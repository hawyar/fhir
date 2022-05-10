package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

// Essential middleware for FHIR server
func (s *Server) MountMiddlewares() {
	r := s.Router
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.AllowContentType("application/fhir+json", "application/fhir+xml"))
	r.Use(SetDefaultTimeZone)
	r.Use(ResourceContentType)
}

// ResourceContentType sets the requested resource content type
func ResourceContentType(next http.Handler) http.Handler {
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

// PreferTransactionCtx specifies what the server return for a transaction request. Could be one of minimal, representative. See
func PreferTransactionCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// if not a transaction request (POST, PUT, PATCH) then skip
		if r.Method != "POST" && r.Method != "PUT" && r.Method != "PATCH" {
			next.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()

		prefer := r.Header.Get("Prefer")

		// if not set then skip
		if prefer == "" {
			next.ServeHTTP(w, r)
			return
		}

		if prefer == "return=minimal" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if prefer == "return=representation" {
			ctx := context.WithValue(ctx, "PreferTxReturn", "representation")
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		if prefer == "return=OperationOutcome" {
			ctx := context.WithValue(ctx, "PreferTxReturn", "OperationOutcome")
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		next.ServeHTTP(w, r)
	})

}

func SetDefaultTimeZone(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Date", time.Now().Format(time.RFC1123))
		next.ServeHTTP(w, r)
	})
}
