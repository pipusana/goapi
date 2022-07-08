package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pipusana/goapi/entities"
	"github.com/pipusana/goapi/usecases"
)

type NisitHandlersAdapter interface {
	FindAllNisit(w http.ResponseWriter, r *http.Request)
	UpdateOneNisit(w http.ResponseWriter, r *http.Request)
	FindOneNisit(w http.ResponseWriter, r *http.Request)
	DeleteOneNisit(w http.ResponseWriter, r *http.Request)
	CreateNisit(w http.ResponseWriter, r *http.Request)
}

type nisitHandlers struct {
	nisitUsecase usecases.NisitUsecase
}

func NewNisitHandlers(nisitUsecase usecases.NisitUsecase) NisitHandlersAdapter {
	return &nisitHandlers{
		nisitUsecase: nisitUsecase,
	}
}

func (h nisitHandlers) FindAllNisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result, err := h.nisitUsecase.FindAllNisit()
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(result)
}

func (h nisitHandlers) UpdateOneNisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nisit_update entities.NisitUpdate
	nisit_id := mux.Vars(r)["id"]
	json.NewDecoder(r.Body).Decode(&nisit_update)

	err := h.nisitUsecase.UpdateOneNisit(nisit_id, nisit_update)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode("Nisit Updated Successfully!!")
}

func (h nisitHandlers) FindOneNisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	nisit_id := mux.Vars(r)["id"]
	result, err := h.nisitUsecase.FindOneNisit(nisit_id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(result)
}

func (h nisitHandlers) DeleteOneNisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	nisit_id := mux.Vars(r)["id"]
	err := h.nisitUsecase.DeleteOneNisit(nisit_id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode("Nisit Deleted Successfully!!")
}

func (h nisitHandlers) CreateNisit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nisit entities.Nisit
	json.NewDecoder(r.Body).Decode(&nisit)

	result, err := h.nisitUsecase.CreateNisit(&nisit)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(result)
}
