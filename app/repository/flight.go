package repository

import (
	"github.com/enrinal/crispy-spoon/model"
	"gorm.io/gorm"
)

type FlightRepository interface {
	FindAll() ([]model.Flight, error)
	FindById(id int) (*model.Flight, error)
	Insert(data *model.Flight) error
	Update(id int, data *model.Flight) error
	Delete(id int) error
}

type flightRepositoryImpl struct {
	db *gorm.DB
}

func NewFlightRepo(db *gorm.DB) *flightRepositoryImpl {
	return &flightRepositoryImpl{db}
}

func (f *flightRepositoryImpl) FindAll() ([]model.Flight, error) {
	var flights []model.Flight

	result := f.db.Find(&flights)
	if result.Error != nil {
		return nil, result.Error
	}

	return flights, nil
}

func (f *flightRepositoryImpl) FindById(id int) (*model.Flight, error) {
	var flight model.Flight

	result := f.db.First(&flight, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &flight, nil
}

func (f *flightRepositoryImpl) Insert(data *model.Flight) error {
	result := f.db.Create(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (f *flightRepositoryImpl) Update(id int, data *model.Flight) error {
	result := f.db.Model(&model.Flight{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (f *flightRepositoryImpl) Delete(id int) error {
	result := f.db.Delete(&model.Flight{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
