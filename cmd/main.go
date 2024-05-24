package main

import (
	"fmt"
	"log"
	"net/http"

	db "exo-planets/pkg/db"
	pl "exo-planets/pkg/planets"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := db.Connection_mysql()
	if err != nil {
		log.Println("DB error is:", err)
		return
	}

	err = db.AutoMigrate(&pl.Planets{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}

	client := &pl.Client{DB: db}

	fmt.Println("hii ....this is exo-planets")

	r := mux.NewRouter()
	r.HandleFunc("/add", client.AddExoPlanets).Methods("POST")
	r.HandleFunc("/list", client.ListExoPlanets).Methods("GET")
	r.HandleFunc("/get/{id}", client.GetExoPlanetById).Methods("GET")
	r.HandleFunc("/delete/{id}", client.DeleteExoplanetById).Methods("GET")
	r.HandleFunc("/update", client.UpdateExoplanetById).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))

}
