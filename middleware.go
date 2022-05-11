package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

type RequestCtx struct {
	Format string `json:"format"` // The format to return the resource in, could be either json or xml
	Prefer string `json:"prefer"` // The format to return the resource in, could be either minimal, representation, or OperationOutcome
}

func (s *Server) MountMiddlewares() {
	r := s.Router
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.AllowContentType("application/fhir+json", "application/fhir+xml"))
	r.Use(middleware.StripSlashes)
	r.Use(PreferRequestCtx)
	r.Use(SetDefaultTimeZone)
}

func PreferRequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqCtx := RequestCtx{
			Format: "json",
			Prefer: "representation",
		}

		contentType := r.URL.Query().Get("_format")

		if contentType == "xml" {
			reqCtx.Format = "xml"
		}

		// only check if its transaction request
		if r.Method != "POST" && r.Method != "PUT" && r.Method != "PATCH" {
			// check for prefer header
			prefer := r.Header.Get("Prefer")

			if prefer == "return=minimal" {
				reqCtx.Prefer = "minimal"
			}

			if prefer == "return=OperationOutcome" {
				reqCtx.Prefer = "OperationOutcome"
			}
		}

		ctx := context.WithValue(r.Context(), RequestCtx{}, reqCtx)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func SetDefaultTimeZone(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Date", time.Now().Format(time.RFC1123))
		next.ServeHTTP(w, r)
	})
}
