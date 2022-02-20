package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/segmentio/ksuid"
)

var pool = NewPool()

func main() {

	port := ":4141"
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Use(middleware.Heartbeat("/ping"))

	r.With(formatCtx).Route("/v1", func(r chi.Router) {
		r.Route("/Patient", func(r chi.Router) {
			r.With(PatientCtx).Post("/", NewPatientHandler)
		})
		r.Get("/metadata", CapabilityStmt)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Resource not found"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("Noop"))
	})

	fmt.Println("Server: http://127.0.0.1:4141/")

	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	http.ListenAndServe(port, r)
}

func CapabilityStmt(w http.ResponseWriter, r *http.Request) {
	cap, err := GetCapabilityStatement(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	format := r.Context().Value("format").(string)

	if format == "xml" {
		xml, err := xml.Marshal(cap)

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

	json, err := json.Marshal(cap)

	if err != nil {
		fmt.Print("beep boop")
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func formatCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		param := r.URL.Query().Get("_format")
		format := "json"
		if param == "xml" {
			format = "xml"
		}
		ctx = context.WithValue(ctx, "format", format)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func PatientCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func NewPatientHandler(w http.ResponseWriter, r *http.Request) {
	patient, err := CreatePatient(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	//client := pool.Get()
	//_, err = client.Do("SET", patient.Id, patient)

	//defer client.Close()

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
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
		w.WriteHeader(http.StatusCreated)
		w.Write(xml)
		return
	}

	json, err := json.Marshal(patient)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := pool.Get()

	res, err := client.Do("SET", *patient.Id, string(json))

	defer client.Close()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(res)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println(res)

	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func NewID() string {
	id := ksuid.New()
	return id.String()
}
