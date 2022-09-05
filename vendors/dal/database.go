package dal

import (
	"time"

	"github.com/golang/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MAXRETRIES int = 10
	MAXTIME    int = 5
)

func GetConnection(dsn string) (*gorm.DB, error) {
	count := 0
RETRY:
	count++
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if count <= MAXRETRIES {
			glog.Warning("Trying to connect to database-->", count)
			time.Sleep(time.Second * time.Duration(MAXTIME))
			goto RETRY
		} else {
			return nil, err
		}
	}
	return db, err
}
