package api

import (
	"github.com/enrinal/crispy-spoon/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (api *API) FetchAllFlight(c *gin.Context) {

	flight, err := api.flightRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var resp []FlightFetchResponse
	for _, v := range flight {
		resp = append(resp, FlightFetchResponse{
			ID:           int(v.ID),
			FlightNumber: v.FlightNumber,
			Origin:       v.Origin,
			Destination:  v.Destination,
			Price:        v.Price,
			Airline:      v.Airline,
		})
	}

	log.Println(resp)

	c.JSON(http.StatusOK, resp)
	return
}

func (api *API) FetchFlightByID(c *gin.Context) {

	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)

	flight, err := api.flightRepo.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp := FlightFetchResponse{
		ID:           int(flight.ID),
		FlightNumber: flight.FlightNumber,
		Origin:       flight.Origin,
		Destination:  flight.Destination,
		Price:        flight.Price,
		Airline:      flight.Airline,
	}

	log.Println(resp)

	c.JSON(http.StatusOK, resp)
	return
}

func (api *API) CreateFlight(c *gin.Context) {

	var req FlightRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	flight := model.Flight{
		FlightNumber: req.FlightNumber,
		Origin:       req.Origin,
		Destination:  req.Destination,
		Price:        req.Price,
		Airline:      req.Airline,
	}

	err = api.flightRepo.Insert(&flight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp := FlightFetchResponse{
		ID:           int(flight.ID),
		FlightNumber: flight.FlightNumber,
		Origin:       flight.Origin,
		Destination:  flight.Destination,
		Price:        flight.Price,
		Airline:      flight.Airline,
	}

	log.Println(resp)

	c.JSON(http.StatusCreated, resp)
	return
}

func (api *API) UpdateFlight(c *gin.Context) {

	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)

	var req FlightRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	flight := model.Flight{
		FlightNumber: req.FlightNumber,
		Origin:       req.Origin,
		Destination:  req.Destination,
		Price:        req.Price,
		Airline:      req.Airline,
	}

	err = api.flightRepo.Update(id, &flight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp := FlightFetchResponse{
		ID:           int(flight.ID),
		FlightNumber: flight.FlightNumber,
		Origin:       flight.Origin,
		Destination:  flight.Destination,
		Price:        flight.Price,
		Airline:      flight.Airline,
	}

	log.Println(resp)

	c.JSON(http.StatusOK, resp)
	return
}

func (api *API) DeleteFlight(c *gin.Context) {

	strId := c.Param("id")
	id, _ := strconv.Atoi(strId)

	err := api.flightRepo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	resp := SuccessResponse{
		Message: "Flight deleted successfully",
	}

	log.Println(resp)

	c.JSON(http.StatusOK, resp)
	return
}
