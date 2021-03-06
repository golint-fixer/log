# log [![Build Status](https://travis-ci.org/toashd/log.svg)](https://travis-ci.org/toashd/log)

log provides a small set of functions for extended logging. Log messages have different levels (Debug, Info, Warning, Error), are colorized per level and are prepended with an optional timestamp. Only messages greater or equal to the set log level will be logged to standard out/error.

## Installation

Standard `go get`:

```
$ go get github.com/toashd/log
```

## Usage & Example

For usage and examples see the [Godoc](http://godoc.org/github.com/toashd/log).

![alt log](https://cloud.githubusercontent.com/assets/1436166/12043339/a920b3f4-ae84-11e5-87da-ab70e5ce1fe0.png)

Usage is easy enough:

```go

log.Debug("debug message")
log.Info("info message")
log.Warn("warning message")
log.Error("error message")

log.Errorf("%s error message", "formatted")

```

Color and timestamp are enabled by default but can be easily disabled.

## License
MIT
