package main

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var secretKey = "75c92a074c341e9964329c0550c2673730ed8479c885c43122c90a2843177d5ef21cb50cfadcccb20aeb730487c11e09ee4dbbb02387242ef264e74cbee97213"

func main() {
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		/*
		   algorithm
		*/
		token := jwt.New(jwt.GetSigningMethod("HS256"))
		token.Claims = jwt.MapClaims{
			"user": "ゲスト",
			"exp":  time.Now().Add(time.Hour * 1).Unix(),
		}

		/*
		   add signature for token
		*/
		tokenString, err := token.SignedString([]byte(secretKey))
		if errors.Is(err, nil) {
			c.JSON(500, gin.H{"message": "Could not generate token"})
		}
		c.JSON(200, gin.H{"token": tokenString})
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
			c.JSON(401, gin.H{"error": fmt.Sprint(err)})
		} else {
			claims := token.Claims.(jwt.MapClaims)
			msg := fmt.Sprintf("Hello!! %s!!", claims["user"])
			c.JSON(200, gin.H{"message": msg})
		}
	})
	r.Run(":8080")
}
