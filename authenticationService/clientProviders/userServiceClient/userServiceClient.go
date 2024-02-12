package userServiceClient

import (
	"fmt"
	"log"
	"os"

	"sync"

	ser "github.com/AFORANURAG/microServices/backend/userService/services/userService"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var userServiceGrpcClient ser.UserServiceClient
var once sync.Once

func NewUserServiceServiceClientProvider() ser.UserServiceClient {
	once.Do(
		func() {
			godotenv.Load()
			port := os.Getenv("USER_SERVICE_PORT")
			host := os.Getenv("SERVER_HOST")
			connectionString := fmt.Sprintf("%s:%s", host, port)
			userServiceConn, err := grpc.Dial(connectionString, grpc.WithInsecure())

			if err != nil {
				log.Printf("failed to connect:%v", err)
			}
			userServiceGrpcClient = ser.NewUserServiceClient(userServiceConn)
			if err != nil {
				log.Printf("Error while getting user: %v", err)
			}
		})
	return userServiceGrpcClient
}
