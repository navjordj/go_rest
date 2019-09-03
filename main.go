package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	name string
	age  int
}

func (p *person) birthday() {
	p.age += 1
}

func (p *person) ageInTenYears() int {
	return p.age + 10
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("hmm")
}

func addPerson(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		w.Header().Set("Content-type", "application/json")

		//var newPerson person

		for k, v := range mux.Vars(r) {
			fmt.Print(k, v)
		}
	}

}

var Personer []person

func main() {

	per := person{"Per", 30}
	per.birthday()

	Personer = append(Personer, per)

	fmt.Println("Starter serveren")
	r := mux.NewRouter()
	r.HandleFunc("/person", getPerson).Methods(("GET"))
	r.HandleFunc("/person", getPerson).Methods(("POST"))
	http.ListenAndServe(":8000", r)
	fmt.Println("Lytter nå på port 8000")
}
