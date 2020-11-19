package routes

import (
	"IOTAPI/controllers"

	"github.com/gofiber/fiber/v2"
)

func registerAPI(apiRoute fiber.Router) {
	registerMQTT(apiRoute)
}

func registerMQTT(apiRoute fiber.Router) {
	mqtt := apiRoute.Group("/mqtt")

	mqtt.Put("/", controllers.MQTTPublish())
}
