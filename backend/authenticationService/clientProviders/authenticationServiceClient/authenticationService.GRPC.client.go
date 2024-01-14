package userServiceClient

import (
	"fmt"
	"log"

	"sync"

	ser "github.com/AFORANURAG/microServices/backend/userService/services/userService"
	"google.golang.org/grpc"
)

var userServiceGrpcClient ser.UserServiceClient
var once sync.Once

func UserServiceServiceClientProvider(port string, host string) ser.UserServiceClient {
	once.Do(
		func() {

			connectionString := fmt.Sprintf("%s:%s", host, port)
			userServiceConn, err := grpc.Dial(connectionString, grpc.WithInsecure())

			if err != nil {
				log.Fatalf("failed to connect:%v", err)
			}
			defer userServiceConn.Close()
			userServiceGrpcClient = ser.NewUserServiceClient(userServiceConn)
			if err != nil {
				log.Fatalf("Error while getting user: %v", err)
			}
		})
	return userServiceGrpcClient
}
