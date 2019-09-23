# hardwareid provides support for reading the unique hardware address of most host OS's 
![Image of Gopher 47](logo.png)

… because sometimes you just need to reliably identify your hardwares.

[![GoDoc](https://godoc.org/github.com/sandipmavani/hardwareid?status.svg)](https://godoc.org/github.com/sandipmavani/hardwareid) [![Go Report Card](https://goreportcard.com/badge/github.com/sandipmavani/hardwareid)](https://goreportcard.com/report/github.com/sandipmavani/hardwareid)

## Main Features

* Cross-Platform (tested on Win7+, Debian 8+, Ubuntu 14.04+, OS X 10.6+, FreeBSD 11+)
* No admin privileges required
* Hardware independent (no usage of MAC, BIOS or CPU — those are too unreliable, especially in a VM environment)
* IDs are unique<sup>[1](#unique-key-reliability)</sup> to the installed OS

## Installation

Get the library with

```bash
go get github.com/sandipmavani/hardwareid
```

You can also add the cli app directly to your `$GOPATH/bin` with

```bash
go get github.com/sandipmavani/hardwareid/cmd/hardwareid
```

## Usage

```golang
package main

import (
  "fmt"
  "log"
  "github.com/sandipmavani/hardwareid"
)

func main() {
  id, err := hardwareid.ID()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(id)
}
```

Or even better, use securely hashed hardware IDs:

```golang
package main

import (
  "fmt"
  "log"
  "github.com/sandipmavani/hardwareid"
)

func main() {
  id, err := hardwareid.ProtectedID("myAppName")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(id)
}
```

### Function: ID() (string, error)

Returns original hardware address as a `string` like `MM:MM:MM:SS:SS:SS`.

### Function: ProtectedID(appID string) (string, error)

Returns hashed version of the hardware ID as a `string`. The hash is generated in a cryptographically secure way, using a fixed, application-specific key (calculates HMAC-SHA256 of the app ID, keyed by the hardware ID).

## What you get

This package returns the hardware address mac address of your system.


Do something along these lines:

```golang
package main

import (
  "crypto/hmac"
  "crypto/sha256"
  "fmt"
  "github.com/sandipmavani/hardwareid"
)

const appKey = "WowSuchNiceApp"

func main() {
  id, _ := hardwareid.ID()
  fmt.Println(protect(appKey, id))
  // Output: dbabdb7baa54845f9bec96e2e8a87be2d01794c66fdebac3df7edd857f3d9f97
}

func protect(appID, id string) string {
  mac := hmac.New(sha256.New, []byte(id))
  mac.Write([]byte(appID))
  return fmt.Sprintf("%x", mac.Sum(nil))
}
```

Or simply use the convenience API call:

```golang
hashedID, err := hardwareid.ProtectedID("myAppName")
```


## Credits

The Go gopher was created by [Denis Brodbeck](https://github.com/sandipmavani) with [gopherize.me](https://gopherize.me/), based on original artwork from [Renee French](http://reneefrench.blogspot.com/).

## License

The MIT License (MIT) — [Denis Brodbeck](https://github.com/sandipmavani). Please have a look at the [LICENSE.md](LICENSE.md) for more details.
