package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const (
	INVALID_TOKEN = "Invalid Token"
)

func VerifyJWT(tokenString string, secretKey string) (bool, string) {
token,err:=jwt.Parse(tokenString,func (token *jwt.Token)(interface{},error)  {
	if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
		return nil,fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	

})




	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	return []byte(secretKey), nil
	// })
	// if err != nil {
	// 	log.Printf("Error While Decoding  JWT :%v\n", err)
	// 	return false, ""
	// }
	// fmt.Printf("Token is: %t\n", token.Valid)

	// if token.Valid {
	// 	fmt.Printf("Hello world ")
	// 	claims, ok := token.Claims.(jwt.MapClaims)
	// 	fmt.Printf("Ok ok : %v\n", ok)
	// 	if ok {
	// 		if email, ok := claims["email"].(string); ok {
	// 			fmt.Printf("Email is : %v\n", email)
	// 			return true, email
	// 		}
	// 	}
	// 	if !ok {
	// 		return false, ""
	// 	}
	// }
	// return false, INVALID_TOKEN
}