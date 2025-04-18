package models

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	Location string
}
