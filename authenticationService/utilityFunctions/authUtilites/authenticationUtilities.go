package authenticationUtilities

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	INVALID_TOKEN = "Invalid Token"
)

type TwoFactorVerifyOTPResponse struct {
	Status string
	Details string
}
type DataToEmbed struct {
	PhoneNumber string `json:"phoneNumber"`
	/*
		all other standards claims to be encoded in the token
	*/
	jwt.StandardClaims
}

func GenerateToken(phoneNumber string, secretKey string) (string, error) {
	claims := DataToEmbed{
		PhoneNumber: phoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 1 hour
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

func  Verify2factorOTP(url string,PhoneNumber string,otp int) bool {
	var completeURL string=fmt.Sprintf("%sVERIFY3/%s/%d",url,PhoneNumber,otp)
	 	resp,err:=http.Get(completeURL)
		 if err != nil {
        log.Printf("Failed to send GET request: %v", err)
		return false 
    }
    defer resp.Body.Close()

	 body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Failed to read response body: %v", err)
		return false
    }

    // Print the response status and body
    fmt.Printf("Response Status: %s\n", resp.Status)
    fmt.Printf("Response Body: %s\n", string(body))

	var apiResponse TwoFactorVerifyOTPResponse;
	// pointer to apiResponse because it is gonna mutate the original predeclared apiResponse of type TwoFactorVerifyOTPResponse
	err=json.Unmarshal(body,&apiResponse)
	
	if err != nil {
        fmt.Printf("Failed to parse JSON response: %v", err)
		return false
    }
	if apiResponse.Details== "OTP Matched"{
		return true
	}


    // Print the parsed response
    fmt.Printf("Parsed Response: %+v\n", apiResponse)
	
return false
}