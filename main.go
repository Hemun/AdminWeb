package main

import (
	"air-q/Routes"
	db "air-q/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//db connection
	db.Connect()
	fmt.Println("Hello Api...")
	//app
	app := fiber.New()
	app.Use(cors.New())
	Routes.Setup(app)

	app.Listen(":3030")
	//if err != nil {
	//	return
	//}
}
