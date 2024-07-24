package main

import (
	"fmt"
	"log"
	"net"
	"os"

	authenticationService "github.com/AFORANURAG/microServices/authenticationService/services/authService"

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

	authenticationService.RegisterAuthenticationServiceServer(grpcServer, authenticationService.InitializeAuthenticationService(""))
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve grpc server:%v", err)
	}
	fmt.Printf("Server listening on port :%s")
	log.Printf("Server listening on port :%s")
}
