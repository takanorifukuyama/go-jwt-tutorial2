package main

import (
	"github.com/gin-gonic/gin"
    "github.com/takanorifukuyama/go-jwt-tutorial2/auth"
)

func main() {
	r := gin.Default()

	r.GET("/api", auth.CreateToken)
	r.GET("/api/private", auth.UseToken)
	r.Run(":8080")
}
