# birdid

Bird + Cupid: Find the first interaction between two Twitter users

[![hydrun CI](https://github.com/pojntfx/birdid/actions/workflows/hydrun.yaml/badge.svg)](https://github.com/pojntfx/birdid/actions/workflows/hydrun.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/pojntfx/birdid.svg)](https://pkg.go.dev/github.com/pojntfx/birdid)
[![Matrix](https://img.shields.io/matrix/birdid:matrix.org)](https://matrix.to/#/#birdid:matrix.org?via=matrix.org)
[![Binary Downloads](https://img.shields.io/github/downloads/pojntfx/birdid/total?label=binary%20downloads)](https://github.com/pojntfx/birdid/releases)

## Overview

`birdid` finds the first connection (tweet to, mention etc.) between two Twitter users.

## Installation

Static binaries are available on [GitHub releases](https://github.com/pojntfx/birdid/releases).

On Linux, you can install them like so:

```shell
$ curl -L -o /tmp/birdid "https://github.com/pojntfx/birdid/releases/latest/download/birdid.linux-$(uname -m)"
$ sudo install /tmp/birdid /usr/local/bin
```

On macOS, you can use the following:

```shell
$ curl -L -o /tmp/birdid "https://github.com/pojntfx/birdid/releases/latest/download/birdid.darwin-$(uname -m)"
$ sudo install /tmp/birdid /usr/local/bin
```

On Windows, the following should work (using PowerShell as administrator):

```shell
PS> Invoke-WebRequest https://github.com/pojntfx/birdid/releases/latest/download/birdid.windows-x86_64.exe -OutFile \Windows\System32\birdid.exe
```

You can find binaries for more operating systems and architectures on [GitHub releases](https://github.com/pojntfx/birdid/releases).

## Usage

To find the earliest contact between two users, run the following (you can get the client ID and client secret from the [Twitter developer portal](https://developer.twitter.com/en/portal/dashboard)). Be sure to adjust the limit to something higher (i.e. 5000):

```shell
$ birdid --client-id=your-client-id --client-secret=your-client-secret --candidate-one=jagger27 --candidate-two=pojntfx --limit=10
Earliest tweet from jagger27 to pojntfx: ID 1502328878479683590 at Fri Mar 11 17:01:31 +0000 2022 with URL https://twitter.com/jagger27/status/1502328878479683590 and text RT @pojntfx: Want a good argument for massive investments in public rail infrastructure? Look at how aid is being delivered to Ukraine.
Earliest tweet from pojntfx to jagger27: ID 1502272146147622912 at Fri Mar 11 13:16:05 +0000 2022 with URL https://twitter.com/pojntfx/status/1502272146147622912 and text @jagger27 🥺
```

Be sure to check out the [reference](#reference) for more information.

## Reference

```bash
$ birdid --help
Usage of birdid:
  -candidate-one string
    	First candidate's Twitter handle
  -candidate-two string
    	Second candidate's Twitter handle
  -client-id string
    	Twitter API client ID (can also be set using the CLIENT_ID env variable)
  -client-secret string
    	Twitter API client secret (can also be set using the CLIENT_SECRET env variable)
  -limit int
    	Maximum amount of events to look back in timeline (default 1000)
  -verbose
    	Enable verbose logging
```

## Acknowledgements

- This project would not have been possible were it not for [@dghubble](https://github.com/dghubble)'s [go-twitter](https://github.com/dghubble/go-twitter) package; be sure to check it out too!
- All the rest of the authors who worked on the dependencies used! Thanks a lot!

## Contributing

To contribute, please use the [GitHub flow](https://guides.github.com/introduction/flow/) and follow our [Code of Conduct](./CODE_OF_CONDUCT.md).

To build birdid locally, run:

```shell
$ git clone https://github.com/pojntfx/birdid.git
$ cd birdid
$ make depend
$ make
$ out/birdid
```

Have any questions or need help? Chat with us [on Matrix](https://matrix.to/#/#birdid:matrix.org?via=matrix.org)!

## License

birdid (c) 2022 Felicitas Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
