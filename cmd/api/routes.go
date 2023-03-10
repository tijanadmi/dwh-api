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
	// mux.Get("/weather/{year}", app.getWeather)
	// mux.Get("/weather_forecast", app.getWeatherForecast)
	// mux.Get("/weather_history/{year}", app.getWeatherHistory)
	// mux.Get("/dozvole_1/{year}", app.getPermissions1)
	// mux.Get("/dozvole_2_3/{year}", app.getPermissions23)
	// mux.Get("/zahtevi_1/{year}", app.getRequests1)
	// mux.Get("/zahtevi_2_3/{year}", app.getRequests23)
	// mux.Get("/ispadi/{year}", app.getOutages)
	// mux.Get("/iskljucenja/{year}", app.getDisconnectors)
	// mux.Get("/planovi/{year}", app.getPlans)

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(app.authRequired)

		mux.Get("/weather/{year}", app.getWeather)
		mux.Get("/weather_forecast", app.getWeatherForecast)
		mux.Get("/weather_history/{year}", app.getWeatherHistory)
		mux.Get("/dozvole_1/{year}", app.getPermissions1)
		mux.Get("/dozvole_2_3/{year}", app.getPermissions23)
		mux.Get("/zahtevi_1/{year}", app.getRequests1)
		mux.Get("/zahtevi_2_3/{year}", app.getRequests23)
		mux.Get("/ispadi/{year}", app.getOutages)
		mux.Get("/iskljucenja/{year}", app.getDisconnectors)
		mux.Get("/planovi/{year}", app.getPlans)

	})

	return mux

}
