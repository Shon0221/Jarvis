package main

import (
	"IOTAPI/config"
	"IOTAPI/database"
	"IOTAPI/routes"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
)

// @title IOT API 文件
// @version 1.0
// @description IOT API 文件
// @contact.email holmes.lin@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:80
// @BasePath /
func main() {
	CONFIG := config.New("./config.yaml")

	// Initialize database
	db, err := database.New(&database.DatabaseConfig{
		Driver:   CONFIG.DB.Driver,
		Host:     CONFIG.DB.Host,
		Username: CONFIG.DB.Username,
		Password: CONFIG.DB.Password,
		Port:     CONFIG.DB.Port,
		Database: CONFIG.DB.Database,
	})

	// Auto-migrate database models
	if err != nil {
		log.Println("failed to connect to database:", err.Error())
	} else {
		if db == nil {
			log.Println("failed to connect to database: db variable is nil")
		}
		//} else {
		//	//app.DB = db
		//	//err := app.DB.AutoMigrate(&models.Role{})
		//	//if err != nil {
		//	//	fmt.Println("failed to automigrate role model:", err.Error())
		//	//	return
		//	//}
		//	//err = app.DB.AutoMigrate(&models.User{})
		//	//if err != nil {
		//	//	fmt.Println("failed to automigrate user model:", err.Error())
		//	//	return
		//	//}
		//}
	}

	app := fiber.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/apidoc/doc.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	app.Get("/apidoc/*", swagger.Handler)

	apiv1 := app.Group("/v1/api")
	routes.Register(apiv1, db)

	//data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	//fmt.Println(string(data))

	log.Fatal(app.Listen(":" + CONFIG.App.Port))
}
