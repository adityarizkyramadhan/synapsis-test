package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	httpHandler "github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/handler/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env")
	}
	router := gin.New()

	userGroup := router.Group("/user")

	userHandler := httpHandler.NewUserRoutes()

	if err := userHandler.Init(userGroup); err != nil {
		log.Fatalf("failed to initialize user handler: %v", err)
	}

	authorGroup := router.Group("/author")
	authorHandler := httpHandler.NewAuthorRoutes()

	if err := authorHandler.Init(authorGroup); err != nil {
		log.Fatalf("failed to initialize author handler: %v", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server started on port %s\n", os.Getenv("PORT"))

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Use select to block the main goroutine and keep cron running
	select {
	case <-quit:
		log.Println("Shutting down server...")
	}

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
