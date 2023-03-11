<!-- .slide: class="title" -->

## Web shell
###  A hacker's favorite backdoor

![](/modules/web_shell/webshell.jpg)

---

<!-- .slide: class="content" -->
## Web shell
### A hacker's favorite backdoor

A webshell is a way for an attacker to obtain a shell over the web
* Sometimes this is done by adding malicious active content (PHP,ASP,CGI)
* Often it is installed as an additional service that looks like a web service
* Can contain authentication and/or encryption.

---

<!-- .slide: class="content" -->
## Web shell
### Exercise

In this exercise we start an opensource web shell, typical of
malicious web shells.

```bash
$ git clone https://github.com/JiangKlijna/web-shell
$ cd web-shell
$ make
$ ./web-shell -s -P 2020 -u admin -p admin -cmd bash
```

This creates the server web shell. (Alternatively download the Github
release)

---

<!-- .slide: class="content" -->
## Web shell
### The client

Connect to the local webshell server using the client over HTTP (with
the above username and password for authentication.

```
$ ./web-shell -c -H localhost -P 2020 -u admin -p admin
$ ping www.google.com
PING www.google.com (172.217.24.36) 56(84) bytes of data.
64 bytes from hkg07s23-in-f4.1e100.net (172.217.24.36): icmp_seq=1 ttl=114 time=27.9 ms
64 bytes from hkg07s23-in-f36.1e100.net (172.217.24.36): icmp_seq=2 ttl=114 time=27.7 ms
```

---

<!-- .slide: class="content" -->
## Detection
### Detecting the web shell via network

How would we detect this webshell?

* Search for listening ports and enrich

Use the `Linux.Network.NetstatEnriched` artifact to look for listening
ports.

---

<!-- .slide: class="content" -->
## Detection
### Listening ports

![](/modules/web_shell/netstat_enriched.png)

---

<!-- .slide: class="content" -->
## Detection
### Listening ports

![](/modules/web_shell/netstat_enriched_results.png)

---

<!-- .slide: class="content" -->
## Detection
### Process Tree

![](/modules/web_shell/process_tree.png)

---

<!-- .slide: class="content" -->
## Detection
### Looking in memory

Some attack tools do not have files on disk
1. File may be packed/encrypted
2. On Linux the file may be deleted (but process is running)
3. Malicious code may be injected into another process

Scan in memory!  *Yara scan*

---

<!-- .slide: class="content" -->
## Detection
### Yara scan in memory

![](/modules/web_shell/yara_scan_memory.png)

---

<!-- .slide: class="content" -->
## Detection
### Yara scan in memory

![](/modules/web_shell/yara_scan_hits.png)

---

<!-- .slide: class="title" -->

## Custom Detection
### Developing custom detection artifacts

<img src="/modules/web_shell/detection.jpg" style="height: 400px" >

---

<!-- .slide: class="content" -->
## What are we looking for?

Ultimately a Web Shell is:

1. A shell process (e.g. Bash or Powershell)
2. Spawned from another process along the call chain
3. One of the processes along the call chain is listening for network
   connections.

Examples:
```
    systemd -> screen -> bash -> mc -> bash -> web-shell -> bash
    systemd -> sshd -> sshd -> bash
```

---

<!-- .slide: class="content" -->
## Writing custom artifacts
### Developing in VQL

```sql
-- Get all listening Pids and store in a lookup structure
LET ListeningPids <= memoize(
key="Pid",
query={
    SELECT Laddr.ip AS IP,
         Laddr.port AS Port,
         str(str=Pid) AS Pid
  FROM connections()
  WHERE Status =~ "LISTEN"
   AND Pid != "0"
})

-- A Function to determine if any pid in the Pids list is listening
LET IsPidListening(Pids) = SELECT _value FROM foreach(row=Pids)
  WHERE get(item=ListeningPids, field=str(str=_value))

-- Get all bash processes and their callchain
LET CallChains = SELECT Pid,
  Name,
  CommandLine,
  join(array=process_tracker_callchain(id=Pid).Data.Name, sep=" -> ") AS CallChain,
  process_tracker_callchain(id=Pid).Data.Pid AS CallChainPids
FROM process_tracker_pslist()
WHERE Name =~ "bash"

-- A Suspicious shell is one that is spawned from
-- any process which also has a listening port.
SELECT Pid, Name, CommandLine, CallChain,
  {
    SELECT process_tracker_get(id=_value).Data AS Details
    FROM IsPidListening(Pids=CallChainPids)
    LIMIT 1
  } AS ListeningProcess
FROM CallChains
WHERE ListeningProcess
```

---

<!-- .slide: class="content" -->
## Writing custom artifacts
### Testing in a notebook

![](/modules/web_shell/create_notebook.png)

---

<!-- .slide: class="content" -->
## Writing custom artifacts
### Add a new VQL Cell

![](/modules/web_shell/add_vql_cell.png)

---

<!-- .slide: class="content" -->
## Writing custom artifacts
### Running test VQL on server

![](/modules/web_shell/develop_vql_in_cell.png)

---

<!-- .slide: class="content" -->
## Writing custom artifacts
### Create a new artifact

![](/modules/web_shell/create_new_artifact.png)
