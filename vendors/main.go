package main

import (
	"flag"
	"net/http"
	"os"
	"vendors/dal"
	"vendors/handlers"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  string
)

func main() {

	flag.StringVar(&PORT, "port", "50082", "--port 50082")
	flag.StringVar(&DSN, "db", "host=localhost user=admin password=admin123 dbname=vendorsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", "--db host=localhost user=admin password=admin123 dbname=vendorsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	flag.Parse()

	if port := os.Getenv("PORT"); port != "" {
		PORT = port
	}

	if dbConn := os.Getenv("DB_CONN"); dbConn != "" {
		DSN = dbConn
	}

	db, err := dal.GetConnection(DSN)
	if err != nil {
		panic(err)
	}

	vdb := new(dal.VendorDB)
	vdb.DB = db

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
		ctx.Abort()
	}, Ping)

	r.GET("/ping", Ping)
	r.GET("/greet", Greet("Hello Folks!"))

	vhandler := new(handlers.VendorHandler)
	vhandler.IVendor = vdb

	r.POST("v1/vendor", vhandler.Create())

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
