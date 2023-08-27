package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleRoutes(addr string) {
	app := gin.Default()

	app.GET("/:query", handleGetPron)

	srv := http.Server{
		Addr:    addr,
		Handler: app,
	}

	go srv.ListenAndServe()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shuting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v\n", err)
	}

	select {
	case <-ctx.Done():
		log.Println("server shutdown timeout")
	}

	log.Println("change da world, my final message. goodbye!")
}
