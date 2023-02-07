package repository

import (
	"database/sql"

	"github.com/tijanadmi/dwh_api/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetWeather(year string) ([]*models.WeatherData, error)
	GetWeatherForecast() ([]*models.WeatherData, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	GetWeatherHistory(year string) ([]*models.WeatherDataHistory, error)
}
