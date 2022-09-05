package models_test

import (
	"testing"
	"vendors/models"
)

func TestValidatePositive(t *testing.T) {
	vendor := new(models.Vendor)
	vendor.Email = "Vendor-1@Vendor.Com"
	vendor.ContactNo = "999999999"
	vendor.Name = "Vendor-1"
	err := vendor.Validate()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestValidateNegative(t *testing.T) {
	vendor1 := new(models.Vendor)
	vendor1.ContactNo = "999999999"
	vendor1.Name = "Vendor-1"
	err := vendor1.Validate()
	if err == nil {
		//t.Log(err.Error())
		t.Fail()
	}

	vendor2 := new(models.Vendor)
	vendor2.Email = "Vendor-1@Vendor.Com"
	//vendor.ContactNo = "999999999"
	vendor2.Name = "Vendor-1"
	err2 := vendor2.Validate()
	if err2 == nil {
		//t.Log(err.Error())
		t.Fail()
	}
	vendor3 := new(models.Vendor)

	vendor3.Email = "Vendor-1@Vendor.Com"
	vendor3.ContactNo = "999999999"
	//vendor.Name = "Vendor-1"
	err3 := vendor3.Validate()
	if err3 == nil {
		//t.Log(err.Error())
		t.Fail()
	}
}
