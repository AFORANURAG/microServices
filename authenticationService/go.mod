module github.com/AFORANURAG/microServices/authenticationService

go 1.21

toolchain go1.22.5

require (
	github.com/AFORANURAG/microServices/emailService v0.0.0-20240726154403-12ea93e1b44c
	github.com/AFORANURAG/microServices/userService v0.0.0-20240725071507-461f0f420531
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.8.1
	github.com/google/wire v0.6.0
	github.com/joho/godotenv v1.5.1
	github.com/rabbitmq/amqp091-go v1.10.0
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240730163845-b1a4ccb954bf // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
)
