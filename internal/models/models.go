package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Zastita is the type for all signals of protective devices
type Signal struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type WeatherData struct {
	Name                     string  `json:"name"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	Height                   int     `json:"height"`
	TimezoneAbbrevation      string  `json:"timezone_abbrevation"`
	UtcTimeoffset            float64 `json:"utc_timeoffset"`
	ModelrunUtc              string  `json:"modelrun_utc"`
	ModelrunUpdatetimeUtc    string  `json:"modelrun_updatetime_utc"`
	Time                     string  `json:"time"`
	Pictocode                float64 `json:"pictocode"`
	Uvindex                  float64 `json:"uvindex"`
	TemperatureMax           float64 `json:"temperature_max"`
	TemperatureMin           float64 `json:"temperature_min"`
	TemperatureMean          float64 `json:"temperature_mean"`
	FelttemperatureMax       float64 `json:"felttemperature_max"`
	FelttemperatureMin       float64 `json:"felttemperature_min"`
	Winddirection            float64 `json:"winddirection"`
	PrecipitationProbability float64 `json:"precipitation_probability"`
	Rainspot                 string  `json:"rainspot"`
	PredictabilityClass      float64 `json:"predictability_class"`
	Predictability           float64 `json:"predictability"`
	Precipitation            float64 `json:"precipitation"`
	Snowfraction             float64 `json:"snowfraction"`
	SealevelpressureMax      float64 `json:"sealevelpressure_max"`
	SealevelpressureMin      float64 `json:"sealevelpressure_min"`
	SealevelpressureMean     float64 `json:"sealevelpressure_mean"`
	WindspeedMax             float64 `json:"windspeed_max"`
	WindspeedMean            float64 `json:"windspeed_mean"`
	WindspeedMin             float64 `json:"windspeed_min"`
	RelativehumidityMax      float64 `json:"relativehumidity_max"`
	RelativehumidityMin      float64 `json:"relativehumidity_min"`
	RelativehumidityMean     float64 `json:"relativehumidity_mean"`
	ConvectivePrecipitation  float64 `json:"convective_precipitation"`
	PrecipitationHours       float64 `json:"precipitation_hours"`
	Humiditygreater90Hours   float64 `json:"humiditygreater90_hours"`
	CreationDate             string  `json:"creation_date"`
}

type WeatherDataHistory struct {
	Name                     string  `json:"name"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	Height                   int     `json:"height"`
	TimezoneAbbrevation      string  `json:"timezone_abbrevation"`
	UtcTimeoffset            float64 `json:"utc_timeoffset"`
	ModelrunUtc              string  `json:"modelrun_utc"`
	ModelrunUpdatetimeUtc    string  `json:"modelrun_updatetime_utc"`
	Time                     string  `json:"time"`
	Pictocode                string  `json:"pictocode"`
	Uvindex                  string  `json:"uvindex"`
	TemperatureMax           string  `json:"temperature_max"`
	TemperatureMin           string  `json:"temperature_min"`
	TemperatureMean          string  `json:"temperature_mean"`
	FelttemperatureMax       string  `json:"felttemperature_max"`
	FelttemperatureMin       string  `json:"felttemperature_min"`
	Winddirection            string  `json:"winddirection"`
	PrecipitationProbability string  `json:"precipitation_probability"`
	Rainspot                 string  `json:"rainspot"`
	PredictabilityClass      string  `json:"predictability_class"`
	Predictability           string  `json:"predictability"`
	Precipitation            string  `json:"precipitation"`
	Snowfraction             string  `json:"snowfraction"`
	SealevelpressureMax      string  `json:"sealevelpressure_max"`
	SealevelpressureMin      string  `json:"sealevelpressure_min"`
	SealevelpressureMean     string  `json:"sealevelpressure_mean"`
	WindspeedMax             string  `json:"windspeed_max"`
	WindspeedMean            string  `json:"windspeed_mean"`
	WindspeedMin             string  `json:"windspeed_min"`
	RelativehumidityMax      string  `json:"relativehumidity_max"`
	RelativehumidityMin      string  `json:"relativehumidity_min"`
	RelativehumidityMean     string  `json:"relativehumidity_mean"`
	ConvectivePrecipitation  string  `json:"convective_precipitation"`
	PrecipitationHours       string  `json:"precipitation_hours"`
	Humiditygreater90Hours   string  `json:"humiditygreater90_hours"`
	CreationDate             string  `json:"creation_date"`
}

type Permission struct {
	RbrPk           int64  `json:"rbr_pk"`
	BrZahtevaFK     int64  `json:"br_zahteva_fk"`
	SapSifra        string `json:"sap_sifra"`
	BrojIsk         string `json:"broj_isk"`
	BrDozvole       string `json:"br_dozvole"`
	TipRadova       string `json:"tip_radova"`
	EeObjekat       string `json:"ee_objekat"`
	Tip             string `json:"tip"`
	Oznaka          string `json:"oznaka"`
	StatusEl        string `json:"status_el"`
	Izuzev          string `json:"izuzev"`
	Napomena        string `json:"napomena"`
	DozIzdao        string `json:"doz_izdao"`
	DozPrimio       string `json:"doz_primio"`
	DatPrijemaDoz   string `json:"dat_prijema_doz"`
	VremePrijemaDoz string `json:"vreme_prijema_doz"`
	StatusDoz       string `json:"status_doz"`
	NapomenaZavRad  string `json:"napomena_zav_rad"`
	DozZavIzdao     string `json:"doz_zav_izdao"`
	DozZavPrimio    string `json:"doz_zav_primio"`
	DatZavRadova    string `json:"dat_zav_radova"`
	VremeZavRad     string `json:"vreme_zav_rad"`
}

type Request struct {
	RbrPk      int64  `json:"rbr_pk"`
	Rco        string `json:"rco"`
	BrojIsk    string `json:"broj_isk"`
	BrZahteva  string `json:"br_zahteva"`
	BrojKps    string `json:"broj_kps"`
	Grupa      string `json:"grupa"`
	Int        string `json:"int"`
	SapSifra   string `json:"sap_sifra"`
	EeObjekat  string `json:"ee_objekat"`
	Tip        string `json:"tip"`
	Oznaka     string `json:"oznaka"`
	Opis       string `json:"opis"`
	PlDatumOd  string `json:"pl_datum_od"`
	PlVremeOd  string `json:"pl_vreme_od"`
	PlDatumDo  string `json:"pl_datum_do"`
	PlVremeDo  string `json:"pl_vreme_do"`
	Nodob      string `json:"nodob"`
	TipRadova  string `json:"tip_radova"`
	Uslovi     string `json:"uslovi"`
	NapVeza    string `json:"nap_veza"`
	PodZahteva string `json:"pod_zhteva"`
	IskOdobrio string `json:"isk_odobrio"`
}

type Outage struct {
	Datizv        string `json:"datizv"`
	Vrepoc        string `json:"vrepoc"`
	Vrezav        string `json:"vrezav"`
	Traj          string `json:"trajanje"`
	IpsId         string `json:"ips_id"`
	TipOb         string `json:"tipob"`
	Opis          string `json:"opis"`
	ImeDalekovoda string `json:"ime_dalekovoda"`
	PoljeTrafo    string `json:"polje_trafo"`
	Org1          string `json:"org1"`
	Org2          string `json:"org2"`
	Nazvrpd       string `json:"nazvrpd"`
	Uzrok         string `json:"uzrok"`
	VrmUsl        string `json:"vrm_usl"`
	Tekst         string `json:"tekst"`
	IdStavke      string `json:"id_stavke"`
	IdSeq         string `json:"id_seq"`
}

type Disconnection struct {
	Datizv        string `json:"datizv"`
	Vrepoc        string `json:"vrepoc"`
	Vrezav        string `json:"vrezav"`
	Traj          string `json:"trajanje"`
	IpsId         string `json:"ips_id"`
	TipOb         string `json:"tipob"`
	Opis          string `json:"opis"`
	ImeDalekovoda string `json:"ime_dalekovoda"`
	PoljeTrafo    string `json:"polje_trafo"`
	Org1          string `json:"org1"`
	Org2          string `json:"org2"`
	Nazvrpd       string `json:"nazvrpd"`
	Razlog        string `json:"razlog"`
	Tekst         string `json:"tekst"`
}

type Plan struct {
	IdPogOdr  string `json:"pog_odr_id"`
	IdSapFLok string `json:"sap_id"`
	IdIPSFLok string `json:"ips_id"`
	Opis      string `json:"opis"`
	IdPogPlan string `json:"id_pog_plan"`
	PlOdr     string `json:"pl_odr"`
	TksStOd   string `json:"tks_st_od"`
	TipNaloga string `json:"tip_naloga"`
	DatumPoc  string `json:"datum_poc"`
	DatumZav  string `json:"datum_zav"`
	Id        string `json:"id"`
}

// User is the type for users
type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
