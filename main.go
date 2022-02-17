package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

type PatientReq struct {
	Patient fhir.Patient `json:"patient"`
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Lite FHIR Server - v1"))
		})
		r.Get("/patient/{id}", func(w http.ResponseWriter, r *http.Request) {

			id := chi.URLParam(r, "id")

			fmt.Println(id)

			var patient fhir.Patient

			injson, err := patient.MarshalJSON()

			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(injson)
		})

		r.Post("/patient", func(w http.ResponseWriter, r *http.Request) {
			var patient fhir.Patient

			fmt.Println(patient)
			fmt.Println(r.Body)

			injson, err := patient.MarshalJSON()

			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write(injson)
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Resource not found"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("Noop"))
	})
	fmt.Println("Server running: http://127.0.0.1:4141/")

	http.ListenAndServe(":3000", r)
}
