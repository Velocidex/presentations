<!-- .slide: class="title" -->

# Offline collection

## Collecting artifacts without Velociraptor clients

---

<!-- .slide: class="content" -->
## Why Offline collection?

* I want to collect artifacts from an endpoint
* But Velociraptor is not installed on the endpoint!
* Or the endpoint is inaccessible to the Velociraptor server (no egress, firewalls etc).
* But Velociraptor is just a VQL engine!  It does not really need a server anyway

---

<!-- .slide: class="content" -->
## Create an offline collector

<div class="container">
<div class="col">

![](/modules/offline_collector/offline_builder.png)

</div>
<div class="col">

#### Let's select two artifacts to collect:

1. Windows.KapeFiles.Targets
    * Select the Basic Collection to get many forensic artifacts
2. Generic.Collectors.SQLECmd

</div>

---

<!-- .slide: class="full_screen_diagram" -->

### Selecting the Windows.KapeFiles.Targets artifact

![](/modules/offline_collector/offline_kape_targets.png)

---

<!-- .slide: class="full_screen_diagram" -->

### Configuring the collector to encrypt output

![](/modules/offline_collector/offline_configure.png)

---

<!-- .slide: class="full_screen_diagram" -->

### Downloading the prepared binary

![](/modules/offline_collector/offline_download_binary.png)

---

<!-- .slide: class="content" -->

## Offline collector binaries

* Preconfigured to collect the required artifacts
* No user interaction needed - just run with no command line args
* Prepare armoured Zip file with all the results in them

---

<!-- .slide: class="full_screen_diagram" -->

### Acquire data!

![](/modules/offline_collector/offline_acquire.png)

---

<!-- .slide: class="content" -->

## Acquired file is encrypted

* Due to limitations in the Zip format, file names can not be encrypted.
* Therefore, Velociraptor creates a second protected Zip file inside
  the outer container.
* Several encryption schemes supported:
    1. Regular password
    2. X509 - random password generated and encrypted with the server's certificate.
    3. GPG - random password generated and encrypted with the GPG public key.


---

<!-- .slide: class="content" -->

## Acquired file is encrypted

![](/modules/offline_collector/offline_encrypted.png)

---

<!-- .slide: class="content" -->

## Importing into Velociraptor

* Velociraptor can automatically decrypted offline containers when
  importing.
* Use the Server.Utils.ImportCollection artifact to import collections
* The server uses its private key to unlock the container automatically.
* This preserves PII and confidential information in transit!


---

<!-- .slide: class="full_screen_diagram" -->

### Import the collection into the Velociraptor server

![](/modules/offline_collector/offline_import.png)

---

<!-- .slide: class="full_screen_diagram" -->

### Inspect the import process

![](/modules/offline_collector/offline_import_inspect.png)

---

<!-- .slide: class="full_screen_diagram" -->

### Inspect the collected data

![](/modules/offline_collector/offline_import_inspect_data.png)
