package main

import (
	"github.com/enrinal/crispy-spoon/app/api"
	"github.com/enrinal/crispy-spoon/app/repository"
	"github.com/enrinal/crispy-spoon/config"
	"github.com/enrinal/crispy-spoon/model"
)

func main() {
	db := config.NewDB()
	conn, err := db.Connect()
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(&model.Flight{})

	flightRepo := repository.NewFlightRepo(conn)
	mainAPI := api.NewAPI(flightRepo)
	mainAPI.Start()
}
