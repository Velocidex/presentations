
<!-- .slide: class="title" -->

# Velociraptor Installation and GUI tour

---


<!-- .slide: class="content" -->
## Create a local server

* Create a local server on your windows system.
* We will use this server's notebook feature to learn about windows artifacts
* Run Velociraptor on your machine
    * Download Velociraptor from GitHub (.msi or .exe)

```
velociraptor-v0.6.6-2-windows-amd64.exe gui
```

---


<!-- .slide: class="full_screen_diagram" -->

The "gui" command creates an instant temporary server/client with self
  signed SSL and a hard coded admin/password.

![](/modules/gui_tour/velociraptor-gui.png)

---

<!-- .slide: class="full_screen_diagram" -->
## Your Velociraptor is ready to use!

![](/modules/gui_tour/GUI.png)
---

<!-- .slide: class="title" -->
# A Velociraptor GUI tour

<img src="/modules/gui_tour/tour-bus.png" class="title-inset">

---

<!-- .slide: class="content" -->

## The Dashboard

* The Dashboard shows the current state of the installation:
    * How many clients are connected
    * Current CPU load and memory footprint on the server.
    * When running hunts or intensive processing, memory and CPU requirements will increase but not too much.
    * You can customize the dashboard - it’s also just an artifact.

---

<!-- .slide: class="full_screen_diagram" -->

## The Dashboard

![](/modules/gui_tour/dashboard.png)

---

<!-- .slide: class="title" -->

# Interactively investigate individual clients

---


<!-- .slide: class="content" -->
## Searching for a client

To work with a specific client we need to search for it.
Press the **Search** or **Show All** icon to see some clients

![](/modules/gui_tour/search_clients.png)

---

<!-- .slide: class="content" -->
## Search for clients
### hostname, label, or client ID.
* You can start typing the hostname to auto-complete
* Some common terms:
   * host: hostnames
   * mac: Mac addresses
   * ip: last seen IP address
   * label: Search by labels

---


<!-- .slide: class="content small-font" -->
## Client Overview

* Internally the client id is considered the most accurate source of
endpoint identity

![](/modules/gui_tour/client_overview.png)

---

<!-- .slide: class="content small-font" -->
## Shell commands

* Velociraptor allows running shell commands on the endpoint using
  Powershell/Cmd/Bash
    * Only Velociraptor users with the administrator role are allowed to
  do this!
    * Actions are logged and audited

![](/modules/gui_tour/shell_commands.png)

Note:

```powershell
Get-LocalGroupMember -Group "Administrators"
```

---


<!-- .slide: class="title" -->
# Interactively fetching files from the endpoint

<img src="/modules/gui_tour/fetch.png" class="title-inset">

---

<!-- .slide: class="content small-font" -->
## The VFS View

Remember that the VFS view is simply a server side cache of
information we know about the endpoint - it is usually out of date!

![](/modules/gui_tour/vfs_view.png)

---


<!-- .slide: class="content small-font" -->
## Navigating the interface

* Click the “Refresh this directory” will schedule a directory listing
  artifact and wait for the results (usually very quick if the
  endpoint is online).
* The “Recursively refresh this directory” will schedule a recursive
  refresh - this may take some time! After this operation a lot of the
  VFS will be pre-populated already.
* “Collect from client” will retrieve the file data to the
  server. After which, the floppy disk sign indicates that we have
  file data available and you can click the “Download” link to get a
  copy of the file.


---

<!-- .slide: class="full_screen_diagram" -->
## The VFS interface

![](/modules/gui_tour/vfs_view_2.png)

---


<!-- .slide: class="title" -->
# Velociraptor Artifacts

## Fast, Efficient, Surgical

<img src="/modules/gui_tour/surgical.png" style="bottom: -200px" class="title-inset">

---

<!-- .slide: class="content small-font" -->

## Why a query language?
* Able to dynamically adapt to changing requirements - without needing to rebuild clients or servers.
* For example, a new IOC is released for detection of a specific threat
    * Immediately write a VQL artifact for the threat, upload the artifact and hunt everywhere for it.
    * Turn around from IOC to full hunt: A few minutes.
    * Share artifacts with the community
* VQL Artifacts are simply YAML files with VQL queries.
    * Can be easily shared and cross pollinate other Artifacts
    * Can be customized by callers.
    * [Public Artifact Reference](https://docs.velociraptor.app/artifact_references/)


---


<!-- .slide: class="content" -->

## What is VQL?

```sql
SELECT X, Y, Z FROM plugin(arg=1) WHERE X = 1
```

* `X, Y, Z` are called Column Selectors
* `plugin(arg=1)` is termed a VQL Plugin with Args
* `X = 1` is the Filter Condition


---


<!-- .slide: class="content" -->
## Velociraptor artifacts

Velociraptor is just a VQL engine!

* We package VQL queries in Artifacts:
    * YAML files
    * Include human description
    * Package related VQL queries into “Sources”
    * Take parameters for customization
    * Can in turn be used in VQL as well...

---

<!-- .slide: class="content" -->

## What does the VFS view do under the cover?

* Refreshing the VFS simply schedules new artifacts to be collected - it is just a GUI convenience.

![](/modules/gui_tour/vfs_collections.png)

---


<!-- .slide: class="content" -->
## Velociraptor uses expert knowledge to find the evidence

A key objective of Velociraptor is encapsulating DFIR knowledge into
the platform, so you don’t need to be a DFIR expert.  We have high
level questions to answer We know where to look for evidence of user /
system activities

We build artifacts to collect and analyze the evidence in order to answer our investigative questions.

---

<!-- .slide: class="full_screen_diagram" -->
## Anatomy of an artifact

![](/modules/gui_tour/artifacts.png)

---

<!-- .slide: class="content small-font" -->
## Collecting new artifacts

* To collect a new artifact, from the Collected Artifacts screen,
click Collect new artifact and search for it. Select Add to add it to
this collection. When finished simply click Next.

![](/modules/gui_tour/new-collections.png)

---

<!-- .slide: class="content small-font" -->
## Configuring the artifact collection

* Many artifacts take parameters that can control the way they work.

![](/modules/gui_tour/configure_artifacts.png)

---


<!-- .slide: class="content small-font" -->
## What do artifacts return?

* All artifacts produce rows since they are just queries.
* Some artifacts also upload files. You can create a download zip to export all the uploaded files.

![](/modules/gui_tour/artifact-results.png)

---


<!-- .slide: class="content small-font" -->
## Uploaded files

* The uploads tab shows the file's location on the server.
* You can download each one individually.

![](/modules/gui_tour/artifact-uploads.png)

---


<!-- .slide: class="content small-font" -->
## Artifact query logs

* As the query is running on the endpoint any log messages are sent to the server.
* Click the log tab to see if there were any errors and how many rows are expected.

![](/modules/gui_tour/artifact-logs.png)

---


<!-- .slide: class="content small-font" -->
## Artifacts return multiple tables (sources)
* `Source Selector`: Viewing the result tab shows the rows sent from
  every artifact and source.

![](/modules/gui_tour/artifact-sources.png)

---

<!-- .slide: class="content small-font" -->
## Exporting artifact collections
* Use the GUI to create a zip export of the collection
![](/modules/gui_tour/export-collection.png)
