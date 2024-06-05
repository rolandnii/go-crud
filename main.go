package main

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/rolandnii/roland-auth/auth/routes"
	"github.com/rolandnii/roland-auth/database"
)




func main() {

		
	
	f, _ :=  os.OpenFile("main.log",os.O_CREATE |os.O_WRONLY | os.O_APPEND, 0666)
	log.SetOutput(f)
	
	err := godotenv.Load(".env")

	if err != nil {
		log.Error(err)
	}

	dbConfig := database.DatabaseSetup()
	dsn , _ := dbConfig.Url()

	err = database.Connect(dsn)
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	routes.Init(app)

	app.Listen(":"+os.Getenv("PORT"))

	
}