<!-- .slide: class="title" -->

## Forensic Enrichment
### Enriching events with endpoint state

---

<!-- content -->

## Enrichment with endpoint state

* So far we only applied detection on event logs
* Usually endpoint state provides important context
   * Helps us to triage hits to eliminate false positives

* Many SOC workflows rely on a detection/response loop
   * Analysts need to respond to the detection to retrieve more
     endpoint context

* Velociraptor allows direct enrichment
  * Forensic enrichment adds data outside the pure log source directly
    to the detected event.

---

<!-- content small-font -->
## The Velociraptor Sigma Engine Extensions
### Pre-evaluation enrichment

* Prior to evaluating the rule, the Velociraptor Sigma engine can
  transform the original event using VQL functions.
* This allows us to add additional fields from the live system.
* In this case we add the field `CallChain` to be the [full call
  chain](https://docs.velociraptor.app/vql_reference/other/process_tracker_callchain/)
  of the process mentioned in the event.
   * Velociraptor's process tracker is extremely efficient way to
     calculate process lineage
* Let's try this detection logic:
   * A file is created in the Windows directory
   * The process that created the file is launched interactively

---

<!-- content small-font -->
## Exercise: Enrich data with Process Call Chain

```sql
LET SigmaRules = '''
title: Test Sigma Rule
logsource:
  category: etw
  product: windows
  service: file

detection:
  selection_event_type:
    EventType: CreateNewFile
  selection_filename:
    FileName|re: ^C:\\Windows
  selection_callchain_interactive:
    EventData.CallChain|re: cmd.exe|powershell.exe

  condition: selection_event_type and selection_filename and selection_callchain_interactive

vql: |
  x=>dict(
    EventData=x.EventData + dict(
      CallChain=process_tracker_callchain(id=x.System.ProcessID).Data.Exe))
'''
SELECT *
FROM Artifact.Windows.Sigma.ETWBase(SigmaRules=SigmaRules)
```

---

<!-- content small-font -->
## Enriching detection with endpoint state

<img src="enriched_detection.svg" style="">

---

<!-- content small-font -->
## Post Evaluation Enrichment

* Enriching events before evaluation can be expensive - it happens on
  all events, even the ones that do not match.

* Sometimes we just want to get additional information when the rule
  fires!
  * This is much cheaper as it only happens when the rule fires 🙂
  * But you can not use that enrichment in a detection condition! 🤨

---

<!-- content small-font -->
## Post Evaluation Enrichment

* Expensive enrichment can be added to the `enrichment` section of the rule
* When the rule fires, the enrichment will be added to the results
    * This allows us to capture this state in the event itself!
    * If the file is removed subsequently we still have important context!

```sql

enrichment: |
   x=>dict(PEInfo=parse_pe(file=x.EventData.ProcInfo.Exe))

```

---

<!-- full_screen_diagram small-font -->

## Post Evaluation Enrichment

<img src="post_detection_enrichment.svg" style="">
