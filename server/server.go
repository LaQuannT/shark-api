package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/LaQuannT/shark-api/database"
)

func main() {
	var Logger *logrus.Logger
	Logger.SetFormatter(&logrus.JSONFormatter{FieldMap: logrus.FieldMap{
		logrus.FieldKeyTime:  "@timestamp",
		logrus.FieldKeyLevel: "@level",
		logrus.FieldKeyMsg:   "@message",
	}})

	err := godotenv.Load()
	if err != nil {
		Logger.Fatal("ENV: Failed to load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		Logger.Fatal("ENV: 'PORT' variable not set")
	}

	connstr := os.Getenv("DB_CONNSTR")
	if connstr == "" {
		Logger.Fatal("ENV: 'DB_CONNSTR' variable not set")
	}

	DB, err := database.Init(connstr)
	if err != nil {
		Logger.Fatalf("Database error: %v\n", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	srv := &http.Server{
		Addr:    fmt.Sprintf(`:%v`, port),
		Handler: r,
	}

	go func() {
		Logger.Infof("Server is listening on port :%v", port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			Logger.Fatalf("Server ListenAndServe: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	<-quit
	Logger.Info("Server shutting down...")

	gracefullCtx, ShutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ShutdownCancel()

	if err := srv.Shutdown(gracefullCtx); err != nil {
		Logger.Fatalf("Server force shutdown:", err)
	}
	Logger.Info("Server exiting")
}
