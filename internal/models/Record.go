package models

import "errors"

type Record struct {
	Uuid        string `json:"uuid" db:"uuid"`
	FirstName   string `json:"first_name" binding:"required" db:"first_name"`
	LastName    string `json:"last_name" binding:"required" db:"last_name"`
	MobilePhone string `json:"mobile_phone" binding:"required" db:"mobile_phone"`
	HomePhone   string `json:"home_phone" db:"home_phone"`
}

type RecordInput struct {
	FirstName   string `json:"first_name" form:"firstName"`
	LastName    string `json:"last_name" form:"lastName"`
	MobilePhone string `json:"mobile_phone" form:"mobilePhone"`
	HomePhone   string `json:"home_phone" form:"homePhone"`
}

func (i RecordInput) Validate() error {
	if i == *new(RecordInput) {
		return errors.New("structure has no values")
	}

	return nil
}
