package models

import (
	"github.com/goravel/framework/database/orm"
)

type Enquiry struct {
	orm.Model
	Name    string `gorm:"size:255;not null" json:"name"`
	Email   string `gorm:"size:255;not null" json:"email"`
	Phone   string `gorm:"size:50;not null" json:"phone"`
	Service string `gorm:"size:100" json:"service"`
	Message string `gorm:"type:text;not null" json:"message"`
}
