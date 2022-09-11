<!-- .slide: class="title" -->
# Cloud Native Velociraptor

<div class="inset">

## VeloCon 2022

### Mike Cohen, Digital Paleontologist

</div>

---

<!-- .slide: class="content" -->
## History of Velociraptor

Google GRR was the first large scale hunting platform but (at the
time) had issues with performance

* Written in a slow language
* A lot of work done server side
* Extensive client data model
* Cloud first architecture - worker, frontend, GUI
* Uses a database to store bulk data

---

<!-- .slide: class="full_screen_diagram" -->
## GRR Overview

<img style="height: 400px"
    src="grr_architecture.png">
<div>
    <a href="https://storage.googleapis.com/docs.grr-response.com/MoserCohenHunting.pdf">GRR Architecture</a>
</div>

---

<!-- .slide: class="content" -->
## Velociraptor was born!

Velociraptor was designed to address some of the problems with earlier
systems.

* Written in a fast and efficient language
* Designed for a small deployment (<10k endpoints)
* Query language allows the server to be very simple
* No inherent data model - Schema-less
* Easy to deploy
* File storage backend

---

<!-- .slide: class="full_screen_diagram" -->
## Velociraptor Architecture

<img style="height: 400px"
    src="https://docs.velociraptor.app/blog/html/2018/08/10/design_differences_between_velociraptor_and_grr/image1.png">
<div>
    <a href="https://docs.velociraptor.app/blog/html/2018/08/10/design_differences_between_velociraptor_and_grr/">Design Differences Between Velociraptor And Grr</a>
</div>


---

<!-- .slide: class="content" -->
## Multi Frontend Velociraptor

To scale larger than one server we support multi-frontend architecture

* Uses distributed filesystem (EFS)
* Multiple server in a Master-Minion architecture
* Message passing between servers ensures conherency.
* Still no additional dependency on cloud infrastructure

---

<!-- .slide: class="full_screen_diagram" -->
## Multi-frontend Architecture

![Velociraptor Multi-frontend Architecture](https://docs.velociraptor.app/blog/img/0cb6CiHW_m4COLHtf)

---

<!-- .slide: class="content" -->
## Velociraptor in the cloud

To scale even more we need to replace the backend with cloud centric
services

* Horizontally auto-scale servers
* Use cloud managed storage
* Remove need for state on the server
* Not for everyone! More complexity and cost involed...

---

<!-- .slide: class="content" -->
## Organization of codebase

* We do not want to remove the current architecture!
* Cloud Velociraptor is a separate project with different goals!
* Previously code abstracted at the storage layer
* Now code is abstracted at a higher layer we call services.

---

<!-- .slide: class="content" -->
## Velociraptor Services

Velociraptor services are swappable high level utilities that are used
in the codebase. e.g.

* Indexing service: Used to search for clients.
* Label service: Used to label a client.
* ClientInfo service: Used to manage information about clients.
* Repository service: Used to manage artifacts
* Launcher service: Used to compile and launch collections

---

<!-- .slide: class="full_screen_diagram" -->
## Velociraptor Services

![Velociraptor Services](velociraptor_services.png)


---

<!-- .slide: class="content" -->
## Velociraptor Services

Having high level services allows us to swap implementations freely
and compose Velociraptor around different backend architectures

* This is a much more flexible abstraction than around data storage
  alone.
* Allows us to use a separate project to manager different
  implementations

---

<!-- .slide: class="content" -->
## Cloud Velociraptor

A new experimental GitHub Project https://github.com/Velocidex/cloudvelo

* Uses S3 for bulk storage
* Opensearch for structured storage
* Stateless servers (can be horizontally scaled)
* Not all features are possible in this architecture.

---

<!-- .slide: class="content" -->
## Cloud Architecture

<div style="text-align: center">
<img src="cloud_velo_architecture.svg"
    style="height: 50vh">
</div>

---

<!-- .slide: class="content" -->
## Cloud Architecture

* Clients connect to a stateless frontend with an ingestor
* The ingestor writes the data in S3 or Opensearch
* The GUI runs in a stateless server
* The foreman is a batch processing job - schedules hunts, updates client event queries.

---

<!-- .slide: class="content" -->
<h1 style="margin-top: 20vh">Demo Time</h1>

---

<!-- .slide: class="content" -->
## Current limitations

The Cloud Velociraptor implementation has some limitations currently
(but they might be removed in future)

* Clients are currently polling.
* Missing notifications so updates are slower.
* Opensearch backend is much slower than attached storage!
* More expensive and complex to maintain because we have more moving pieces...

---

<!-- .slide: class="content" -->
## Future improvements

* Potential performance improvements through better schema / OpenSearch optimizations
* Currently we only use OpenSearch as a simple database - maybe a more performant database?
* Currently no server event monitoring!
* Will remain a separate project - an alternative to the standard Velociraptor...

---

<!-- .slide: class="content" -->
## Conclusions

* Velociraptor in the cloud is a viable alternative.
* We hope it can scale to much larger deployment sizes.
* It can be a lot more expensive to run though!

<table class="noborder">
<tr>
    <td>Docs</td><td>
        <a href="https://docs.velociraptor.app/">https://docs.velociraptor.app/</a>
    </td>
</tr>
<tr>
    <td>Github</td><td>
        <a href="https://github.com/Velocidex/cloudvelo">https://github.com/Velocidex/cloudvelo</a>
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
