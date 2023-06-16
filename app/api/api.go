package api

import (
	"fmt"
	"github.com/enrinal/crispy-spoon/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type API struct {
	gin        *gin.Engine
	flightRepo repository.FlightRepository
}

func NewAPI(flightRepository repository.FlightRepository) API {
	r := gin.Default()
	api := API{
		flightRepo: flightRepository,
		gin:        r,
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
		return
	})

	r.POST("/login", api.Login)

	flightV1 := r.Group("/api/v1", Auth())
	{
		flightV1.GET("/flights", api.FetchAllFlight)
		flightV1.GET("/flights/:id", api.FetchFlightByID)
		flightV1.POST("/flights", api.CreateFlight)
	}

	r.GET("/admin", Auth(), AuthzAdmin(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "admin",
		})
		return
	})

	return api
}

func (api *API) Handler() *gin.Engine {
	return api.gin
}

func (api *API) Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":"+port, api.Handler())
}
