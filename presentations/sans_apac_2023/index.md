<!-- .slide: class="title" -->

<h1 style="font-size: 4ex">The evolving frontier of DFIR readiness</h1>

<div class="inset">

## SANS Summit APAC 2023

### Mike Cohen, Digital Paleontologist
### Rapid 7 Inc

</div>

---

<!-- .slide: class="content" -->

## Forensic Readiness

1. Maximising an environment‟s ability to collect credible digital
   evidence
2. Minimising the cost of forensics in an incident response.

[Forensic Readiness John Tan @Stake (2001)](https://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.480.6094&rep=rep1&type=pdf)

---

<!-- .slide: class="content" -->

## What is Digital Forensics anyway?

* Trying to reconstruct the past
  * What occurred on this system?
  * When did this occur?

* Locard's exchange principal

> "Every contact leaves a trace”

---

<!-- .slide: class="full_screen_diagram" -->

<a href="https://www.utica.edu/academic/institutes/ecii/publications/articles/A0AC5A7A-FB6C-325D-BF515A44FDEE7459.pdf">

![](physical_analogy.png)

</a>


---

<!-- .slide: class="content" -->

## Enhancing the efficacy of Forensics

* In the physical world:

  * Surveillance CCTV cameras assist in forensic investigation
  * Procedural methods (registration, auditing etc).
  * Enhanced tracking - e.g. cell phones, GPS etc

* Better forensics acts as a deterrent!

---

<!-- .slide: class="content" -->

## Why do we need Digital Forensics?

* Much of the time we arrive at the "crime scene" after the fact

* Try to reconstruct what happened from incidental information

* Forensics by its nature is **Making the best of a bad hand!**

---

<!-- .slide: class="content" -->

## Why do we need Digital Forensics?

* We use digital forensics to answer tactical questions!

   * We rely on artifacts that were not specifically designed to tell
     us what we want to know.
   * Requires a lot of interpretation to tie unrelated artifacts to
     infer what happened.

* We need to be lucky!

---

<!-- .slide: class="content" -->

## What if we could prepare for forensics?

* Sometimes we go into an incident unprepared, but a lot of the time
  we can prepare in advance!

  * In a corporate setting we can actually prepare for forensic
    investigation and incident response

  * Similar but orthogonal to system hardening

>  Taking steps in advance to increase our chances of successfully
>  investigating an incident!

---

<!-- .slide: class="content" -->

## Levels of preparedness

* Information security is a continuum and a tradeoff between
  resourcing and usability

* What can we do to improve our forensic readiness?

  * Simple things can be done cheaply!
     * Set configuration parameters in the environment.

  * More sophisticated things may involve more efforts
     * Install an agent, EDR etc.

* Consider how likely a forensic investigation will occur?
  * Tradeoff between cost and completeness

---

<!-- .slide: class="content" -->

## What types of interventions can we employ?

* Configuration change
  * Increases the system's ability to support forensic analysis

* Installation of EDR/Agents/Endpoint visibility software
  * Increases resillience againt malicious anti-forensics.

---

<!-- .slide: class="content" -->

## Reimagine the forensic process

* Let's re-examine the digital forensic process critically

* Identify the things that can go wrong, the gaps and improvements

* Can we increase our chances of success?
   * Passive target: No deliberate interference with the DFIR process
   * Active Adversary: Employing Anti-Forensics to frustrate investigation



---

<!-- .slide: class="content" -->

## DFIR Mind map

<a href="https://www.sans.org/posters/windows-forensic-analysis/" >
 The SANS Windows Forensic Analysis Poster
</a>

<img src="windows_forensic_analysis_poster.png" style="height: 400px;">

---

<!-- .slide: class="content" -->

## Event logs

<img src="account_activity.png" style="height: 400px;">

---

<!-- .slide: class="content" -->

## Event logs

* Event logs are a huge source of forensic information!
* But they have some problems:

  * Rotation of event logs. By default event log size is very small
      (20 mb)
  * We can adjust the maximum size of log files through group policy
    or the registry.

---

<!-- .slide: class="content" -->

## Setting event log size

<img src="gpo_event_logs.png">

---

<!-- .slide: class="content" -->

## Clearing the event logs

* Many attackers clear the event logs to frustrate forensic analysis.
   * Use Volume Shadow copies to periodically snapshot the disk
   * Forward events off the system, e.g. for SIEM or even built in
     Event Log forwarding.
   * Detect event log configuration modifications (Registry changes)

---

<!-- .slide: class="content" -->

## Forwarding event logs off the system

* The best practice for protecting event logs is to forward them off
  the system.

   * Built in facility within Windows: [Windows Event Forwarding (WEF)](https://learn.microsoft.com/en-us/windows/security/threat-protection/use-windows-event-forwarding-to-assist-in-intrusion-detection?source=recommendations)

   * Use agent like Elastic or Velociraptor

* Tuning which events to forward
  * Sending only relevant events means less volume
  * Indexing events on the server side may increase costs (can just
    forward for backup).

---

<!-- .slide: class="full_screen_diagram" -->

## Example: Forwarding Sysmon logs

<img src="forwarding_sysmon.png">

---

<!-- .slide: class="content" -->

## Disabling of event logs

* Attackers can actively change logging configuration

<img src="disable-bits-log.png" style="height: 400px;">

---

<!-- .slide: class="content" -->

## What can we do with an agent?

* We can constantly check configuration for compliance

![](checking_log_modifications.png)

---

<!-- .slide: class="content" -->

## What can we do with an agent?

* Real time alerting on configuration modifications

![](detection_log_modifications.png)

---

<!-- .slide: class="content" -->

## File based forensic artifacts

<div class="container">
<div class="col">

## Prefetch files

* Useful to determine evidence of execution
* Only enabled on non-server Windows Versions.

</div>
<div class="col">

<img src="prefetch.png">

</div>

---

<!-- .slide: class="content" -->

## Prefetch files

Enable prefetcher on windows server OS's

```text
reg add "HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\Session Manager\Memory Management\PrefetchParameters" /v EnablePrefetcher /t REG_DWORD /d 3 /f

reg add "HKEY_LOCAL_MACHINE\Software\Microsoft\Windows NT\CurrentVersion\Prefetcher" /v MaxPrefetchFiles /t REG_DWORD /d 8192 /f

powershell /c "Enable-MMAgent -OperationAPI"
```

---

<!-- .slide: class="content" -->

## Filesystem artifacts: USN

* The USN Journal records filesystem operations

   * Operations are recorded in the hidden NTFS file `$Extend\$UsnJrnl:$J`

   * The USN journal rolls over fairly quickly (Approx 30mb)

---

<!-- .slide: class="content" -->

## Using the USN journal in IR

Filtering the USN journal for prefetch file modifications gives
useful timestamps related to program execution!

<img src="usn_filter_for_pf.png">

---

<!-- .slide: class="content" -->

## Querying the USN journal

* The [fsutil tool](https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/fsutil-usn) can be used to manipulate the USN journal.

<div class="container">
<div class="col">

<img src="usn_info.png">

</div>
<div class="col">

<img src="usn_increase.png">

</div>

---

<!-- .slide: class="content" -->

## Attackers may clear the USN journal

* Attacker can completely delete the USN Journal.

* We are forced to carve the disk for USN records!

![](carving_usn.png)

---

<!-- .slide: class="content" -->

## What can we do with an agent?

* Forward USN records off the system in a timely fasion.

<img src="forwarding_usn.png" style="height: 400px">

---

<!-- .slide: class="content" -->

## Forwarding USN logs off the system

<img src="forwarding_usn_prefetch.png" style="height: 400px">

---

<!-- .slide: class="content" -->

## Leveraging ETW for visibility

* Using ETW we gain access to more valuable forensic information
  anyway!

<img src="file_creation_etw.png" style="height: 400px">


---

<!-- .slide: class="content" -->

## Leveraging ETW for visibility

<img src="file_creation_enrichment.png" style="height: 400px">

---

<!-- .slide: class="content" -->

## Process execution

<div class="container">
<div class="col">

* Many artifacts around process execution

* None of them are perfect!

* Missing data:
   * Parent process (Call chain)
   * User that launched the process
   * Limited number of last run timestamps

</div>
<div class="col">

<img src="process_execution_artifacts.png">

</div>

---

<!-- .slide: class="content" -->

## Preparing for Process Execution

* By far the best preparation is to install Sysmon or process logging.

* This has some very small overhead but it is so worth it!

* Process execution logging gives us context as to activity on the
  system.

---

<!-- .slide: class="content" -->
## Tracking processes

* One of the critical questions we ask is `Where did this process come
  from?`

* Context of where the process came from is important in establishing
  initial access vector!

* We could collect all process execution from all endpoints, but:
    * This will generate a large volume of events.
    * Vast majority of events are not interesting.
    * Often determining which process is interesting is determined by
      context.

---

<!-- .slide: class="full_screen_diagram" -->
## Where did notepad process come from?

<img src="process_hacker.png" style="height: 600px" >

---

<!-- .slide: class="content" -->
## Using Generic.System.Pstree

![](collecting_pstree.png)

---

<!-- .slide: class="content" -->
## View process tree

![](pstree.png)

---

<!-- .slide: class="content" -->
## Inspect the process call chain

![](powershell_pstree.png)

---

<!-- .slide: class="content" -->

## Windows Crash Dumps

* Windows Error Reporting (WER) records important information about an
  application when it crashes or hangs.
* This can be extremely important for investigating intrusions
* Exploit code will often general an application crash (E.g. Buffer
  Overflow)

    * https://bmcder.com/blog/extracting-cobalt-strike-from-windows-error-reporting

---

<!-- .slide: class="content" -->

## Windows Crash Dumps

* Settings can be applied for each process

<div class="container">
<div class="col">

    * DumpCount
    * DumpFolder
    * DumpType:
        0. Custom Dump
        1. Mini Dump
        2. Full Dump (Process)

</div>
<div class="col">

![](wer_configuration.png)

</div>

---

<!-- .slide: class="content" -->

## Powershell script logging

* Attackers use PowerShell extensively!
* Script block logging provides visibility into PowerShell activities.

```sh
New-Item -Path "HKLM:\SOFTWARE\Wow6432Node\Policies\Microsoft\Windows\PowerShell\ScriptBlockLogging" -Force

Set-ItemProperty -Path "HKLM:\SOFTWARE\Wow6432Node\Policies\Microsoft\Windows\PowerShell\ScriptBlockLogging" -Name "EnableScriptBlockLogging" -Value 1 -Force
```

---

<!-- .slide: class="content" -->

## Powershell script logging

* Use Group Policy to enable it everywhere

![](powershell_script_logging_gpo.png)

---

<!-- .slide: class="content" -->
## Powershell script logging

* Logging powershell gives a unique view at of attacker activities.

![](powershell_script_logging.png)
