
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

---


<!-- .slide: class="content" -->
## Automate file collection+parsing

The goal of VQL is to automate as much of the routine DFIR work as possible

* Example: Collect all browser artifacts

---


<!-- .slide: class="content small-font" -->
## Generic.Collectors.SQLECmd

* Use Generic.Collectors.SQLECmd to automatically locate and parse SQLite files.
* SQLECmd is an open source project to document location of SQLite files
used by various programs (like browsers)

![](/modules/artifacts_and_vql_intro/sqlite-parsing.png)


<!-- .slide: class="content" -->
## Velociraptor's plugins are robust

Handle file reading errors gracefully

![](/modules/artifacts_and_vql_intro/sqlite-error-recovery.png)

---

<!-- .slide: class="content small-font" -->
## Inspect the data with the table widget
* Show or hide columns
    * Export the modified table to CSV or JSON

![](/modules/artifacts_and_vql_intro/hide-columns.png)


---

<!-- .slide: class="content small-font" -->
## Inspect the data with the table widget
* Filter or sort using the table widget

![](/modules/artifacts_and_vql_intro/table-filter.png)

---

<!-- .slide: class="content small-font" -->
## Inspect the data with the table widget
![](/modules/artifacts_and_vql_intro/transform-table.png)
