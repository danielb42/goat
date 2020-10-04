# goat

![Tag](https://img.shields.io/github/v/tag/danielb42/goat)
![Go Version](https://img.shields.io/github/go-mod/go-version/danielb42/goat)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/danielb42/goat)](https://pkg.go.dev/github.com/danielb42/goat)
[![Go Report Card](https://goreportcard.com/badge/github.com/danielb42/goat)](https://goreportcard.com/report/github.com/danielb42/goat)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Golang connector for `at(1)`.  

Schedule external command executions, powered by the `at(1)`-utility:

```golang
// let /hello/world.sh be run at <execTime>
execTime := time.Date(<in the future>)
jobID, err := at.AddJob("/hello/world.sh", execTime)

// changed your mind?
at.RemoveJob(jobID)
```

## Prerequisites

`at(1)` must be present on your system. If it's not, try `apt install at` or `yum install at` or `pacman -S at` or `apk add at` according to your linux flavor.
