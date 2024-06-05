package resource

import (
	"time"

	"github.com/rolandnii/roland-auth/model"
)


type User struct {
	ID        uint      ` json:"id"`
	FirstName string    `json:"first_name"`
	Phone     *string   `json:"phone"`
	Email     string    `json:"email"`
	LastName  *string   `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Companies []model.Company `json:"companies"`
}



func UserResource (user model.User) User {

	return User{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Phone: user.Phone,
		Companies: user.Companies,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}


func UserColllection(users []model.User) []User {
	var userCollection []User

	for  _, v := range users {
		userCollection = append(userCollection, UserResource(v))
	}

	return userCollection
}