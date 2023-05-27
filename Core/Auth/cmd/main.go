package main

import (
	"auth/infra/application"
	"auth/logger"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func initMain() logger.Logger {
	gin.SetMode(gin.DebugMode)
	var err error
	var myLogger logger.Logger
	myLogger, err = logger.New()
	if err != nil {
		log.Fatalf("logger initialization failed: %s", err.Error())
	}
	return myLogger
}

func main() {
	var myLogger = initMain()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	app := application.New(myLogger)
	err := app.Start()
	if err != nil {
		myLogger.Sugar().Fatalf("app not started: %s", err.Error())
	}

	<-ctx.Done()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopCancel()

	err = app.Stop(stopCtx)
	if err != nil {
		myLogger.Sugar().Fatalf("error in app.Stop: %s", err.Error())
	}
}
