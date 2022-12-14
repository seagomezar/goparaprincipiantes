package main

import (
	"encoding/json"
	"net/http"
)

type Persona struct {
	Nombre string
}

func main() {
	http.HandleFunc("/persona", personaHandler)
	http.ListenAndServe(":8080", nil)
}

func personaHandler(w http.ResponseWriter, r *http.Request) {
	p := Persona{Nombre: "Sebastian"}
	json.NewEncoder(w).Encode(p)
}
