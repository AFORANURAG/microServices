package main

import (
	"backend/userService/services/userService"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()
	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("SERVER_HOST")
	connectionString := fmt.Sprintf("%s:%s", host, port)
	userServiceConn, err := grpc.Dial(connectionString, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect:%v", err)
	}
	defer userServiceConn.Close()
	client := userService.NewUserServiceClient(userServiceConn)
	res, err := client.GetUserById(context.Background(), &userService.Request{Name: "anurag"})
	if err != nil {
		log.Fatalf("Error while getting user: %v", err)
	}
	log.Println(res.Name)
}
