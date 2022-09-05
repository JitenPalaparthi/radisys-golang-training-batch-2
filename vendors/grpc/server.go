package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
	"vendors/dal"
	"vendors/interfaces"
	"vendors/models"
	pb "vendors/protos"

	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type VService struct {
	pb.UnimplementedVendorServiceServer
	IVendor interfaces.IVendor
}

func (s *VService) Add(ctx context.Context, in *pb.Vendor) (*pb.VendorResponse, error) {
	out := new(pb.VendorResponse)
	vendor := new(models.Vendor)
	vendor.Name = in.Name
	vendor.Email = in.Email
	vendor.ContactNo = in.ContactNo
	vendor.Address = in.Address
	vendor.WebSite = in.Website
	vendor.Status = "created"
	vendor.LastModified = time.Now().Unix()

	err := vendor.Validate()
	if err != nil {
		return nil, err
	}

	v, err := s.IVendor.Add(vendor)
	if err != nil {
		return nil, err
	}
	out.Code = 201
	out.Message = "Successfully Added"
	vendorOut := new(pb.Vendor)
	vendorOut.ID = int32(v.ID)
	vendorOut.Name = v.Name
	vendorOut.Email = v.Email
	vendorOut.ContactNo = v.ContactNo
	vendorOut.Address = v.Address
	vendorOut.Website = v.WebSite
	vendorOut.Status = v.Status
	vendorOut.LastModified = v.LastModified
	out.VendorOut = vendorOut
	return out, nil
}

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

	if port := os.Getenv("APP_PORT"); port != "" {
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

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		glog.Fatal(err)
	}

	vservice := new(VService)
	vservice.IVendor = vdb
	s := grpc.NewServer()
	pb.RegisterVendorServiceServer(s, vservice)
	glog.Infof("Server is listening on %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		glog.Fatal(err)
	}
}
