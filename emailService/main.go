package main

import (
	"fmt"
	"log"
	"net"
	"os"

	emailService "github.com/AFORANURAG/microServices/backend/emailService/services"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	connectionString := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", connectionString)
	if err != nil {
		log.Printf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	emailService.RegisterEmailServiceServer(grpcServer, emailService.InitializeAuthenticationService())
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve grpc server:%v", err)
	}
	log.Printf("Server listening on port :%s", port)
}
