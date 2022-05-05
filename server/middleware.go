package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func (s *Server) MountMiddlewares() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))
}

func FHIRWarmup(next http.Handler) {

	name := "fhir-server"
	url := "http://localhost:8080/v1/CapabilityStatement"

	cap, err := NewCapabilityStatement(fhir.CapabilityStatement{
		Url:         &url,
		FhirVersion: fhir.FHIRVersion4_0_1,
		Name:        &name,
		Status:      fhir.PublicationStatusDraft,
		Kind:        fhir.CapabilityStatementKindInstance,
	})

	if err != nil {
		log.Println(err)
		return
	}

	json, err := json.Marshal(cap)

	if err != nil {
		log.Println(err)
		return
	}

	key := "CapabilityStatement-" + *cap.Id

	Set(key, string(json))
	fmt.Println("Warmup complete")

	next.ServeHTTP(nil, nil)
}

func ResourceFormat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		param := r.URL.Query().Get("_format")

		format := "json"

		if param == "xml" {
			format = "xml"
			w.Header().Set("Content-Type", "application/fhir+xml")
		} else {
			w.Header().Set("Content-Type", "application/fhir+json")
		}

		ctx = context.WithValue(ctx, "format", format)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
