<!-- .slide: class="title" -->

<h1 style="font-size: 4ex">Velociraptor: Digging Deeper</h1>

<div class="inset">

## Introducing Velociraptor

### Mike Cohen, Digital Paleontologist
### Rapid 7 Inc

</div>

---

<!-- full_screen_diagram small-font -->
## What is Velociraptor?
### A DFIR tool to handle every stage of the attack timeline

![](attack_timeline.svg)

---

<!-- full_screen_diagram small-font -->
## How can I use Velociraptor?
### Flexible tool to handle every use case

![](use_cases.svg)

---

<!-- .slide: class="content" -->
## Scalable, fast, accurate

* Support Linux, Windows, MacOS, FreeBSD …
* Server simply collects the results of queries - clients do all the
  heavy lifting.
* Client memory and CPU usage is controlled via throttling and active
  cancellations.
* Server is optimized for speed and scalability
* Concurrency control ensures stability
* Bandwidth limits ensure network stability
* Single or multi-server modes (20k EP/server).

---

<!-- .slide: class="content small-font" -->
## Interactively investigate clients

Digital forensic plugins turn VQL into a high quality DFIR tool

![](/modules/gui_tour/vfs_view.png)

---

<!-- .slide: class="content small-font" -->
## Velociraptor Artifacts
### Artifacts encode VQL into user sharable code snippets

![](/modules/artifacts_introduction/artifacts.png)

---

<!-- .slide: class="content small-font" -->
## Hunts - Collecting at scale
### Collecting artifacts at scale from multiple endpoints

![](/modules/secure_shell/select_hunt_artifacts.png)

---

<!-- .slide: class="content small-font" -->
## Postprocessing using Notebooks
### Collaborative shared VQL execution environments

![](/modules/secure_shell/postprocess_hunt.png)

---

<!-- .slide: class="content" -->

## Triaging Using Sigma

* Endpoint tools can directly evaluate Sigma rules on the event logs


<img src="velociraptor_sigma_flow.png" style=" height: 40vh;"/>

---

<!-- .slide: class="content small-font" -->

##  Collecting the sigma artifact

![](collecting_sigma_rules.png)

---

<!-- .slide: class="content small-font" -->

## Triaging an endpoint

![](query_logs.png)


---

<!-- .slide: class="content small-font" -->

## Stacking rules by title

![](stacking_a_column.png)

---

<!-- .slide: class="content small-font" -->

## Viewing the stacking stats

![](viewing_column_stack.png)

---

<!-- .slide: class="content small-font" -->

## Viewing common rows

![](viewing_common_rows.png)

---

<!-- .slide: class="content" -->

## Detection vs. Forensics

* VQL Sigma rules bridge detection with forensics.
* Forensics: `What happened here?`
    * Recover all the information - relevant or not
    * Get a full picture.

* Detection: `What bad things happened here?`
    * Take the forensic information and rapidly zero in on obvious bad
      signals.
    * Not designed to be exhaustive! Triage oriented

* Complimentary capabilities

---

<!-- .slide: class="content" -->

## Real Time Sigma alerting
### VQL is fully asynchronous - real time queries.

<img src="client_events_arch.png" style=" height: 60vh;"/>

---

<!-- .slide: class="content" -->

## Real Time Sigma alerting

### Configuring Velociraptor's client monitoring

![](configuring_client_monitoring.png)

---

<!-- .slide: class="content" -->

## Real Time Sigma alerting

### Configuring Velociraptor's client monitoring

![](configuring_client_monitoring_logs.png)


---

<!-- .slide: class="content small-font" -->

## Live detection with Sigma

![](live_sigma_detection.png)


---

<!-- .slide: class="content small-font" -->
## Administration and automation

* All server administration tasks can be automated with VQL artifacts
* API access available for external automation
* Automatic upload of data to Elastic/Slack/Discord
* Open ended architecture enables novel use cases.

---

<!-- full_screen_diagram small-font -->
## The Velociraptor Ecosystem

<img src="ecosystem.svg" style="height: 70vh;">

---

<!-- .slide: class="content" -->
## Conclusions

We only scratched the surface of what Velociraptor can do!

Check out the following links and join our community…

<table class="noborder">
<tr>
    <td>Docs</td><td>
        <a href="https://docs.velociraptor.app/">https://docs.velociraptor.app/</a>
    </td>
</tr>
<tr>
    <td>Github</td><td>
        <a href="https://github.com/Velocidex/velociraptor">https://github.com/Velocidex/velociraptor</a>
    </td>
</tr>
<tr>
    <td>Discord</td><td>
        <a href="https://docs.velociraptor.app/discord/">https://docs.velociraptor.app/discord/</a>
    </td>
</tr>
<tr>
    <td>Mailing list</td><td>
        <a href="mailto:velociraptor-discuss@googlegroups.com">velociraptor-discuss@googlegroups.com</a>
    </td>
</tr>
</table>
