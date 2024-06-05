package services

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)




type Mail struct {
	From string
	To []string
	Subject string
}


type SendMail interface {
	Send()
	SendHmtl()
}



func (email *Mail) SendHmtl(filePath string, data interface{}) error {

	auth := smtp.PlainAuth(
		"",
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_HOST"),
	)

	content, err := HtmlToString(filePath,data)

	if err != nil {
		log.Errorf("Email couldnt be sent \n Error: %v",err)
		return fiber.NewError(500, "An internal server ocurred")
	}
	
	 msg := "Subject: "+ email.Subject +"\n" +header() +"\n\n"+  content

	 addr := os.Getenv("MAIL_HOST") +":" +os.Getenv("MAIL_PORT")
	 from := email.From

	 log.Info(addr)

	 if from == "" {
		from = os.Getenv("MAIL_USERNAME")
	 }

	err = smtp.SendMail(
		addr,
		auth,
		from,
		email.To,
		[]byte(msg),
	)



	if err != nil {
		log.Errorf("Email couldnt be sent \n Error: %v",err)
		return fiber.NewError(500, "An internal server ocurred")
	}

	return nil
}



func header() string {
	return "Mime-Version: 1.0;\r\nContent-Type: text/html; charset=UTF-8 \r\n"
}



func HtmlToString(path string, data interface{}) (htmlContent string,Error error ){

	var content bytes.Buffer
	
	t , err := template.ParseFiles("./template/otp-verification.html")

	if err != nil {
		return "" ,fiber.NewError(500, fmt.Sprint(err))
	}

	err = t.Execute(&content,data)

	if err != nil {
		return "", fiber.NewError(500, fmt.Sprint(err))
	}

	return content.String(), nil
}