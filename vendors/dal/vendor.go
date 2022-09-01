package dal

import (
	"vendors/models"

	"gorm.io/gorm"
)

type VendorDB struct {
	*gorm.DB // promoted type and also can acheive inheritance thru this
}

func (v *VendorDB) Add(vendor *models.Vendor) (*models.Vendor, error) {
	v.AutoMigrate(&models.Vendor{})
	tx := v.Create(vendor)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return vendor, nil
}
