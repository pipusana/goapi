package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pipusana/goapi/adapters"
	"github.com/pipusana/goapi/configs"
	"github.com/pipusana/goapi/handlers"
	"github.com/pipusana/goapi/repositories"
	"github.com/pipusana/goapi/usecases"
)

func main() {
	// Load Configurations
	appConfig := configs.ReadConfig()

	// Set up
	mongoAdapter := adapters.NewMongoAdapter(appConfig.MongoURI, appConfig.DATABASE, appConfig.COLLECTION)
	rabbitMQAdapter := adapters.NewRabbbitMqAdapter(appConfig.QueueName, appConfig.RabbitMQURI)

	nisitRepo := repositories.NewNisitRepository(mongoAdapter)
	logRepo := repositories.NewLoggerRepository(rabbitMQAdapter)

	nisitUsecase := usecases.NewNisitUseCase(nisitRepo, logRepo)

	// Initialize the handles
	nisitHandlers := handlers.NewNisitHandlers(nisitUsecase)

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterNisitRoutes(router, nisitHandlers)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", appConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", appConfig.Port), router))
}

func RegisterNisitRoutes(router *mux.Router, nisitHandlers handlers.NisitHandlersAdapter) {
	router.HandleFunc("/", HelloWorld).Methods("GET")
	router.HandleFunc("/nisits", nisitHandlers.FindAllNisit).Methods("GET")
	router.HandleFunc("/nisits", nisitHandlers.CreateNisit).Methods("POST")
	router.HandleFunc("/nisits/{id}", nisitHandlers.FindOneNisit).Methods("GET")
	router.HandleFunc("/nisits/{id}", nisitHandlers.UpdateOneNisit).Methods("PATCH")
	router.HandleFunc("/nisits/{id}", nisitHandlers.DeleteOneNisit).Methods("DELETE")
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Hello World !!!")
}
