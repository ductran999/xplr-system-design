package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ductran999/xplr-system-design/client"
	"github.com/ductran999/xplr-system-design/server"
)

func main() {
	fmt.Println("Welcome to Go Net Cat")
	fmt.Println("1. Start Server")
	fmt.Println("2. Start Client")
	fmt.Print("=> Enter Your option:")

	var option int
	_, err := fmt.Scanln(&option)
	if err != nil {
		log.Fatal("invalid input:", err.Error())
	}

	switch option {
	case 1:
		startServer()
	case 2:
		startClient()
	default:
		slog.Error("Invalid option. Please enter 1 or 2.")
	}
}

func startServer() {
	srv := server.NewServer()
	if err := srv.Open(); err != nil {
		log.Fatal("failed to open server", err.Error())
	}
	slog.Info("server listening on port 8080")

	go func() {
		if err := srv.Serve(); err != nil && !errors.Is(err, net.ErrClosed) {
			slog.Error("server got error:", "error", err.Error())
		}
	}()

	gracefulShutdown(srv.Close)
}

func startClient() {
	c := client.NewClient()
	if err := c.Dial(); err != nil {
		log.Fatal("failed to dial to host", err.Error())
	}
	slog.Info("server accept your call")

	go func() {
		if err := c.Send(); err != nil &&
			!errors.Is(err, net.ErrClosed) &&
			!errors.Is(err, io.EOF) {
			slog.Error(err.Error())
		}
	}()

	gracefulShutdown(c.Close)
}

func gracefulShutdown(close func() error) {
	// wait for Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	slog.Info("Shutting down...")

	err := close()
	if err != nil {
		slog.Error("failed to shut down:", "error", err.Error())
	} else {
		slog.Info("shutdown cleanly")
	}
}
