package dbrepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/tijanadmi/dwh_api/internal/models"
)

/* holds our connection to the database*/
type OracleDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *OracleDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *OracleDBRepo) GetWeather(year string) ([]*models.WeatherData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select NAME,
					LATITUDE,
					LONGITUDE,
					HEIGHT,
					TIMEZONEABBREVATION,
					UTCTIMEOFFSET,
					MODELRUNUTC,
					MODELRUNUPDATETIMEUTC,
					TIME,
					PICTOCODE,
					UVINDEX,
					TEMPERATUREMAX,
					TEMPERATUREMIN,
					TEMPERATUREMEAN,
					FELTTEMPERATUREMAX,
					FELTTEMPERATUREMIN,
					WINDDIRECTION,
					PRECIPITATIONPROBABILITY,
					RAINSPOT,
					PREDICTABILITYCLASS,
					PREDICTABILITY,
					PRECIPITATION,
					SNOWFRACTION,
					SEALEVELPRESSUREMAX,
					SEALEVELPRESSUREMIN,
					SEALEVELPRESSUREMEAN,
					WINDSPEEDMAX,
					WINDSPEEDMEAN,
					WINDSPEEDMIN,
					RELATIVEHUMIDITYMAX,
					RELATIVEHUMIDITYMIN,
					RELATIVEHUMIDITYMEAN,
					CONVECTIVEPRECIPITATION,
					PRECIPITATIONHOURS,
					HUMIDITYGREATER90HOURS,
					CREATION_DATE
			  from dwh_weather where substr(time,0,4)= :1`

	rows, err := m.DB.QueryContext(ctx, query, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var weatherData []*models.WeatherData

	for rows.Next() {
		var data models.WeatherData
		err := rows.Scan(
			&data.Name,
			&data.Latitude,
			&data.Longitude,
			&data.Height,
			&data.TimezoneAbbrevation,
			&data.UtcTimeoffset,
			&data.ModelrunUtc,
			&data.ModelrunUpdatetimeUtc,
			&data.Time,
			&data.Pictocode,
			&data.Uvindex,
			&data.TemperatureMax,
			&data.TemperatureMin,
			&data.TemperatureMean,
			&data.FelttemperatureMax,
			&data.FelttemperatureMin,
			&data.Winddirection,
			&data.PrecipitationProbability,
			&data.Rainspot,
			&data.PredictabilityClass,
			&data.Predictability,
			&data.Precipitation,
			&data.Snowfraction,
			&data.SealevelpressureMax,
			&data.SealevelpressureMin,
			&data.SealevelpressureMean,
			&data.WindspeedMax,
			&data.WindspeedMean,
			&data.WindspeedMin,
			&data.RelativehumidityMax,
			&data.RelativehumidityMin,
			&data.RelativehumidityMean,
			&data.ConvectivePrecipitation,
			&data.PrecipitationHours,
			&data.Humiditygreater90Hours,
			&data.CreationDate,
		)

		if err != nil {
			return nil, err
		}

		weatherData = append(weatherData, &data)
	}

	return weatherData, nil
}

func (m *OracleDBRepo) GetWeatherForecast() ([]*models.WeatherData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select NAME,
					LATITUDE,
					LONGITUDE,
					HEIGHT,
					TIMEZONEABBREVATION,
					UTCTIMEOFFSET,
					MODELRUNUTC,
					MODELRUNUPDATETIMEUTC,
					TIME,
					PICTOCODE,
					UVINDEX,
					TEMPERATUREMAX,
					TEMPERATUREMIN,
					TEMPERATUREMEAN,
					FELTTEMPERATUREMAX,
					FELTTEMPERATUREMIN,
					WINDDIRECTION,
					PRECIPITATIONPROBABILITY,
					RAINSPOT,
					PREDICTABILITYCLASS,
					PREDICTABILITY,
					PRECIPITATION,
					SNOWFRACTION,
					SEALEVELPRESSUREMAX,
					SEALEVELPRESSUREMIN,
					SEALEVELPRESSUREMEAN,
					WINDSPEEDMAX,
					WINDSPEEDMEAN,
					WINDSPEEDMIN,
					RELATIVEHUMIDITYMAX,
					RELATIVEHUMIDITYMIN,
					RELATIVEHUMIDITYMEAN,
					CONVECTIVEPRECIPITATION,
					PRECIPITATIONHOURS,
					HUMIDITYGREATER90HOURS,
					CREATION_DATE
			  from dwh_weather_forecast`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var weatherData []*models.WeatherData

	for rows.Next() {
		var data models.WeatherData
		err := rows.Scan(
			&data.Name,
			&data.Latitude,
			&data.Longitude,
			&data.Height,
			&data.TimezoneAbbrevation,
			&data.UtcTimeoffset,
			&data.ModelrunUtc,
			&data.ModelrunUpdatetimeUtc,
			&data.Time,
			&data.Pictocode,
			&data.Uvindex,
			&data.TemperatureMax,
			&data.TemperatureMin,
			&data.TemperatureMean,
			&data.FelttemperatureMax,
			&data.FelttemperatureMin,
			&data.Winddirection,
			&data.PrecipitationProbability,
			&data.Rainspot,
			&data.PredictabilityClass,
			&data.Predictability,
			&data.Precipitation,
			&data.Snowfraction,
			&data.SealevelpressureMax,
			&data.SealevelpressureMin,
			&data.SealevelpressureMean,
			&data.WindspeedMax,
			&data.WindspeedMean,
			&data.WindspeedMin,
			&data.RelativehumidityMax,
			&data.RelativehumidityMin,
			&data.RelativehumidityMean,
			&data.ConvectivePrecipitation,
			&data.PrecipitationHours,
			&data.Humiditygreater90Hours,
			&data.CreationDate,
		)

		if err != nil {
			return nil, err
		}

		weatherData = append(weatherData, &data)
	}

	return weatherData, nil
}

func (m *OracleDBRepo) GetWeatherHistory(year string) ([]*models.WeatherDataHistory, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select NAME,
					LATITUDE,
					LONGITUDE,
					HEIGHT,
					TIMEZONEABBREVATION,
					UTCTIMEOFFSET,
					MODELRUNUTC,
					MODELRUNUPDATETIMEUTC,
					TIME,
					PICTOCODE,
					UVINDEX,
					TEMPERATUREMAX,
					TEMPERATUREMIN,
					TEMPERATUREMEAN,
					FELTTEMPERATUREMAX,
					FELTTEMPERATUREMIN,
					WINDDIRECTION,
					PRECIPITATIONPROBABILITY,
					RAINSPOT,
					PREDICTABILITYCLASS,
					PREDICTABILITY,
					PRECIPITATION,
					SNOWFRACTION,
					SEALEVELPRESSUREMAX,
					SEALEVELPRESSUREMIN,
					SEALEVELPRESSUREMEAN,
					WINDSPEEDMAX,
					WINDSPEEDMEAN,
					WINDSPEEDMIN,
					RELATIVEHUMIDITYMAX,
					RELATIVEHUMIDITYMIN,
					RELATIVEHUMIDITYMEAN,
					CONVECTIVEPRECIPITATION,
					PRECIPITATIONHOURS,
					HUMIDITYGREATER90HOURS,
					CREATION_DATE
			  from dwh_weather_history where ('20'|| substr(time,length(time)-1,2)) = :1`

	rows, err := m.DB.QueryContext(ctx, query, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var weatherDataHistory []*models.WeatherDataHistory

	for rows.Next() {
		var data models.WeatherDataHistory
		err := rows.Scan(
			&data.Name,
			&data.Latitude,
			&data.Longitude,
			&data.Height,
			&data.TimezoneAbbrevation,
			&data.UtcTimeoffset,
			&data.ModelrunUtc,
			&data.ModelrunUpdatetimeUtc,
			&data.Time,
			&data.Pictocode,
			&data.Uvindex,
			&data.TemperatureMax,
			&data.TemperatureMin,
			&data.TemperatureMean,
			&data.FelttemperatureMax,
			&data.FelttemperatureMin,
			&data.Winddirection,
			&data.PrecipitationProbability,
			&data.Rainspot,
			&data.PredictabilityClass,
			&data.Predictability,
			&data.Precipitation,
			&data.Snowfraction,
			&data.SealevelpressureMax,
			&data.SealevelpressureMin,
			&data.SealevelpressureMean,
			&data.WindspeedMax,
			&data.WindspeedMean,
			&data.WindspeedMin,
			&data.RelativehumidityMax,
			&data.RelativehumidityMin,
			&data.RelativehumidityMean,
			&data.ConvectivePrecipitation,
			&data.PrecipitationHours,
			&data.Humiditygreater90Hours,
			&data.CreationDate,
		)

		if err != nil {
			return nil, err
		}

		weatherDataHistory = append(weatherDataHistory, &data)
	}

	return weatherDataHistory, nil
}

// Authenticate authenticates a user
/*func (m *OracleDBRepo) Authenticate(username, testPassword string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()



	var user models.User

	query := `select id, username, password from tis_services_users where username = :1`

	row := m.DB.QueryRowContext(ctx, query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return errors.New("incorrect password")
	} else if err != nil {
		return err
	}
	return nil
}*/

func (m *OracleDBRepo) GetUserByUsername(username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, password from tis_services_users where username = :1`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, username)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *OracleDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, username, password from tis_services_users where id = :1`

	var user models.User
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
