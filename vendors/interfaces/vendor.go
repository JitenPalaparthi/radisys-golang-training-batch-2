package interfaces

import "vendors/models"

type IVendor interface {
	Add(*models.Vendor) (*models.Vendor, error)
	// Get(id int) *models.Vendor
	// Update(id int, data map[string]interface{}) (*models.Vendor, error)
	// Delete(id int) (int, error)
	// GetAll() []models.Vendor
}
