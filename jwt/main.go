package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func main() {
	token := GenToken()
	ParseToken(token)
}

// Create the Claims
func GenToken() string {
	mySigningKey := []byte("AllYourBase")

	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss
}

func ParseToken(tokenString string) {
	// token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte("AllYourBase"), nil
	// })
	// if token.Valid {
	// 	fmt.Println("You look nice today")
	// } else if ve, ok := err.(*jwt.ValidationError); ok {
	// 	if ve.Errors&jwt.ValidationErrorMalformed != 0 {
	// 		fmt.Println("That's not even a token")
	// 	} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
	// 		// Token is either expired or not active yet
	// 		fmt.Println("Timing is everything")
	// 	} else {
	// 		fmt.Println("Couldn't handle this token:", err)
	// 	}
	// } else {
	// 	fmt.Println("Couldn't handle this token:", err)
	// }
	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Foo, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
}
