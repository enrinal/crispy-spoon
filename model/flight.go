package model

import "gorm.io/gorm"

type Flight struct {
	gorm.Model
	FlightNumber string  `json:"flight_number"`
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	Price        float64 `json:"price"`
	Airline      string  `json:"airline"`
}
