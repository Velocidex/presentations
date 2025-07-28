<!-- .slide: class="title" -->

<h1 style="font-size: 4ex">Detection Engineering</h1>

<div class="inset">

## The art and science of detecting intrusions!

### Mike Cohen, Digital Paleontologist, Rapid7

> [Link to Presentation](https://present.velocidex.com/presentations/2025-auscert-detection_workshop/)

</div>

---

<!-- content -->

## Overview

* This workshop will discuss Detection Engineering with Velociraptor
* We will use Velociraptor as an example platform
   * Not everything we show will be possible on all platforms
   * Nevertheless it is useful to see what is possible!
   * Think about the additional detections that can be implemented
     with better capabilities!
* We also introduce the Sigma detection "standard"

---

<!-- content -->
## Create a local server

* Create a local server on your windows system.
* We will use this server's notebook feature to learn about windows artifacts
* Run Velociraptor on your machine
    * Download Velociraptor from the
      [Download](https://docs.velociraptor.app/downloads/) page (`.msi`
      or `.exe`)

```
velociraptor-v0.74.2-windows-amd64.exe gui
```

---

<!-- .slide: class="full_screen_diagram" -->

The "gui" command creates an instant temporary server/client with self
  signed SSL and a hard coded admin/password.

<img src="/modules/gui_tour/velociraptor-gui.png" style="height: 60vh">

---

<!-- .slide: class="full_screen_diagram" -->
## Your Velociraptor is ready to use!

![](/modules/gui_tour/GUI.png)
