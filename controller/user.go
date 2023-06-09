package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"Golang-CRUD-API/models"
	"Golang-CRUD-API/service"

	"github.com/gorilla/mux"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		log.Println(err)
		respondWithError(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		respondWithError(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	err = uc.userService.CreateUser(&user)
	if err != nil {
		log.Println(err)
		respondWithError(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Println(err)
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := uc.userService.GetUserByID(uint(id))
	if err != nil {
		log.Println(err)
		respondWithError(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Println(err)
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		respondWithError(w, "Failed to update user", http.StatusBadRequest)
		return
	}

	user.ID = uint(id)
	err = uc.userService.UpdateUser(&user)
	if err != nil {
		log.Println(err)
		respondWithError(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Println(err)
		respondWithError(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = uc.userService.DeleteUser(uint(id))
	if err != nil {
		log.Println(err)
		respondWithError(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
