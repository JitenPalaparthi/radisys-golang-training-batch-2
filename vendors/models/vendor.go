package models

import (
	"encoding/json"
	"errors"
)

var (
	ErrEmptyNameField      = errors.New("empty name field")
	ErrEmptyEmailField     = errors.New("empty email field")
	ErrEmptyContactNoField = errors.New("empty contactNo field")
)

type Vendor struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"index"`
	Email        string `json:"email"`
	ContactNo    string `json:"contactNo" gorm:"column:contactNo"`
	WebSite      string `json:"webSite"`
	Address      string `json:"address"`
	Status       string `json:"status"`
	LastModified int64  `json:"lastModified" gorm:"column:lastModified"`
}

func (v *Vendor) Validate() error {
	if v.Name == "" {
		//return errors.New("name field is empty")
		return ErrEmptyNameField
	}
	if v.Email == "" {
		return ErrEmptyEmailField
	}
	if v.ContactNo == "" {
		return ErrEmptyContactNoField
	}
	return nil
}

func (v *Vendor) ToBytes() ([]byte, error) {
	return json.Marshal(v)
}

func (v *Vendor) ToJSON() (string, error) {
	if buf, err := v.ToBytes(); err != nil {
		return "", nil
	} else {
		return string(buf), nil
	}
}
