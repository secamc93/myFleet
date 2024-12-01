package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);unique;not null"`
	ParentID  *uint
	Parent    *User     `gorm:"foreignKey:ParentID"`
	Children  []User    `gorm:"foreignKey:ParentID"`
	Companies []Company `gorm:"many2many:user_companies;"`
}

type Company struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null;unique"`
	Users []User `gorm:"many2many:user_companies;"`
}
