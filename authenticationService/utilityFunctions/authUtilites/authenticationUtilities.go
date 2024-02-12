package authenticationUtilities

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	INVALID_TOKEN = "Invalid Token"
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

func VerifyJWT(tokenString string, secretKey string) (bool, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Printf("Error While Decoding  JWT :%v\n", err)
		return false, ""
	}
	fmt.Printf("Token is: %s\n", token.Valid)

	if token.Valid {
		fmt.Printf("Hello world ")
		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Printf("Ok ok : %v\n", ok)
		if ok {
			if email, ok := claims["email"].(string); ok {
				fmt.Printf("Email is : %v\n", email)
				return true, email
			}
		}
		if !ok {
			return false, ""
		}
	}
	return false, INVALID_TOKEN
}
