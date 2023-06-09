package main

import (
	"log"
	"net/http"

	"Golang-CRUD-API/controller"
	"Golang-CRUD-API/models"
	"Golang-CRUD-API/repository"
	"Golang-CRUD-API/service"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load the configuration
	config := LoadConfig()

	// Initialize the database connection
	dsn := config.DBConfig.Username + ":" + config.DBConfig.Password + "@tcp(" + config.DBConfig.Host + ":" + config.DBConfig.Port + ")/" + config.DBConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate the user model
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	// Create repository and service instances
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	// Create a new instance of the Gorilla Mux router
	router := mux.NewRouter()

	// Define the API endpoints
	router.HandleFunc("/api/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", userController.GetUserByID).Methods("GET")
	router.HandleFunc("/api/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}