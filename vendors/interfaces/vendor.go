package interfaces

import "vendors/models"

type IVendor interface {
	Add(*models.Vendor) (*models.Vendor, error)
	GetBy(id int) (*models.Vendor, error)
	UpdateBy(id int, data map[string]interface{}) (*models.Vendor, error)
	DeleteBy(id int) (int, error)
	GetAll() ([]models.Vendor, error)
}
