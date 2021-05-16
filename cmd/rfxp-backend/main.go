package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/stebunting/rfxp-backend/server"
)

func main() {
	godotenv.Load()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_DSN"),
		Environment: os.Getenv("SENTRY_ENV"),
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	server := server.NewServer(server.ServerConfig{
		Port: port,
	})

	server.Start()
	server.Stop()
}
