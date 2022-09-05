package dal

import (
	"time"
	"vendors/models"

	"github.com/golang/glog"
	"gorm.io/gorm"
)

type VendorDB struct {
	*gorm.DB // promoted type and also can acheive inheritance thru this
}

func (v *VendorDB) Add(vendor *models.Vendor) (*models.Vendor, error) {
	v.AutoMigrate(&models.Vendor{})
	tx := v.Create(vendor)
	if tx.Error != nil {
		glog.Error(tx.Error)
		return nil, tx.Error
	}
	return vendor, nil
}

func (v *VendorDB) GetBy(id int) (*models.Vendor, error) {
	vendor := new(models.Vendor)
	tx := v.First(vendor, id)
	if tx.Error != nil {
		glog.Error(tx.Error)
		return nil, tx.Error
	}
	return vendor, nil
}

func (v *VendorDB) GetAll() ([]models.Vendor, error) {
	//vendors := make([]models.Vendor, 0)
	var vendors []models.Vendor
	tx := v.Find(&vendors)
	if tx.Error != nil {
		glog.Error(tx.Error)
		return nil, tx.Error
	}
	return vendors, nil
}

func (v *VendorDB) GetAllByStatus(status string) ([]models.Vendor, error) {
	//vendors := make([]models.Vendor, 0)
	var vendors []models.Vendor
	tx := v.Find(&vendors, map[string]interface{}{"status": status})
	if tx.Error != nil {
		glog.Error(tx.Error)
		return nil, tx.Error
	}
	return vendors, nil
}

func (v *VendorDB) UpdateBy(id int, data map[string]interface{}) (*models.Vendor, error) {
	vendor := new(models.Vendor)
	vendor.ID = id
	data["lastModified"] = time.Now().Unix()
	tx := v.DB.Model(vendor).Updates(data)
	if tx.Error != nil {
		glog.Error(tx.Error)
		return nil, tx.Error
	}
	return v.GetBy(id)
}

func (v *VendorDB) DeleteBy(id int) (int, error) {
	vendor := new(models.Vendor)
	vendor.ID = id
	tx := v.Delete(vendor)
	if tx.Error != nil {
		glog.Error(tx.Error)
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
