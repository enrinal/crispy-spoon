package api

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type FlightFetchResponse struct {
	ID           int     `json:"id"`
	FlightNumber string  `json:"flight_number"`
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	Price        float64 `json:"price"`
	Airline      string  `json:"airline"`
}
