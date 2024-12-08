# libcmon
[![Go Reference](https://pkg.go.dev/badge/pindorama.net.br/libcmon.svg)](https://pkg.go.dev/pindorama.net.br/libcmon)

## <img src="https://github.com/user-attachments/assets/ea106675-94e2-4a67-bbc6-c992e9159ac8" height=40> **WARNING**: This library is in W.I.P. state, the APIs may change during the development.

## What is it?

C'mon (/kəˈmɔn/) is a general library for the Go Programming Language.  
Its spirit is the same as [Heirloom Toolchest's
``libcommon/``](https://github.com/Projeto-Pindorama/heirloom-ng/tree/master/libcommon).  
Still work in progress. Meant to be used in many projects, from
[pacote](https://github.com/Projeto-Pindorama/pacote) to the L.E.``mount``
rewrite in Go.

It provides the following functionality:

* General file-to-``[]byte`` functions and ``str(n)cpy``(3) "replicas" available
  at ``bass/``. The primitive is ``bass.Walk()``, the other ones are based on it;
* Functions for collecting disk information at ``disks/``, also planned to write
  on disks in the future;
* General wrappers at ``porcelana/``.

## Licence

The
[BSD 3 Clause licence](https://github.com/Projeto-Pindorama/libcmon?tab=License-1-ov-file),
but the authors also request that this project is used for good, not evil.
