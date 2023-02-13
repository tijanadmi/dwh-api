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
	GetPermissions1(year string) ([]*models.Permission, error)
	GetPermissions23(year string) ([]*models.Permission, error)
	GetRequests1(year string) ([]*models.Request, error)
	GetRequests23(year string) ([]*models.Request, error)
	GetOutages(year string) ([]*models.Outage, error)
	GetDisconnections(year string) ([]*models.Disconnection, error)
	GetPlans(year string) ([]*models.Plan, error)
}
