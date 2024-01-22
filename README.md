# Go Grace

Go graceful process shutdown.

## Install

```text
go get https://github.com/deusexec/go-grace
```

## How To Use

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "syscall"
    "testing"

    "github.com/deusexec/go-grace"
)

type Disposable interface {
    Dispose()
}

type DB struct{}
type Redis struct{}
type RabbitMQ struct{}

func (d *DB) Dispose() {
    fmt.Println("[*] Disposing: DB...")
}

func (d *Redis) Dispose() {
    fmt.Println("[*] Disposing: Redis...")
}

func (d *RabbitMQ) Dispose() {
    fmt.Println("[*] Disposing: RabbitMQ...")
}

func dispose(resources []Disposable) func() {
    return func() {
        fmt.Println("[*] Sever shutting down...")
        for _, resource := range resources {
            resource.Dispose()
        }
        fmt.Println("[*] Done")
    }
}

func main() {
    resources := []Disposable{
        new(DB),
        new(Redis),
        new(RabbitMQ),
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Welcome"))
    })

    go func() {
        log.Println("Server is running...")
        if err := http.ListenAndServe("localhost:8080", nil); err != nil {
            log.Fatal(err)
        }
    }()

    grace.Shutdown(dispose(resources), syscall.SIGINT, syscall.SIGTERM)
}
```

**Terminal 1**

```bash
$ go run .
2025/01/01 15:25:00 Running a server...
2025/01/01 15:25:00 Server is running...
```

**Terminal 2**

```bash
$ kill -s INT <PID>
[*] Sever shutting down...
[*] Disposing: DB...
[*] Disposing: Redis...
[*] Disposing: RabbitMQ...
[*] Done
```
