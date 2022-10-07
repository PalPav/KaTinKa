package main

import (
	_ "github.com/lib/pq"
	"katinka/http/server"
	"katinka/services/container"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	log.Println("Hello i'm KaTinKa!")

	// Run http server
	go server.Run()

	sig := <-cancelChan
	// shutdown other goroutines gracefully
	// close other resources

	log.Printf("Caught %v", sig)
	log.Printf("Bye bye ) Cya <3")
	if container.DB() != nil {
		container.DB().Close()
	}
}
