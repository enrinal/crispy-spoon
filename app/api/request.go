package api

type FlightRequest struct {
	FlightNumber string  `json:"flight_number" binding:"required"`
	Origin       string  `json:"origin" binding:"required"`
	Destination  string  `json:"destination" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Airline      string  `json:"airline" binding:"required"`
}
