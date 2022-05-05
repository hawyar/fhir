package server

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
	"github.com/segmentio/ksuid"
)

func PostCapabilityStatementHandler(w http.ResponseWriter, r *http.Request) {
	capStmt, err := CreateCapabilityStatement(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	format := r.Context().Value("format").(string)

	json, err := json.Marshal(capStmt)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	key := "CapabilityStatement-" + *capStmt.Name

	Set(key, string(json))

	if format == "xml" {
		xml, err := xml.Marshal(capStmt)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(xml)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(json)
	return
}

// func GetCapabilityStatement(_ *http.Request) (fhir.CapabilityStatement, error) {
// 	// check the
// }

func NewPatientCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var patient fhir.Patient

		err := json.NewDecoder(r.Body).Decode(&patient)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}

		empty := ""

		if patient.Id == &empty {
			id := NewID()
			patient.Id = &id
		}

		ctx := context.WithValue(r.Context(), "patientID", patient)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func patientCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		if r.Method == "GET" {
			id := chi.URLParam(r, "id")

			if id == "" {
				http.Error(w, "Please provide an id", 400)
				return
			}

			fmt.Println(id)

			ctx := r.Context()
			ctx = context.WithValue(ctx, "patientId", id)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		next.ServeHTTP(w, r)
	})
}

func GetPatientHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "patientId")

	var patient fhir.Patient

	pat := Get(id)

	if pat == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Resource not found"))
	}

	err := json.Unmarshal([]byte(pat), &patient)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	format := r.Context().Value("format").(string)

	if format == "xml" {
		xml, err := xml.Marshal(patient)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}

		w.Header().Set("Content-Type", "application/fhir+xml")
		w.WriteHeader(http.StatusOK)
		w.Write(xml)
		return
	}

	json, err := json.Marshal(patient)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
	patient, err := CreatePatient(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	json, err := json.Marshal(patient)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	key := "Patient-" + *patient.Id

	Set(key, string(json))

	format := r.Context().Value("format").(string)

	if format == "xml" {
		xml, err := xml.Marshal(patient)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		w.Header().Set("Content-Type", "application/fhir+xml")
		w.WriteHeader(http.StatusCreated)
		w.Write(xml)
		return
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func NewProcedureHandler(w http.ResponseWriter, r *http.Request) {
	procedure, err := CreateProcedure(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	json, err := json.Marshal(procedure)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	key := "Procedure-" + ksuid.New().String()

	Set(key, string(json))

	format := r.Context().Value("format").(string)

	if format == "xml" {
		xml, err := xml.Marshal(procedure)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		w.Header().Set("Content-Type", "application/fhir+xml")
		w.WriteHeader(http.StatusCreated)
		w.Write(xml)
		return
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func NewID() string {
	id := ksuid.New()
	return id.String()
}

func NewObservationHandler(w http.ResponseWriter, r *http.Request) {
	observation, err := NewObservation(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	json, err := json.Marshal(observation)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	key := "Observation-" + *observation.Id

	Set(key, string(json))

	format := r.Context().Value("format").(string)

	if format == "xml" {
		xml, err := xml.Marshal(observation)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		w.Header().Set("Content-Type", "application/fhir+xml")
		w.WriteHeader(http.StatusCreated)
		w.Write(xml)
		return
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func NewLocationHandler(w http.ResponseWriter, r *http.Request) {
	location, err := NewLocation(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	json, err := json.Marshal(location)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	key := "Location-" + *location.Id

	Set(key, string(json))

	format := r.Context().Value("format").(string)

	if format == "xml" {
		xml, err := xml.Marshal(location)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		w.Header().Set("Content-Type", "application/fhir+xml")
		w.WriteHeader(http.StatusCreated)
		w.Write(xml)
		return
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func NewOrganizationHandler(w http.ResponseWriter, r *http.Request) {
	organization, err := NewOrganization(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	json, err := json.Marshal(organization)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	key := "Organization-" + *organization.Id

	Set(key, string(json))

	format := r.Context().Value("format").(string)

	if format == "xml" {
		xml, err := xml.Marshal(organization)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		w.Header().Set("Content-Type", "application/fhir+xml")
		w.WriteHeader(http.StatusCreated)
		w.Write(xml)
		return
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}
