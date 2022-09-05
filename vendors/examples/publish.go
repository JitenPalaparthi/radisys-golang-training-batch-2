package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
	"vendors/dal"
	"vendors/mb"
)

var (
	DSN  string
	CONN []string = []string{"localhost:29092"}
)

func main() {
	flag.StringVar(&DSN, "db", "host=localhost user=admin password=admin123 dbname=vendorsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai", "--db host=localhost user=admin password=admin123 dbname=vendorsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	flag.Parse()

	db, err := dal.GetConnection(DSN)
	if err != nil {
		panic(err)
	}
	vdb := new(dal.VendorDB)
	vdb.DB = db
	for {
		time.Sleep(time.Second * 5)
		vendors, err := vdb.GetAllByStatus("created")
		for _, vendor := range vendors {
			//fmt.Println(v)
			val, _ := vendor.ToBytes()
			err = mb.Publish(CONN, "vendor.created", []byte(strconv.Itoa(vendor.ID)), val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(vdb.UpdateBy(vendor.ID, map[string]interface{}{"status": "active"}))
			}
		}
	}
}
