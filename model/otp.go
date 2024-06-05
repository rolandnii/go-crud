package model

import "time"


type Otp struct {
	Recipient string `gorm:"primarKey" validate:"required" json:"recipient"`
	Token string  `gorm:"unique" json:"token"`
	CreatedAt time.Time `json:"created_at"`
}


