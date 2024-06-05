package model

import (
	"time"
)



type Company struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string  `json:"name"`
	UserID uint `json:"user_id"` 	
	User User 
	CreatedAt time.Time `json:"created_at"`
}
