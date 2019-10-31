package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var secretKey = os.Getenv("TOKEN_SECRET_KEY")

func main() {
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		/*
		   algorithm
		*/
		token := jwt.New(jwt.GetSigningMethod("HS256"))
		token.Claims = jwt.MapClaims{
			"user": "Takanori",
			"exp":  time.Now().Add(time.Hour * 1).Unix(),
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
	})

	r.GET("/api/private", func(c *gin.Context) {
		/*
		   tokenの検証
		*/
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
			msg := fmt.Sprintf("Hello!! %s!!", claims["user"])
			c.JSON(200, gin.H{"message": msg})
		} else {
			c.JSON(401, gin.H{"error": fmt.Sprint(err)})
		}
	})
	r.Run(":8080")
}
