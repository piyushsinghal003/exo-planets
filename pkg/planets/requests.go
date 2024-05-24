package planets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Planets struct {
	gorm.Model
	Name        string
	Description string
	Distance    int
	Radius      int
	Mass        int
	Category    string
}

type Client struct {
	DB *gorm.DB
	Planets
}

func (client *Client) AddExoPlanets(w http.ResponseWriter, r *http.Request) {
	planet := Planets{}
	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		panic(err)
	}

	result := client.DB.Create(&planet)
	if result.Error != nil {
		fmt.Println("failed to create record: " + result.Error.Error())
		return
	}
}

func (client *Client) ListExoPlanets(w http.ResponseWriter, _ *http.Request) {
	var planets []Planets
	result := client.DB.Find(&planets)
	if result.Error != nil {
		// handle error
		fmt.Println("failed to retrieve records: " + result.Error.Error())
		return
	}
	fmt.Fprintf(w, "planets Endpoint")
	json.NewEncoder(w).Encode(planets)
}

func (cl *Client) GetExoPlanetById(w http.ResponseWriter, r *http.Request) {
	var planet Planets

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := cl.DB.Where("ID = ?", id).Find(&planet)
	if result.Error != nil {
		// handle error
		fmt.Println("failed to retrieve records by id: %d with error: "+result.Error.Error(), id)
		return
	}
	//result.Scan(planet)
	fmt.Fprintf(w, "planets Endpoint")
	json.NewEncoder(w).Encode(&planet)
}

func (cl *Client) DeleteExoplanetById(w http.ResponseWriter, r *http.Request) {
	var planet Planets
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result := cl.DB.Where("ID = ?", id).Delete(&planet)
	if result.Error != nil {
		// handle error
		fmt.Println("failed to retrieve records by id: %d with error: "+result.Error.Error(), id)
		return
	} else if result.RowsAffected == 0 {
		fmt.Fprintf(w, "no data found with id: %d", id)
	} else {
		fmt.Fprintf(w, "Entry deleted")
	}
}

func (cl *Client) UpdateExoplanetById(w http.ResponseWriter, r *http.Request) {
	var planet Planets
	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		panic(err)
	}
	result := cl.DB.Where("ID = ?", planet.ID).UpdateColumns(&planet)
	if result.Error != nil {
		// handle error
		fmt.Println("failed to update records by id: %d with error: "+result.Error.Error(), planet.ID)
		return
	} else if result.RowsAffected == 0 {
		fmt.Fprintf(w, "no data found with id: %d", planet.ID)
	} else {
		fmt.Fprintf(w, "Entry updated")
	}
}

func (cl *Client) FuelEstimation(w http.ResponseWriter, r *http.Request) {
	var planet Planets
	params := mux.Vars(r)
	cap, err := strconv.Atoi(params["cap"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		panic(err)
	}
	g := calculateGravity(&planet)
	fe := float64(planet.Distance) / (g * g) * float64(cap)
	fmt.Fprintf(w, "Fuel Estimation to reach this exoplanet: %s will be %f", planet.Name, fe)
}

func calculateGravity(p *Planets) float64 {
	var g float64
	if p.Category == "Gas Giant" {
		g = 0.5 / float64(p.Radius) * float64(p.Radius)
	}
	if p.Category == "Terrestrial" {
		g = float64(p.Mass) / float64(p.Radius) * float64(p.Radius)
	}
	return g

}
