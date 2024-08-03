package main

import (
	"fmt"
	"log"
	"net"
	"os"

	authenticationService "github.com/AFORANURAG/microServices/authenticationService/services/authService"
	queueservicetypes "github.com/AFORANURAG/microServices/authenticationService/types/queueServiceTypes"
	userservicetypes "github.com/AFORANURAG/microServices/authenticationService/types/userServiceTypes"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	// HOST := os.Getenv("SERVER_HOST")
	godotenv.Load()
	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	connectionString := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", connectionString)
	if err != nil {
		log.Printf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	authenticationService.RegisterAuthenticationServiceServer(grpcServer, authenticationService.InitializeAuthenticationService(userservicetypes.UserServicePhrase(os.Getenv("DSN")),queueservicetypes.QueueServicePhrase(os.Getenv("AMQP_CLOUD_URL"))))
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve grpc server:%v", err)
	}
}
