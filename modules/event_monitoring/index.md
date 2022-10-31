
<!-- .slide: class="title" -->

# Monitoring events from endpoints

---

<!-- .slide: class="content" -->
## What are event artifacts?

Event artifacts are never-ending VQL queries that watch for events on
clients and stream those events to the server.

![](/modules/event_monitoring/event-queries.png)

---

<!-- .slide: class="content" -->
## Client event tables

* Monitoring
    * Clients can be made to monitor events and forward them to the server.

* Response
    * Clients can automatically respond to events autonomously
    * E.g. Kill processes, quarantine machines etc.

---

<!-- .slide: class="content" -->

## Enable sysmon collection

<div class="container small-font">
<div class="col">

* Client event queries are targeted by label group.
* Sysmon will be installed automatically and events will be forwarded.

</div>
<div class="col">

<img src="/modules/event_monitoring/adding-monitoring-rules.png" style="bottom: inherit" class="inset" />
</div>
</div>

---

<!-- .slide: class="full_screen_diagram" -->
## Viewing Sysmon events relayed to the server

![](/modules/event_monitoring/viewing-sysmon-events.png)

---

<!-- .slide: class="content" -->
## Turning artifacts into a detection

* We have previously looked at log enable/disable by examining registry keys.
* Can we detect when these registry keys are changing?
* The diff() plugin periodically runs a query and reports on changes.

* Install the `Windows.Events.EventLogModifications` artifact

---

<!-- .slide: class="full_screen_diagram" -->
## Windows.Events.EventLogModifications

![](/modules/event_monitoring/Windows.Events.EventLogModifications.png)

---

<!-- .slide: class="full_screen_diagram" -->
## System changes relayed to server

Good for slow changes

![](/modules/event_monitoring/Windows.Events.EventLogModifications_results.png)

---

<!-- .slide: class="title" -->
# USN Journal monitoring

## File modification monitoring at scale.

---

<!-- .slide: class="content" -->
## USN Journal

* We have previously seen that the USN journal is useful for
  recovering evidence of file modification.
* Sadly in practice the USN journal rolls over fairly quickly (days!)
* Wouldn't it be nice to feed the events to the server continuously?

---

<!-- .slide: class="content" -->
## Windows.Detection.USN

* Enable the Windows.Detection.USN artifact - target paths of
  interest.

![](/modules/event_monitoring/Windows.Detection.USN.png)

---

<!-- .slide: class="full_screen_diagram" -->
## Inspect streaming results

* See direct evidence of execution, tasks creation etc.

![](/modules/event_monitoring/Windows.Detection.USN_results.png)
