package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context) {
	/*
	   algorithm
	*/
	secretKey := os.Getenv("TOKEN_SECRET_KEY")
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"user": "Takanori",
		"exp":  time.Now().Add(time.Minute * 1).Unix(),
	}

	/*
	   add signature for token
	*/
	tokenString, err := token.SignedString([]byte(secretKey))
	if errors.Is(err, nil) {
		c.JSON(200, gin.H{"token": tokenString})
	} else {
		c.JSON(500, gin.H{"message": "Could not generate token"})
	}
}

func UseToken(c *gin.Context) {
	/*
	   token„ÅÆÊ§úË®º
	*/
	secretKey := os.Getenv("TOKEN_SECRET_KEY")
	token, err := request.ParseFromRequest(
		c.Request,
		request.OAuth2Extractor,
		func(oken *jwt.Token) (interface{}, error) {
			b := []byte(secretKey)
			return b, nil
		},
	)
	if errors.Is(err, nil) {
		claims := token.Claims.(jwt.MapClaims)
		c.JSON(200, gin.H{
			"message": claims["user"],
			"time":    claims["exp"],
		})
	} else {
		c.JSON(401, gin.H{"error": fmt.Sprint(err)})
	}
}
