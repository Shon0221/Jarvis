package controllers

import (
	"IOTAPI/models"
	"IOTAPI/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

// MQTTPublish godoc
// @Summary 透過 MQTT 傳送命令到 ESP32
// @Description 透過 MQTT 傳送命令
// @Tags MQTT
// @Produce  json
// @Param mqtt body models.MQTT true
// @Success 200 {object} models.MQTT
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /v1/api/mqtt [put]
func MQTTPublish() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var mqtt models.MQTT
		if err := ctx.BodyParser(&mqtt); err != nil {
			return ctx.SendStatus(400)
		}
		v := validate.Struct(mqtt)
		if !v.Validate() {
			return ctx.Status(400).JSON(fiber.Map{
				"success": false,
				"msg":     v.Errors,
			})
		}

		if ok := services.Publish(&mqtt); ok != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"success": false,
				"msg":     "Send to MQTT message failure",
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"msg":     "Send Message Success",
		})
	}
}

func MQTTSubscribe() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var mqtt models.MQTT
		if err := ctx.BodyParser(&mqtt); err != nil {
			return ctx.SendStatus(400)
		}
		v := validate.Struct(mqtt)
		if !v.Validate() {
			return ctx.Status(400).JSON(fiber.Map{
				"success": false,
				"msg":     v.Errors,
			})
		}

		if ok := services.Publish(&mqtt); ok != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"success": false,
				"msg":     "Send to MQTT message failure",
			})
		}

		return ctx.JSON(fiber.Map{
			"success": true,
			"msg":     "Send Message Success",
		})
	}
}

type HTTPError struct {
	success bool
	message string
}
