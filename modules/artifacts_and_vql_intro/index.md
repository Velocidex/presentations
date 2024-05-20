
<!-- .slide: class="content" -->
## VQL - Velociraptor's magic sauce

Rather than having specific analysis modules, VQL allows generic capabilities to be combined in novel creative ways
* NTFS/MFT/USN/Glob file system analysis
* File parsers - Grok, Sqlite etc
* Built in powerful parser framework for novel binary parsers


---

<!-- .slide: class="content" -->

## Velociraptor Artifacts
Velociraptor comes with a large number of artifact types
* Client Artifacts run on the endpoint
* Client Event artifacts monitor the endpoint
* Server Artifacts run on the server
* Server Event artifacts monitor for events on the server.

---


<!-- .slide: class="content" -->

## Velociraptor Query Language

Using a query language we can string together different forensic
capabilities to create novel analysis

The Power of Open source!

The [Velociraptor artifact
exchange](https://docs.velociraptor.app/exchange/) is a place for the
community to publish useful VQL artifacts for reuse

---

<!-- .slide: class="full_screen_diagram" -->
## The Artifact Exchange

![](/modules/artifacts_and_vql_intro/artifact-exchange.png)


---

<!-- .slide: class="full_screen_diagram" -->
## Automatically import Exchange

![](/modules/artifacts_and_vql_intro/import-exchange.png)

<!-- .slide: class="title" -->
# Searching for files
## Let's start at the beginning….

---

<!-- .slide: class="content small-font" -->

## Manually importing artifact packs

You can manually upload an artifact pack as well (A zip file
containing artifact definitions).

![](/modules/artifacts_and_vql_intro/import_pack.png)

---

<!-- .slide: class="content" -->
## Finding files

DFIR is often about finding files on the endpoint
* Filename is sometimes an indicator
* Word documents in a temp folder may contain macros
* Sometimes we need to filter by file content
* File has signature of malicious macro/script


---


<!-- .slide: class="content" -->
## Windows.Search.FileFinder

<div class="container">
<div class="col">

* Glob based
* Time filters
* Yara for Content

</div>
<div class="col">

![](/modules/artifacts_and_vql_intro/file-finder-args.png)

</div>
</div>

---


<!-- .slide: class="content" -->
## Exercise

Find all executables in the user's home directory
