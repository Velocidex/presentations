
<!-- .slide: class="content" -->
## What is Velociraptor?
* Velociraptor is a unique Free and Open Source DFIR tool, giving you
  power and flexibility through the Velociraptor Query Language
* VQL is used to drive a powerful set of forensic capabilities:
* Using VQL we can write custom "Artifacts" to identify emerging threats quickly and safely
* Hunt for artifacts at scale over thousands of end points within
  minutes!



---


<!-- .slide: class="full_screen_diagram" -->
## Deployment overview
![](/modules/overview/deployment_overview.svg)

---


<!-- .slide: class="content" -->
## Typical deployments

* Support Linux, Windows, MacOS, FreeBSD …
* Server simply collects the results of queries - clients do all the heavy lifting.
* Client memory and CPU usage is controlled via throttling and active cancellations.
* Server is optimized for speed and scalability
* Concurrency control ensures stability
* Bandwidth limits ensure network stability

---


<!-- .slide: class="content" -->
## Typical deployments

Current recommendations
* 10k-15k clients - single server with file based data store (usually cloud VM).
* SSL load is the biggest load - TLS offloading helps a lot!
* 8 GB RAM/8 cores is generous towards the top of the range.
* We recommend Ubuntu/Debian server
    * 15-20k to 150k endpoints we recommend a multi-frontend setup.
