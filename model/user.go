package model

import (
	"time"
	"github.com/rolandnii/roland-auth/services"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"size:100;" json:"first_name" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Phone     *string   `json:"phone" validate:"required,len=10"`
	Email     string    `gorm:"unique" json:"email" validate:"required,email"`
	LastName  *string   `json:"last_name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Companies []Company
}


func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	u.Password , err = services.HashPassword(u.Password)

	return
}

func (u *User) SendUserRegisteredOtpNotification(token string) error {
	

	mail := services.Mail{
		To:      []string{u.Email},
		Subject: "Your registration for my app",
	}

	return mail.SendHmtl("./template/otp-verification.html", struct {
		Token string
	}{
		token,
	})
}
