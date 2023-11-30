package server

import (
	// "context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/udodinho/job-app/pkg/middleware"
)

func (s *Server) Start() {
	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	if PORT == ":" {
		PORT = ":8080"
	}

	idleConnsClosed := make(chan struct{})
	
	wait := make(chan os.Signal, 1)
	// s.SetupRouter()

	// r := Server{
	// 	Db: configs.DB,
	// }

	server := fiber.New()

	// Middleware
	middleware.FiberMiddleware(server)

	// Routes
	s.SetupRouter(server)
	s.NotFoundRoute(server)

	go func() {
		log.Printf("Server started on http://localhost%s\n", PORT)
		if err := server.Listen(PORT); err != nil && err != http.ErrServerClosed {
			fmt.Printf("%v\n", err)
		}

		close(idleConnsClosed)
	}()

	// Ctrl + C on the terminal for example sends a terminate signal.
	signal.Notify(wait, os.Interrupt)
	// Blocks until a signal is sent to the channel.
	<-wait

	log.Printf("Shutting down the server...")
	time.Sleep(time.Second)
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// Gracefully shutdown the server when receiving the TERM or KILL signals.
	if err := server.Shutdown(); err != nil {
		log.Fatal("Shutting down forcefully", err)
	}
	log.Println("Server Shut Down")
}


