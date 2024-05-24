package planets

import "net/http"

type PlanetsImpl interface {
	AddExoPlanets(w http.ResponseWriter, r *http.Request)
	ListExoPlanets(w http.ResponseWriter, _ *http.Request)
	GetExoPlanetById(w http.ResponseWriter, r *http.Request)
	DeleteExoplanetById(w http.ResponseWriter, r *http.Request)
	UpdateExoplanetById(w http.ResponseWriter, r *http.Request)
	FuelEstimation(w http.ResponseWriter, r *http.Request)
}
