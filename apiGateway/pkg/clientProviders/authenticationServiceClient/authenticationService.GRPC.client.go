package authenticationServiceClient

import (
	"fmt"
	"log"
	"os"

	"sync"

	ser "github.com/AFORANURAG/microServices/authenticationService/services/authService"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var authenticationServiceGrpcClient ser.AuthenticationServiceClient
var once sync.Once

func AuthenticationServiceClientProvider() ser.AuthenticationServiceClient {
	once.Do(
		func() {
			err := godotenv.Load(".env")
			port := os.Getenv("AUTHSERVICE_PORT")
			host := os.Getenv("SERVER_HOST")
			fmt.Printf("port is : %v", port)
			connectionString := fmt.Sprintf("%s:%s", host, port)
			authServiceConn, err := grpc.Dial(connectionString, grpc.WithInsecure())

			if err != nil {
				log.Printf("failed to connect:%v", err)
			}
			authenticationServiceGrpcClient = ser.NewAuthenticationServiceClient(authServiceConn)
			if err != nil {
				log.Printf("Error while getting user: %v", err)
			}
		})
	return authenticationServiceGrpcClient
}
