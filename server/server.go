package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("Error finding a port value")
	}

	r := gin.Default()
	r.Use(cors.Default())

	srv := &http.Server{
		Addr:         fmt.Sprintf(`:%v`, port),
		Handler:      r,
	}

	go func() {
		log.Printf("Server is listening on port :%v", port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server ListenAndServe: %v\n", err)
		}
	}()

	quit := make(chan os.Signal,1)
	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	<-quit
	log.Println("Server shutting down...")

	gracefullCtx, ShutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ShutdownCancel()

	if err := srv.Shutdown(gracefullCtx); err != nil {
		log.Fatal("Server force shutdown:", err)
	}
	log.Println("goodbye...")
}
