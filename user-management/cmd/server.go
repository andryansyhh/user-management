package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"net"

	"google.golang.org/grpc"
)

func Run() {
	server := grpc.NewServer()
	deps := InitDependencies(server)

	listener, err := net.Listen("tcp", ":"+deps.Config.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", deps.Config.GRPCPort, err)
	}

	log.Printf("Connected to DB on %s:%s", deps.Config.DBHost, deps.Config.DBPort)

	go func() {
		log.Printf("gRPC server running on %s", deps.Config.GRPCPort)
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	gracefulShutdown(server)
}

func gracefulShutdown(server *grpc.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down gRPC server...")
	server.GracefulStop()
	log.Println("Server stopped.")
}
