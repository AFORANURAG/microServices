package emailService

import (
	context "context"
	"fmt"
	"log"
	"os"

	gomail "gopkg.in/gomail.v2"

	"strconv"

	authenticationUtilities "github.com/AFORANURAG/microServices/emailService/utilities/authUtilities"
	"github.com/joho/godotenv"
)

type EmailServiceImpl struct {
	senderEmail string
	password    string
	port        int
	smtpServer  string
	secretKey   string
}

func (e *EmailServiceImpl) SendEmail(c context.Context, requestBody *EmailServiceRequest) (*EmailServiceResponse, error) {
	log.Println("chali service")
	log.Printf("Email request is : %v", requestBody)
	log.Printf("Email service struct is :%v", e.password)
	m := gomail.NewMessage()
	m.SetHeader("From", e.senderEmail)
	m.SetHeader("To", requestBody.Email)
	m.SetHeader("Subject", "Verification Mail!")
	token, err := authenticationUtilities.GenerateToken(requestBody.Email, e.secretKey)
	if err != nil {
		log.Printf("Error While Generating JWT %v", err)
	}
	link := fmt.Sprintf("Click on  %s/verify?token=%s to Complete Verification Process.", requestBody.OriginURL, token)
	m.SetBody("text/html", link)

	d := gomail.NewDialer(e.smtpServer, e.port, e.senderEmail, e.password)

	// Send the email to Bob, Cora and Dan.

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error while sending email : %v", err)
		return &EmailServiceResponse{Status: 500, Success: false}, err
	}
	return &EmailServiceResponse{Status: 202, Success: true}, nil
}

func (e *EmailServiceImpl) mustEmbedUnimplementedEmailServiceServer() {

}
func NewEmailServiceProvider() *EmailServiceImpl {
	godotenv.Load()
	PORT, Err := strconv.Atoi(os.Getenv("SMPT_PORT"))
	if Err != nil {
		PORT = 587
	}
	Sender_Email := os.Getenv("SENDER_EMAIL")
	Password := os.Getenv("APP_PASSWORD")
	SmtpServer := os.Getenv("SMTP_SERVER")
	SecretKey := os.Getenv("SECRET_KEY")

	return &EmailServiceImpl{senderEmail: Sender_Email, password: Password, port: PORT, smtpServer: SmtpServer, secretKey: SecretKey}
}
