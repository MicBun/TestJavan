package main

import (
	"github.com/MicBun/TestJavan/config"
	"github.com/MicBun/TestJavan/docs"
	"github.com/MicBun/TestJavan/route"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default env")
	}

	description := "This is a simple family asset for JavanTest.\n\n" +
		"Checkout my Github: https://github.com/MicBun\n\n" +
		"Checkout my Linkedin: https://www.linkedin.com/in/MicBun\n\n"

	docs.SwaggerInfo.Title = "Family Asset API"
	docs.SwaggerInfo.Description = description

	database := config.ConnectDataBase()
	sqlDB, _ := database.DB()
	defer sqlDB.Close()
	r := route.SetupRouter(database)
	r.Run()
}
