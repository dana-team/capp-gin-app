package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dana-team/capp-gin-app/server"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go server.RunServer(quit)

	<-quit
	fmt.Println("Shutting down the server...")
}
