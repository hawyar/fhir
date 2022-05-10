package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

var pool = NewPool()
var config = NewConfig()

// Server is the main FHIR server
type Server struct {
	Router    *chi.Mux
	CreatedAt time.Time
	Config    *FHIRConfig
	Port      string
}

// FHIRConfig loads server configf from /config.json
type FHIRConfig struct {
	FHIRVersion string `json:"fhirversion"`
	Name        string `json:"name"`
	Format      string `json:"format"`
	Description string `json:"description"`
}

func init() {
	viper.SetConfigFile("./config.json")

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(config)

	if err != nil {
		fmt.Println("Error unmarshalling config file:", err)
		os.Exit(1)
	}

	fmt.Println("Loaded config from file:", viper.ConfigFileUsed())
}

func main() {
	server := NewServer()
	server.Run()
}

func NewConfig() *FHIRConfig {
	return &FHIRConfig{}
}

func NewServer() *Server {
	return &Server{
		Router:    chi.NewRouter(),
		CreatedAt: time.Now(),
	}
}

func (server *Server) Run() {
	server.Config = config

	server.MountMiddlewares()

	port := ":4141"

	r := server.Router

	r.Use(middleware.Heartbeat("/ping"))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/metadata", func(r chi.Router) {
			r.Get("/", GetCapabilityStatementHandler)
		})
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Not Found"))
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
	})

	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}

	fmt.Println("Server running on port", port)
	fmt.Println("FHIR Server:" + server.Config.Name)
	fmt.Println("version:", server.Config.FHIRVersion)

	http.ListenAndServe(port, r)
}
