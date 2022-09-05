package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"vendors/dal"
	"vendors/handlers"

	"github.com/golang/glog"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  string
	// INFO|WARNING|ERROR|FATAL
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|ERROR|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	//flag.Parse()
}

func main() {

	flag.StringVar(&PORT, "port", "50082", "--port 50082")
	flag.StringVar(&DSN, "db", "host=localhost user=admin password=admin123 dbname=vendorsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", "--db host=localhost user=admin password=admin123 dbname=vendorsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	flag.Set("stderrthreshold", "INFO")
	flag.Parse()

	if port := os.Getenv("PORT"); port != "" {
		PORT = port
	}

	if dbConn := os.Getenv("DB_CONN"); dbConn != "" {
		DSN = dbConn
	}
	glog.Info("Connecting to database...")
	db, err := dal.GetConnection(DSN)
	if err != nil {
		glog.Fatal("Database Connection:", err.Error())
	}

	glog.Info("instantiating vendorDB")
	vdb := new(dal.VendorDB)
	vdb.DB = db

	// vfdb := new(filedb.VendorFDB)
	// vfdb.FileName = "vendors.fdb"

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
		ctx.Abort()
	}, Ping)

	r.GET("/ping", Ping)
	r.GET("/greet", Greet("Hello Folks!"))

	vhandler := new(handlers.VendorHandler)
	vhandler.IVendor = vdb
	//vhandler.IVendor = vfdb
	v1 := r.Group("/v1/public/vendor")
	{
		v1.POST("/add", vhandler.Create())
		v1.GET("/:id", vhandler.GetBy())
		v1.GET("/all", vhandler.GetAll())
		v1.PUT("/:id", vhandler.UpdateBy())
		v1.DELETE("/:id", vhandler.DeleteBy())
	}

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
