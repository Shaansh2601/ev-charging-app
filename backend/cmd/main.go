package main

import (
	backend "backend"
	"backend/pkg/handler"
	"backend/pkg/repository"
	"backend/pkg/service"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewMySQLDB(repository.Config{
		Host:     os.Getenv("DB_Host"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_Username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DBName"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s ", err)
	}
	repos := repository.NewRepository(db)

	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(backend.Server)
	go func() {
		if err := srv.Run(os.Getenv("Server_PORT"), handlers.InitRoutes(os.Getenv("GIN_MODE"))); err != nil {
			logrus.Fatalf("error occured while running http server:  %s", err.Error())
		}
	}()

	logrus.Print("Server Started")

	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	logrus.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown: ", err)
	}

	logrus.Println("Server exiting")
}
