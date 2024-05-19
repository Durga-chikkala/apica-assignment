package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/Durga-chikkala/apica-assignment/cacher"
	"github.com/Durga-chikkala/apica-assignment/handler"
	"github.com/Durga-chikkala/apica-assignment/middlewares"
	"github.com/Durga-chikkala/apica-assignment/service"
	"github.com/Durga-chikkala/apica-assignment/sockets"
	"github.com/Durga-chikkala/apica-assignment/store"
)

func main() {
	// Initialize Sockets
	socket := sockets.NewManager()

	// Initialize Cache
	c := cacher.NewCache(1024, socket)

	// Initialize Configs
	configs := NewConfig()

	// Initialize Routing
	router := mux.NewRouter()

	// Initialize Middlewares
	router.Use(middlewares.CorrelationIDMiddleware, middlewares.CORS, middlewares.RequestLogger, middlewares.SetResponseHeaders)

	server := &http.Server{
		Addr:    ":" + configs.GetOrDefault("HTTP_PORT", "8000"),
		Handler: router,
	}

	// Initialize Handler, Service, Store
	cacheStore := store.New(c)
	cacheSvc := service.New(&cacheStore)
	cacheHandler := handler.New(cacheSvc, socket)

	router.HandleFunc("/cache", cacheHandler.Set)
	router.HandleFunc("/cache/{key}", cacheHandler.Get)
	router.HandleFunc("/cacheKey/{key}", cacheHandler.Delete)
	router.HandleFunc("/keys", cacheHandler.GetALLKeys)

	log.Printf("Running the server at Port: %v", ":"+configs.GetOrDefault("HTTP_PORT", "8000"))
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Error while running the server, Err: %v", err)
		return
	}
}

type Config struct{}

func NewConfig() Config {
	env := os.Getenv("APP_ENV")
	envPath := ""
	if env != "" {
		envPath = "./configs/." + env + ".env"
	} else {
		envPath = "./configs/.env"
	}

	if err := godotenv.Load(envPath); err != nil {
		log.Println("No .env file found")
		return Config{}
	}

	log.Printf("Logs are initialized path: %v", envPath)
	return Config{}
}

type Configs interface {
	Get(key string) string
	GetOrDefault(key, defaultValue string) string
}

func (c Config) Get(key string) string {
	return os.Getenv(key)
}

func (c Config) GetOrDefault(key, defaultValue string) string {
	if key == "" {
		return defaultValue
	}

	return c.Get(key)
}
