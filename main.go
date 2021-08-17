package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	model "github.com/samply/golang-fhir-models/fhir-models/fhir"
)


var patients = make([]model.Patient, 20)

func PatientHandler(w http.ResponseWriter, r *http.Request) {

	id := "12"
	given := "Shane"
	family := "Early"

	for i := 0; i < 20; i++ {
		patients[i] = model.Patient{
			Id: &id,
			Name: []model.HumanName{
				{
					Family: &given,
					Given:  []string{family},
				},
			},
		}

		vars := mux.Vars(r)

		id = vars["id"]

		for _, patient := range patients {
			if *patient.Id == id {
				w.WriteHeader(http.StatusOK)
				m, _ := patient.MarshalJSON()
				w.Write([]byte(m))
				return
			}
		}
	}
}

func PatientsIndexHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"resourceType":"Bundle","entry":[{"resourceType":"Patient"}]}`))
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Consensus Networks <> FHIR Proxy!\n"))
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()

	r.HandleFunc("/", Index)
	r.HandleFunc("/patient/{id}", PatientHandler)
	r.HandleFunc("/patient", PatientsIndexHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
