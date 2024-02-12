package emailclient

import (
	"fmt"
	"log"
	"os"

	"sync"

	emailService "github.com/AFORANURAG/microServices/backend/emailService/services"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var emailServiceGRPCClient emailService.EmailServiceClient
var once sync.Once

func NewEmailServiceServiceClientProvider() emailService.EmailServiceClient {
	once.Do(
		func() {
			godotenv.Load()
			port := os.Getenv("EMAIL_SERVICE_PORT")
			host := os.Getenv("SERVER_HOST")
			connectionString := fmt.Sprintf("%s:%s", host, port)
			userServiceConn, err := grpc.Dial(connectionString, grpc.WithInsecure())

			if err != nil {
				log.Printf("failed to connect:%v", err)
			}
			emailServiceGRPCClient = emailService.NewEmailServiceClient(userServiceConn)
			if err != nil {
				log.Printf("Error while getting user: %v", err)
			}
		})
	return emailServiceGRPCClient
}
