<!-- .slide: class="title" -->
# Velociraptor: Digging Deeper…

<div class="inset">

## DFRWS APAC 2022

### Mike Cohen, Digital Paleontologist

</div>

<img src="/resources/velo_bike.gif" class="fixed" style=" right: 0px;  bottom: 0px; height: 300px; z-index: -10;"/>

---

<!-- .slide: class="content" -->
## Overview

There is no way we can cover all of the capabilities Velociraptor offers in the time today!

* This workshop will give a taste to how modern threat hunting is done
  at scale
* We will look at some of the common scenarios and attacks and some of
  the modern detection techniques we can employ to find these.
* This workshop will be a "Follow Along" workshop - please try these
  exercises on your own machine!

---

<!-- .slide: class="content" -->
## Prerequisites

In order to follow along with this workshop you will need to use a
windows VM with administrator level access.  You can grab a free VM
from Microsoft

* Please ensure your VM has .NET version 4+ with MSBuild -
dotNetFx40_Full_x86_x64.exe
* You can also get the latest Velociraptor for Windows Binary from the
GitHub releases page
    * https://github.com/Velocidex/velociraptor
* Exercise setup scripts if preferred
    * https://gist.github.com/mgreen27/13174916ccfe5fc02b10be13705b2b58


---


<!-- .slide: class="content small-font" -->

## Additional tools

* Memory injection test tool
    * https://github.com/Velocidex/injector/releases
* SelectMyParent.exe
    * https://blog.didierstevens.com/2009/11/22/quickpost-selectmyparent-or-playing-with-the-windows-process-tree/
* Sysinternal tools
    * https://live.sysinternals.com/procdump64.exe
    * https://live.sysinternals.com/sdelete64.exe
    * https://live.sysinternals.com/psexec64.exe
* NTFS Manipulation
    * https://github.com/jschicht/EaTools/raw/master/EaInject64.exe
    * https://github.com/jschicht/EaTools/raw/master/EaQuery64.exe
* WEP Explorer
    * https://github.com/lallousx86/WinTools/tree/master/WEPExplorer
* Process Hacker
    * https://processhacker.sourceforge.io/downloads.php
