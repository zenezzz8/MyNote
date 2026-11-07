package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserID uint   `gorm:"not null" json:"user_id"`
	Type   string `gorm:"type:varchar(50); not null" json:"type"`
	Body   string `gorm:"type:text; not null" json:"body"`
	Check  bool   `gorm:"Default:false" json:"check"`
}
