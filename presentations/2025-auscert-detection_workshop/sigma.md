<!-- .slide: class="title" -->

## Sigma Rules - Vibe detection 🧘

```yaml
title: PsExec Service Start
description: Detects a PsExec service start
author: Florian Roth (Nextron Systems)
logsource:
  category: process_creation
  product: windows
detection:
  selection:
    CommandLine: C:\Windows\PSEXESVC.exe
  condition: selection
```

[proc_creation_win_sysinternals_psexesvc_start.yml](https://github.com/SigmaHQ/sigma/blob/b062d8ad650054cd20836d5ba38031090b8d8c33/deprecated/windows/proc_creation_win_sysinternals_psexesvc_start.yml)

---

<!-- content -->

## Sigma Rules

* A Sigma rule specifies how to detect a particular attack
   * Think of it as **grep** for event logs


* Vibe detection
   * *logsource* section specifies an event source to match the rule against.
   * *detection* clause contains a list of `selections` joined into a logical `condition`.
   * *condition* clause specifies how to combine the detection clauses logically.
   * Selections refer to abstract fields that map to actual fields
     within the event. These mappings are called `Field Mappings`.

---

<!-- content -->
## Sigma Rules

* The Sigma standard does not define
   * What log sources are actually available
   * The specific structure of each event
   * What fields are available and what they are called.

* The intention is to convey the "vibe" of a detection.
   * A Sigma Compiler for the target SIEM is used to convert this
     "vibe" to a concrete detection for a particular engine.

---

<!-- content small-font -->

## Sigma Rule Syntax

* Covered in details [in the Sigma HQ documentation site](https://sigmahq.io/docs/basics/rules.html)
    * Based on YAML
* Detections are written in YAML
    * Include Selections and Condition clause
    * Selections can be OR or AND
    * Condition combine selection logically

```
detection:
  selection_and:
    displaymessage: Max sign in attempts exceeded
  selection_or:
    field_name:
      - this # or
      - that
  condition: selection_and and selection_or
```

---

<!-- content small-font -->

## Sigma Rule Syntax

* Modifiers apply to the field to allow transformation for matching.
* Although it looks like they can be applied in any order, only a few
  combinations make sense.
* Velociraptor `vql` modifier allows running arbitrary VQL code on a field.

```
detection:
  selection:
     - TargetFilename|endswith: '.cmdline'
     - fieldname|base64offset|contains:
         - /bin/bash
     - fieldname|contains: needle
     - fieldname|re: .*needle$
     - fieldname|startswith: needle
     - EventData|vql:
         x=>hash(path=x.Filename).MD5 =~ '12345'
```

---

<!-- full_screen_diagram small-font -->

## The Sigma Compiler

<img src="sigma_architecture.svg" style="height: 90vh">

---

<!-- content -->

## The Sigma Model

### Converting a Vibe to a concrete detection

* To convert a Sigma rule to a concrete detection
    * **Field Mappings**: Mapping between abstract field names and
      concrete field names.

    * **Log sources**: Mapping between abstract log source
      specification and concrete data sources.

    * Code to convert the conditional logical clauses to platform
      specific query against the backend.

* This is done by the Sigma Compiler for the target.
    * e.g. for Elastic

---

<!-- content -->

## The Sigma Model

### Converting a Vibe to a concrete detection

* It generally does not make sense to speak of a `Sigma Rule` without
  knowing the exact `model` used.

* The Sigma standard is not really portable, only the Vibe is
  portable.

  * When converting rules from one model, it helps to develop a new
    model with a 1:1 mappings.

  * For example, a Velociraptor Model to consume Sigma Rules written
    for the ECS stack.

---

<!-- content small-font -->
## The Velociraptor Sigma Architecture

* Velociraptor has a built in Sigma engine
    * Sigma is the preferred built in method for scalable detection and triaging!
* Accepts a model definition:
    * Log sources are VQL queries that generate events
    * Field mappings are VQL Lambda functions that resolve fields in the rule.
* Velociraptor terminology
    * `Sigma Model` == `Sigma Compiler`
* Rules and Models are pushed to the endpoint for direct evaluation
* Only Matches are forwarded to the server.

---

<!-- full_screen_diagram small-font -->
## The Velociraptor Sigma Architecture
### Velociraptor supports multiple models at the same time!

<img src="velociraptor_sigma_flow.svg" style="">

---

<!-- content -->

## The Velociraptor Curated Sigma Project

* A Project to maintain and curate:

   * Useful set of `Sigma Models` for different situations
   * A Curated set of `Sigma Rules` tailored for use in the models.

* Main goal is to separate the `maintainance of the Model` and the
  `writing of the rules`!

> https://sigma.velocidex.com/

---

<!-- full_screen_diagram small-font -->

## The Velociraptor Curated Sigma Project

<img src="sigma_site_models.svg" style="">

---

<!-- content small-font -->

## The Windows Base Model

* Most public Sigma Rules address Windows Event Logs
* Designed to be compatible with Sigma Rules in the wild.
    * Usually this model will work with most Sigma Rules out there.
    * Log sources are compatible with [Sigma HQ](https://sigmahq.io/docs/basics/log-sources.html) Windows Event Logs
* This Sigma Model is used by the `Windows.Hayabusa.Rules` artifact
    * Many projects out there use a similar model (e.g. `Hayabusa`,
      `ChainSaw` etc).
    * This model defines

---

<!-- full_screen_diagram small-font -->

## The Windows Base Model
### A dedicated model for windows event logs

<img src="windows_base_model.svg" style="">

---

<!-- full_screen_diagram small-font -->

## The Windows Base Model
### Log sources consume event logs

<img src="windows_base_model_2.svg" style="">

---

<!-- full_screen_diagram small-font -->

## The Windows Base Model
### Field mappings access event data


<img src="windows_base_model_3.svg" style="">

---

<!-- full_screen_diagram small-font -->

## Importing The Velociraptor Sigma Artifacts
### Using the built in artifact

<img src="import_sigma_artifacts.svg" style="">


---

<!-- full_screen_diagram small-font -->

## Importing The Velociraptor Sigma Artifacts
### By uploading manually

<img src="uploading_artifact_pack.svg" style="">

---

<!-- full_screen_diagram small-font -->

## Recap: What is Sigma ?

> It's just the vibe of the thing!

<iframe width="560" height="315" src="https://www.youtube.com/embed/97IiPli_uXw?si=VLrvR1K82vKOt5OG&amp;start=48" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>

---

<!-- content -->

## Recap: Sigma traps

* You can not just blindly copy Sigma Rules from one system to another
* You have to carefully check that the sigma models are compatible
   * Log sources refer to the same events
   * Field mappings have corresponding mappings in the two models.

* E.g. Sigma rules written for Sysmon EID 1 can not be reliably used
  by Windows Audit EID 4688
