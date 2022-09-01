package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
)

func main() {

	flag.StringVar(&PORT, "port", "50082", "--port 50082")
	flag.Parse()

	if port := os.Getenv("PORT"); port != "" {
		PORT = port
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
		ctx.Abort()
	}, Ping)

	r.GET("/ping", Ping)
	r.GET("/greet", Greet("Hello Folks!"))
	r.Run(":" + PORT)

}

func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func Greet(msg string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": msg})
	}
}
