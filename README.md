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
import "github.com/twiny/flare"

n := flare.New()

// Spawn a goroutine
go func() {
    select {
    case <-n.Done():
        return
    default:
        // Do some work
    }
}()

// Signal the goroutine to stop
n.Cancel()
```

### Integration with context.Context

```go
import (
    "context"
    "github.com/twiny/flare"
)

ctx, cancel := context.WithCancel(context.Background())

n := flare.NewWithContext(ctx)

// Goroutines can listen to n.Done() for exit signals.

// Canceling the context will also signal the flare Notifier.
cancel()
```