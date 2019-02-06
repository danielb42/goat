# goat
[![GoDoc](https://godoc.org/github.com/danielb42/goat?status.svg)](https://godoc.org/github.com/danielb42/goat) 
[![Go Report Card](https://goreportcard.com/badge/github.com/danielb42/goat)](https://goreportcard.com/report/github.com/danielb42/goat) 
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)  

Golang connector for `at(1)`.  

Schedule external command executions, powered by the `at(1)`-utility:

```golang
// let /hello/world.sh be run at <execTime>
execTime := time.Date(<in the future>)
jobID, err := at.AddJob("/hello/world.sh", execTime)

// changed your mind?
at.RemoveJob(jobID)
```

### Prerequisites
`at(1)` must be present on your system. If it's not, try `apt install at` or `yum install at` or `pacman -S at` or `apk add at` according to your linux flavor.
