<!-- .slide: class="title" -->
## Artifacts Of Autumn #37

<a href="https://twitter.com/therealwlambert/status/1586010158848622592" target="_blank">
 <img src="/modules/artifacts_of_autumn/37/tweet.png">
</a>

---

<!-- .slide: class="content" -->

## Mounting ISO files in Windows

* Normally files downloaded from the web have the Mark Of the Web (ADS).
* This limits their ability to run macros etc.
* Files may be embedded inside an ISO file.
* ISO files may be mounted by Explorer.

---

<!-- .slide: class="content" -->
## Exercise - Atomic Red Team

Download the sample ISO file from the [Atomic Red Team](https://github.com/redcanaryco/atomic-red-team/blob/master/atomics/T1553.005/T1553.005.md#atomic-test-4---execute-lnk-file-from-iso)

* You might need to disable Windows Defender for this one!

* Double click on the ISO file to mount it.

---

<!-- .slide: class="content" -->
## Detecting the attack

We can check the event logs for mounting of the ISO.

![](/modules/artifacts_of_autumn/37/artifact.jpeg)
