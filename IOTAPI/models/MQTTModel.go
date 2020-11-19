package models

type MQTT struct {
	Topic   string `json:"topic" xml:"topic" form:"topic" validate:"required"`
	Message string `json:"message" xml:"message" form:"message" validate:"required"`
}
