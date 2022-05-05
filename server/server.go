package server

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

	fmt.Println("Config file loaded")
}

type Server struct {
	Router    *chi.Mux
	CreatedAt time.Time
	Config    *FHIRConfig
	Port      string
}

type FHIRConfig struct {
	Version string `json:"FHIRVersion"`
	Name    string `json:"name"`
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

	r.Use(middleware.AllowContentType("application/fhir+json", "application/fhir+xml"))

	r.Use(middleware.Heartbeat("/ping"))

	r.With(ResourceFormat).Route("/v1", func(r chi.Router) {
		r.Route("/metadata", func(r chi.Router) {
			r.Post("/", PostCapabilityStatementHandler)
			// r.Get("/", GetCapabilityStatementHandler)
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

	fmt.Println(viper.GetViper().AllKeys())

	fmt.Println("Server: http://127.0.0.1:4141/")

	if os.Getenv("PORT") != "" {
		port = ":" + os.Getenv("PORT")
	}
	http.ListenAndServe(port, r)
}
