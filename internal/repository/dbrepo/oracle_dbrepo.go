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

const dbTimeout = time.Second * 10

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
					COALESCE(to_char(PICTOCODE), ''),
					COALESCE(to_char(UVINDEX), ''),
					COALESCE(to_char(TEMPERATUREMAX), ''),
					COALESCE(to_char(TEMPERATUREMIN), ''),
					COALESCE(to_char(TEMPERATUREMEAN), ''),
					COALESCE(to_char(FELTTEMPERATUREMAX), ''),
					COALESCE(to_char(FELTTEMPERATUREMIN), ''),
					COALESCE(to_char(WINDDIRECTION), ''),
					COALESCE(to_char(PRECIPITATIONPROBABILITY), ''),
					COALESCE(to_char(RAINSPOT), ''),
					COALESCE(to_char(PREDICTABILITYCLASS), ''),
					COALESCE(to_char(PREDICTABILITY), ''),
					COALESCE(to_char(PRECIPITATION), ''),
					COALESCE(to_char(SNOWFRACTION), ''),
					COALESCE(to_char(SEALEVELPRESSUREMAX), ''),
					COALESCE(to_char(SEALEVELPRESSUREMIN), ''),
					COALESCE(to_char(SEALEVELPRESSUREMEAN), ''),
					COALESCE(to_char(WINDSPEEDMAX), ''),
					COALESCE(to_char(WINDSPEEDMEAN), ''),
					COALESCE(to_char(WINDSPEEDMIN), ''),
					COALESCE(to_char(RELATIVEHUMIDITYMAX), ''),
					COALESCE(to_char(RELATIVEHUMIDITYMIN), ''),
					COALESCE(to_char(RELATIVEHUMIDITYMEAN), ''),
					COALESCE(to_char(CONVECTIVEPRECIPITATION), ''),
					COALESCE(to_char(PRECIPITATIONHOURS), ''),
					COALESCE(to_char(HUMIDITYGREATER90HOURS), ''),
					COALESCE(to_char(CREATION_DATE), '')
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

func (m *OracleDBRepo) GetPermissions1(year string) ([]*models.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select RBR_PK,
				BR_ZAHTEVA_FK,
				SAP_SIFRA,
				BROJ_ISK,
				BR_DOZVOLE,
				TIP_RADOVA,
				EE_OBJEKAT,
				TIP,
				OZNAKA,
				STATUS_EL,
				IZUZEV,
				NAPOMENA,
				DOZ_IZDAO,
				DOZ_PRIMIO,
				DAT_PRIJEMA_DOZ,
				VREME_PRIJEMA_DOZ,
				STASTUS_DOZ,
				NAPOMENA_ZAV_RAD,
				DOZ_ZAV_IZDAO,
				DOZ_ZAV_PRIMIO,
				DAT_ZAV_RADOVA,
				VREME_ZAV_RAD
			  from dozvole_1 where substr(dat_prijema_doz,length(dat_prijema_doz)-3,4) = :1
			  or  substr(dat_zav_radova,length(dat_zav_radova)-3,4)= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission

	for rows.Next() {
		var data models.Permission
		err := rows.Scan(
			&data.RbrPk,
			&data.BrZahtevaFK,
			&data.SapSifra,
			&data.BrojIsk,
			&data.BrDozvole,
			&data.TipRadova,
			&data.EeObjekat,
			&data.Tip,
			&data.Oznaka,
			&data.StatusEl,
			&data.Izuzev,
			&data.Napomena,
			&data.DozIzdao,
			&data.DozPrimio,
			&data.DatPrijemaDoz,
			&data.VremePrijemaDoz,
			&data.StatusDoz,
			&data.NapomenaZavRad,
			&data.DozZavIzdao,
			&data.DozZavPrimio,
			&data.DatZavRadova,
			&data.VremeZavRad,
		)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, &data)
	}

	return permissions, nil
}

func (m *OracleDBRepo) GetPermissions23(year string) ([]*models.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select RBR_PK,
				BR_ZAHTEVA_FK,
				SAP_SIFRA,
				BROJ_ISK,
				BR_DOZVOLE,
				TIP_RADOVA,
				EE_OBJEKAT,
				TIP,
				OZNAKA,
				STATUS_EL,
				IZUZEV,
				NAPOMENA,
				DOZ_IZDAO,
				DOZ_PRIMIO,
				DAT_PRIJEMA_DOZ,
				VREME_PRIJEMA_DOZ,
				STASTUS_DOZ,
				NAPOMENA_ZAV_RAD,
				DOZ_ZAV_IZDAO,
				DOZ_ZAV_PRIMIO,
				DAT_ZAV_RADOVA,
				VREME_ZAV_RAD
			  from dozvole_2_3 where substr(dat_prijema_doz,length(dat_prijema_doz)-3,4) = :1
			  or  substr(dat_zav_radova,length(dat_zav_radova)-3,4)= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var permissions []*models.Permission

	for rows.Next() {
		var data models.Permission
		err := rows.Scan(
			&data.RbrPk,
			&data.BrZahtevaFK,
			&data.SapSifra,
			&data.BrojIsk,
			&data.BrDozvole,
			&data.TipRadova,
			&data.EeObjekat,
			&data.Tip,
			&data.Oznaka,
			&data.StatusEl,
			&data.Izuzev,
			&data.Napomena,
			&data.DozIzdao,
			&data.DozPrimio,
			&data.DatPrijemaDoz,
			&data.VremePrijemaDoz,
			&data.StatusDoz,
			&data.NapomenaZavRad,
			&data.DozZavIzdao,
			&data.DozZavPrimio,
			&data.DatZavRadova,
			&data.VremeZavRad,
		)

		if err != nil {
			return nil, err
		}

		permissions = append(permissions, &data)
	}

	return permissions, nil
}

func (m *OracleDBRepo) GetRequests1(year string) ([]*models.Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select RBR_PK,
			RCO,
			COALESCE(to_char(BROJ_ISK), ''),
			BR_ZAHTEVA,
			BROJ_KPS,
			GRUPA,
			INT,
			SAP_SIFRA,
			EE_OBJEKAT,
			TIP,
			OZNAKA,
			OPIS,
			PL_DATUM_OD,
			PL_VREME_OD,
			PL_DATUM_DO,
			PL_VREME_DO,
			NODOB,
			TIP_RADOVA,
			USLOVI,
			NAP_VEZA,
			POD_ZAHTEVA,
			ISK_ODOBRIO
			from zahtevi_1 where substr(PL_DATUM_OD,length(PL_DATUM_OD)-3,4) = :1
			or  substr(PL_DATUM_DO,length(PL_DATUM_DO)-3,4)= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var requests []*models.Request

	for rows.Next() {
		var data models.Request
		err := rows.Scan(
			&data.RbrPk,
			&data.Rco,
			&data.BrojIsk,
			&data.BrZahteva,
			&data.BrojKps,
			&data.Grupa,
			&data.Int,
			&data.SapSifra,
			&data.EeObjekat,
			&data.Tip,
			&data.Oznaka,
			&data.Opis,
			&data.PlDatumOd,
			&data.PlVremeOd,
			&data.PlDatumDo,
			&data.PlVremeDo,
			&data.Nodob,
			&data.TipRadova,
			&data.Uslovi,
			&data.NapVeza,
			&data.PodZahteva,
			&data.IskOdobrio,
		)

		if err != nil {
			return nil, err
		}

		requests = append(requests, &data)
	}

	return requests, nil
}

func (m *OracleDBRepo) GetRequests23(year string) ([]*models.Request, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select RBR_PK,
			RCO,
			COALESCE(to_char(BROJ_ISK), ''),
			BR_ZAHTEVA,
			BROJ_KPS,
			GRUPA,
			INT,
			SAP_SIFRA,
			EE_OBJEKAT,
			TIP,
			OZNAKA,
			OPIS,
			PL_DATUM_OD,
			PL_VREME_OD,
			PL_DATUM_DO,
			PL_VREME_DO,
			NODOB,
			TIP_RADOVA,
			USLOVI,
			NAP_VEZA,
			POD_ZAHTEVA,
			ISK_ODOBRIO
			from zahtevi_2_3 where substr(PL_DATUM_OD,length(PL_DATUM_OD)-3,4) = :1
			or  substr(PL_DATUM_DO,length(PL_DATUM_DO)-3,4)= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var requests []*models.Request

	for rows.Next() {
		var data models.Request
		err := rows.Scan(
			&data.RbrPk,
			&data.Rco,
			&data.BrojIsk,
			&data.BrZahteva,
			&data.BrojKps,
			&data.Grupa,
			&data.Int,
			&data.SapSifra,
			&data.EeObjekat,
			&data.Tip,
			&data.Oznaka,
			&data.Opis,
			&data.PlDatumOd,
			&data.PlVremeOd,
			&data.PlDatumDo,
			&data.PlVremeDo,
			&data.Nodob,
			&data.TipRadova,
			&data.Uslovi,
			&data.NapVeza,
			&data.PodZahteva,
			&data.IskOdobrio,
		)

		if err != nil {
			return nil, err
		}

		requests = append(requests, &data)
	}

	return requests, nil
}

func (m *OracleDBRepo) GetOutages(year string) ([]*models.Outage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select to_char(DATIZV,'dd.mm.yyyy'),
		to_char(VREPOC,'dd.mm.yyyy HH24:MI:SS'),
		to_char(VREZAV,'dd.mm.yyyy HH24:MI:SS'),
		TRAJ,
		IPS_ID,
		TIPOB,
		OPIS,
		IME_DALEKOVODA,
		POLJE_TRAFO,
		ORG1,
		ORG2,
		NAZVRPD,
		UZROK,
		VRM_USL,
		TEKST,
		COALESCE(to_char(ID_STAVKE), ''),
		COALESCE(to_char(ID_SEQ), '')
 		from DWH_ISPADI_V  v
 		where TO_CHAR( v.datizv, 'yyyy' )= :1`

	rows, err := m.DB.QueryContext(ctx, query, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var o []*models.Outage

	for rows.Next() {
		var data models.Outage
		err := rows.Scan(
			&data.Datizv,
			&data.Vrepoc,
			&data.Vrezav,
			&data.Traj,
			&data.IpsId,
			&data.TipOb,
			&data.Opis,
			&data.ImeDalekovoda,
			&data.PoljeTrafo,
			&data.Org1,
			&data.Org2,
			&data.Nazvrpd,
			&data.Uzrok,
			&data.VrmUsl,
			&data.Tekst,
			&data.IdStavke,
			&data.IdSeq,
		)

		if err != nil {
			return nil, err
		}

		o = append(o, &data)
	}

	return o, nil
}

func (m *OracleDBRepo) GetDisconnections(year string) ([]*models.Disconnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select to_char(DATIZV,'dd.mm.yyyy'),
		to_char(VREPOC,'dd.mm.yyyy HH24:MI:SS'),
		to_char(VREZAV,'dd.mm.yyyy HH24:MI:SS'),
		TRAJ,
		IPS_ID,
		TIPOB,
		OPIS,
		IME_DALEKOVODA,
		POLJE_TRAFO,
		ORG1,
		ORG2,
		NAZVRPD,
		RAZLOG,
		TEKST
 		from DWH_ISKLJUCENJA_V  v
 		where TO_CHAR( v.datizv, 'yyyy' )= :1`

	rows, err := m.DB.QueryContext(ctx, query, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var d []*models.Disconnection

	for rows.Next() {
		var data models.Disconnection
		err := rows.Scan(
			&data.Datizv,
			&data.Vrepoc,
			&data.Vrezav,
			&data.Traj,
			&data.IpsId,
			&data.TipOb,
			&data.Opis,
			&data.ImeDalekovoda,
			&data.PoljeTrafo,
			&data.Org1,
			&data.Org2,
			&data.Nazvrpd,
			&data.Razlog,
			&data.Tekst,
		)

		if err != nil {
			return nil, err
		}

		d = append(d, &data)
	}

	return d, nil
}

func (m *OracleDBRepo) GetPlans(year string) ([]*models.Plan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select COALESCE(to_char(ID_POG_ODR), ''),
				ID_SAP_F_LOK,
				ID_IPS_F_LOK,
				OPIS,
				COALESCE(to_char(ID_POG_PLAN), ''),
				COALESCE(to_char(PL_ODR), ''),
				TKS_ST_OD,
				TIP_NALOGA,
				to_char(DATUM_POC,'dd.mm.yyyy'),
				to_char(DATUM_ZAV,'dd.mm.yyyy'),
				COALESCE(to_char(ID), '')
				from PLAN_O_VNP_V
				where to_char(DATUM_POC,'yyyy')= :1
				or to_char(DATUM_ZAV,'dd.mm.yyyy')= :2`

	rows, err := m.DB.QueryContext(ctx, query, year, year)
	if err != nil {
		fmt.Println("Pogresan upit ili nema rezultata upita")
		return nil, err
	}
	defer rows.Close()

	var p []*models.Plan

	for rows.Next() {
		var data models.Plan
		err := rows.Scan(
			&data.IdPogOdr,
			&data.IdSapFLok,
			&data.IdIPSFLok,
			&data.Opis,
			&data.IdPogPlan,
			&data.PlOdr,
			&data.TksStOd,
			&data.TipNaloga,
			&data.DatumPoc,
			&data.DatumZav,
			&data.Id,
		)

		if err != nil {
			return nil, err
		}

		p = append(p, &data)
	}

	return p, nil
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

	query := `select id, username, password from dwh_services_users where username = :1`

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

	query := `select id, username, password from dwh_services_users where id = :1`

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
