package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/AFORANURAG/microServices/userService/services/userService"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()
	PORT := os.Getenv("SERVER_PORT")
	// HOST := os.Getenv("SERVER_HOST")
	godotenv.Load()
	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	connectionString := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", connectionString)
	if err != nil {
		log.Printf("Failed to listen: %v", err)
	}
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Printf("Error loading .env file")
	}
	grpcServer := grpc.NewServer()

	userService.RegisterUserServiceServer(grpcServer, userService.InitializeUserService(os.Getenv("DSN")))
	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve grpc server:%v", err)
	}
	fmt.Printf("Server listening on port :%s", PORT)
	log.Printf("Server listening on port :%s", PORT)
}
