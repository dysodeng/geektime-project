package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Server http server
func Server(httpPort string) *http.Server {
	// http server
	server := &http.Server{
		Addr:    ":" + httpPort,
		Handler: Router(),
	}

	// error logger
	logFilename := "var/gin.log"
	errLogFile, _ := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	gin.DefaultErrorWriter = io.MultiWriter(errLogFile, os.Stderr)

	defer func() {
		if err := recover(); err != nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				log.Fatal("http server shutdown error:", err)
			}
		}
	}()

	log.Printf("listening and serving HTTP on: :%s\n", httpPort)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("%+v", errors.Wrap(err, "http server error"))
		}
	}()

	return server
}
