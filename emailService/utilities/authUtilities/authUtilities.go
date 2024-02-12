package authenticationUtilities

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type DataToEmbed struct {
	Email string `json:"email"`
	/*
		all other standards claims to be encoded in the token
	*/
	jwt.StandardClaims
}

func GenerateToken(email string, secretKey string) (string, error) {
	claims := DataToEmbed{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
			IssuedAt:  time.Now().Unix(),
		},
	}
	// Create the token with the claims and the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatalf("Error While Generating Token %v", err)
	}
	return tokenString, nil
}
