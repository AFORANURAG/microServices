package authenticationServiceClient

import (
	"fmt"
	"log"
	"os"

	"sync"

	ser "github.com/AFORANURAG/microServices/backend/userService/services/userService"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var authenticationServiceGrpcClient ser.UserServiceClient
var once sync.Once

func AuthenticationServiceClientProvider() ser.UserServiceClient {
	once.Do(
		func() {
			err := godotenv.Load(".env")
			port := os.Getenv("SERVER_PORT")
			host := os.Getenv("SERVER_HOST")
			connectionString := fmt.Sprintf("%s:%s", host, port)
			userServiceConn, err := grpc.Dial(connectionString, grpc.WithInsecure())

			if err != nil {
				log.Fatalf("failed to connect:%v", err)
			}
			defer userServiceConn.Close()
			authenticationServiceGrpcClient = ser.NewUserServiceClient(userServiceConn)
			if err != nil {
				log.Fatalf("Error while getting user: %v", err)
			}
		})
	return authenticationServiceGrpcClient
}
