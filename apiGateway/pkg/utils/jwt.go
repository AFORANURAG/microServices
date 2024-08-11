package utils

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

const (
	INVALID_TOKEN = "Invalid Token"
)

func VerifyJWT(tokenString string, secretKey string) (bool, string) {
token,err:=jwt.Parse(tokenString,func (token *jwt.Token)(interface{},error)  {
	if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
		return nil,fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(secretKey),nil
})

	if err != nil {
		log.Printf("Error While Decoding  JWT :%v\n", err)
		return false, ""
	}

	if token.Valid{
		claims,ok:=token.Claims.(jwt.MapClaims)
		if ok {
			if phoneNumber, ok := claims["phoneNumber"].(string); ok {
				fmt.Printf("PhoneNumber is : %v\n", phoneNumber)
				return true, phoneNumber
			}
			
		}
		if !ok{
				return false,""
		}
	}
	
	return false, INVALID_TOKEN
}