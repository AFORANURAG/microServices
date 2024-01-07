package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	fmt.Println("GraceFull shutdown enabled")
	idleConnsClosed := make(chan struct{})
	go func() {
		// create a buffered channel to receive os signal
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		// interupptions for example cntr+c etc.
		<-sigint
		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}
		close(idleConnsClosed)
	}()
	fiber_connection_URL, _ := ConnectionURLBuilder("fiber")
	if err := a.Listen(fiber_connection_URL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

func StartServer(a *fiber.App) {
	// Build Fiber connection URL.
	fiberConnURL, _ := ConnectionURLBuilder("fiber")

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
