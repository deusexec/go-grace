package grace

import (
	"fmt"
	"log"
	"net/http"
	"syscall"
	"testing"
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

func Test_Shutdown(t *testing.T) {
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

	Shutdown(dispose(resources), syscall.SIGINT, syscall.SIGTERM)
}
