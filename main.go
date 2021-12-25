package main

import (
	"context"
	"github.com/dysodeng/project/internal/provider"
	"github.com/dysodeng/project/server/grpc"
	"github.com/dysodeng/project/server/http"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	httpPort := "8080"
	rpcPort := "9000"

	// 服务容器
	provider.ServiceProvider()

	// grpc server
	grpcServer := grpc.Server(localIp(), rpcPort)

	// http server
	httpServer := http.Server(httpPort)

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("shutdown http server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("http server shutdown error:", err)
	}
	_ = grpcServer.Stop()
}

// localIp 本机IP
func localIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
