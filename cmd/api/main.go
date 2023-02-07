package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tijanadmi/dwh_api/internal/repository"
	"github.com/tijanadmi/dwh_api/internal/repository/dbrepo"

	_ "github.com/godror/godror"
)

const version = "1.0.0"

const port = 8080

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	DSN          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	logger       *log.Logger
}

func main() {

	var app application

	flag.StringVar(&app.DSN, "dsn", "", "Oracle connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret1", "", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.OracleDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "__Host-refresh_token",
		CookieDomain:  app.CookieDomain,
	}

	logger.Println("Audience and Issuer", app.JWTAudience, app.JWTIssuer)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
