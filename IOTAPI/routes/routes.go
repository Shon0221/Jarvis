package routes

import (
	"IOTAPI/database"

	"github.com/gofiber/fiber/v2"
)

// Register :
func Register(router fiber.Router, db *database.Database) {
	registerAPI(router)

}
