package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/huprince/quick-gin/config"
)

// server run
func Run(router *gin.Engine) {
	server := &http.Server{
		Addr: ":" + config.GetEnv().ServerConfig.ServerPort,
		Handler: router,
	}

	go func ()  {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Listen: %s \n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown: %s \n", err)
	}
	pid := fmt.Sprintf("%d", os.Getpid())
	_, openErr := os.OpenFile("pid", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if openErr == nil {
		_ = ioutil.WriteFile("pid", []byte(pid), 0)
	}

}