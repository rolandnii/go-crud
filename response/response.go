package response

import "github.com/gofiber/fiber/v2"


type StatusCode struct {
	Code int `json:"code"`
	Status string `json:"status"`
}


type SucessResponse struct{
	StatusCode
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode 
	Error interface{} `json:"error"`
}



func Error(c *fiber.Ctx,status *fiber.Error , errorMsg interface{})  error {

	return c.Status(status.Code).JSON(ErrorResponse{
		StatusCode: StatusCode{
			Code: status.Code,
			Status: status.Message,
		},
		Error: errorMsg,
	})
}
func Success(c *fiber.Ctx,code int , data interface{})  error {
	
	return c.Status(code).JSON(SucessResponse{
		StatusCode: StatusCode{
			Code: code,
			Status: "Ok",
		},
		Data: data,
	})
}