package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rolandnii/roland-auth/database"
	"github.com/rolandnii/roland-auth/model"
	"github.com/rolandnii/roland-auth/resource"
	"github.com/rolandnii/roland-auth/response"
	"github.com/rolandnii/roland-auth/services"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return response.Error(c, fiber.ErrBadRequest, err)
	}

	errMsg, err := services.Validator(user)

	if err != nil {

		return response.Error(c, fiber.ErrInternalServerError, err)
	}

	if len(errMsg) > 0 {

		return response.Error(c, fiber.ErrUnprocessableEntity, services.ValidationResponse(errMsg))
	}

	res := database.Db.First(&user,"email = ?", user.Email)

	if res.RowsAffected > 0 {
		
		return response.Error(c, fiber.ErrUnprocessableEntity, "account already exists")
	}


	res = database.Db.Create(&user)

	if res.RowsAffected <= 0 && res.Error != nil {
		log.Error(res.Error.Error())
		return response.Error(c, fiber.ErrInternalServerError, "Backend Failedd")
	}

	
	token, err := GenerateOtp(user, 6)
	if err != nil {
		log.Error(err)
		return response.Error(c, fiber.ErrInternalServerError, "Backend Failed")
	}

	user.SendUserRegisteredOtpNotification(token)

	return response.Success(c, fiber.StatusOK, resource.UserResource(user))
}




func GenerateOtp(user model.User, numOfDigit int) (string, error) {
	token := services.GenerateToken(numOfDigit)

	hashedToken, _ := services.HashPassword(token)

	otp := model.Otp{
		Recipient: user.Email,
		Token:     hashedToken,
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {

		err := database.Db.Where("recipient = ?",otp.Recipient).Delete(&model.Otp{}).Error

		if err != nil {
			return err
		}

		err = database.Db.Create(&otp).Error
		if err != nil {
			return err
		}
		return nil
	})

	return token, err

}
