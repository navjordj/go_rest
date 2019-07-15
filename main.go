package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type person struct {
	name string
	age int
}

func (p *person) birthday() {
	p.age += 1
}

func (p *person) ageInTenYears() int {
	return p.age + 10
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("walla brusjan")
}

func main() {
	//kristian := person{"Jørgen", 30}
	fmt.Println("Starter serveren")
	r := mux.NewRouter()
	r.HandleFunc("/person", getPerson).Methods(("GET"))
	http.ListenAndServe(":8000", r)
	fmt.Println("Lytter nå på port 8000")
}