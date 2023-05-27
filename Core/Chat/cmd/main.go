package main

import (
	"chat/infra/application"
	"chat/logger"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	app := application.CreateApp(myLogger)
	err := app.Begin()
	if err != nil {
		myLogger.Sugar().Fatalf("app not started: %s", err.Error())
	}

	<-ctx.Done()

	stopCtx, stopCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer stopCancel()

	err = app.End(stopCtx)
	if err != nil {
		myLogger.Sugar().Fatalf("error in app.Stop: %s", err.Error())
	}
}
