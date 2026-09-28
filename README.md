<h1 align="center">
  go-readingtime
  <br>
</h1>
<h4 align="center">Estimate how long it takes to read a text</h4>
<h6 align="center"> Coded with 💙 by edoardottt </h6>
<p align="center">
  <a href="https://github.com/edoardottt/go-readingtime/actions">
      <img src="https://github.com/edoardottt/go-readingtime/actions/workflows/go.yml/badge.svg" alt="go action">
  </a>
<p align="center">
  <a href="#install-">Install</a> •
  <a href="#usage-">Usage</a> •
  <a href="#changelog-">Changelog</a> •
  <a href="#contributing-">Contributing</a> •
  <a href="#license-">License</a>
</p>

Install 📡
----------

```console
go install github.com/edoardottt/go-readingtime/cmd/readt@latest
```

Usage 💡
----------

CLI tool

```console
readt <filepath>
```

Golang module

```go
package main

import (
 "fmt"

 reading "github.com/edoardottt/go-readingtime"
)

func main() {
 s := reading.RawEstimate(`Lorem ipsum dolor sit amet, consectetur...`)
 fmt.Println(s) // 130

 d := reading.Estimate(`Lorem ipsum dolor sit amet, consectetur...`)
 fmt.Println(d) // 2m10s

 h := reading.HumanEstimate(`Lorem ipsum dolor sit amet, consectetur...`)
 fmt.Println(h) // 2 minutes
}
```

Read the [`package documentation here`](https://pkg.go.dev/github.com/edoardottt/go-readingtime).

Changelog 📌
-------

Detailed changes for each release are documented in the [release notes](https://github.com/edoardottt/go-readingtime/releases).

Contributing 🛠
-------

Just open an [issue](https://github.com/edoardottt/go-readingtime/issues) / [pull request](https://github.com/edoardottt/go-readingtime/pulls).

Before opening a pull request, download [golangci-lint](https://golangci-lint.run/usage/install/) and run

```bash
golangci-lint run
```

If there aren't errors, go ahead :)

License 📝
-------

This repository is under [MIT License](https://github.com/edoardottt/go-readingtime/blob/main/LICENSE).  
[edoardottt.com](https://edoardottt.com/) to contact me.
