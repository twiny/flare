# Flare 🚀

A simple, lightweight signaling mechanism in Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/twiny/flare)](https://goreportcard.com/report/github.com/twiny/flare)
[![GoDoc](https://godoc.org/github.com/twiny/flare?status.svg)](https://pkg.go.dev/github.com/twiny/flare)

---

## Introduction

Flare provides an alternative to the standard context package in Go for signaling goroutines, without carrying cancellation info or other values. It's a minimalistic tool designed to provide a straightforward mechanism to signal goroutines to stop their work.

## Installation

```bash
go get github.com/twiny/flare
```

## Usage

### Basic Usage

```go
package main

import (
 "github.com/twiny/flare"
)

func main() {
 n := flare.NewNotifier()

 // Spawn a goroutine
 go func() {
  select {
  case <-n.Hold():
   return
  default:
   // Do some work
  }
 }()

 // Signal the goroutine to stop
 n.Signal()
}
```

### Integration with context.Context

```go
func main() {
 n, cancel := flare.NewNotifierWithCancel(context.Background())
 // Canceling the context will also signal the flare Notifier.
 defer cancel()

 // Spawn a goroutine
 go func() {
  select {
  case <-n.Hold():
   return
  default:
   // Do some work
  }
 }()

 // Signal the goroutine to stop
 n.Signal()
}
```

## Wiki

More documentation can be found in the [wiki](https://github.com/twiny/flare/wiki).

## Bugs

Bugs or suggestions? Please visit the [issue tracker](https://github.com/twiny/flare/issues).
