package filedb

import (
	"os"
	"vendors/models"
)

type VendorFDB struct {
	FileName string
}

func (v *VendorFDB) Add(vendor *models.Vendor) (*models.Vendor, error) {
	f, err := os.Create(v.FileName)
	if err != nil {
		return nil, err
	}
	str, err := vendor.ToJSON()
	if err != nil {
		return nil, err
	}
	f.WriteString(str)

	return vendor, nil
}

func (v *VendorFDB) GetBy(id int) (*models.Vendor, error) {
	return nil, nil
}

func (v *VendorFDB) GetAll() ([]models.Vendor, error) {

	return nil, nil
}

func (v *VendorFDB) UpdateBy(id int, data map[string]interface{}) (*models.Vendor, error) {
	return nil, nil
}

func (v *VendorFDB) DeleteBy(id int) (int, error) {
	return 0, nil
}
