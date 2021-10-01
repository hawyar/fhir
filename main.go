package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
	model "github.com/samply/golang-fhir-models/fhir-models/fhir"
)

var patients = make([]model.Patient, 20)

func PatientHandler(w http.ResponseWriter, r *http.Request) {

	id := "12"
	given := "Shane"
	family := "Early"
	var NameUse fhir.NameUse = 0
	nameEnd := "2012"
	phone := "1234567890"
	current := time.Now().String()
	active := true

	for i := 0; i < 20; i++ {
		patients[i] = model.Patient{
			Id: &id,
			Meta: &model.Meta{
				LastUpdated: &current,
				VersionId: &id,
				Extension: []model.Extension{
					{
						Url: "http://hl7.org/fhir/StructureDefinition/us-core-patient-birthDate",
						Id:  &id,
						Extension: []model.Extension{
							{
								Url: "http://hl7.org/fhir/StructureDefinition/us-core-birthDate",
							},
						},
					},
				},
			},
			Name: []model.HumanName{
				{
					Use: &NameUse,
					Family: &given,
					Given:  []string{family},
					Period: &model.Period{
						End:  &nameEnd,
					},
				},
			},
			Telecom: []model.ContactPoint{
				{
					Id: &id,
					Value: &phone,
				},
			},
			Active: &active,
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
		port = "4040"
	}

	r := mux.NewRouter()

	
	r.HandleFunc("/", Index)
	r.HandleFunc("/patient/{id}", PatientHandler)
	r.HandleFunc("/patient", PatientsIndexHandler)

	fmt.Println("Server is running on port: http://localhost:4040/" + port)

	log.Fatal(http.ListenAndServe(":" + port, r))
}
