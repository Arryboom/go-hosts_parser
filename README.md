# HostsParser

a simple hosts parser support cross-platform.

[![Documentation](https://godoc.org/github.com/arryboom/go-hostsparser?status.svg)](https://godoc.org/github.com/arryboom/go-hostsparser)
[![Build Status](https://travis-ci.org/arryboom/go-hostsparser.svg)](https://travis-ci.org/arryboom/go-hostsparser)
[![Report Card](https://goreportcard.com/badge/github.com/arryboom/go-hostsparser)](https://goreportcard.com/report/arryboom/go-hostsparser)

### The /etc/hosts parsing and resolver library for golang

HostsParser is a go (golang) package for parsing hosts files and performing hostname ->IP lookups which are file-based (e.g. via /etc/hosts).

This can be helpful to determine if there is a prettier (or known) hostname
available for an IP address.

Naturally, the (obvious) tradeoff/downside is that this only
works in cases where the IP mapping exists in the hostsfile.

## Get it

    go get -u https://github.com/arryboom/go-hostsparser

## Usage

```go
package main

import (
    "fmt"

    "github.com/arryboom/go-hostsparser"
)

func main() {
    resx, err := NameLookup("example.com")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v maps example.com in the following files: %v", res, hostsfile.HostsPath)
}
```

Output:

    127.0.0.1 maps example.com in the following files: c:\windows\system32\drivers\etc\hosts

## Supported Operating Systems

Tested and verified working on:

* Linux
* Mac OS X
* Windows

It *should* also support Plan9, though admittedly He don't have a plan9
installation available to verify that one ;)

## Unit-tests

Running the unit-tests is straightforward and standard:

    go test


## License

Permissive MIT license.

## Thanks

- jaytaylor  
https://github.com/jaytaylor/go-hostsfile/
