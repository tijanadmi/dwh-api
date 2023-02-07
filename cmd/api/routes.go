package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/*func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}*/

func (app *application) routes() http.Handler {
	
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Post("/authenticate", app.authenticate)
	mux.Get("/refresh", app.refreshToken)
	mux.Get("/logout", app.logout)

	mux.Get("/status", app.statusHandler)
	mux.Get("/weather/{year}", app.getWeather)
	mux.Get("/weather_forecast", app.getWeatherForecast)
	mux.Get("/weather_history/{year}", app.getWeatherHistory)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequired)

		mux.Get("/weather", app.getWeather)

	})

	return mux

}
