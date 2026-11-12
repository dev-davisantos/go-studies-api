package category

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type service interface {
	FindById(id int) (*CategoryResponse, error)
	FindAll() (*[]CategoryResponse, error)
	Create(cat *CategoryRequest) (*CategoryResponse, error)
	UpdateDescription(id int, cat *CategoryRequest) (*CategoryResponse, error)
	Delete(id int) error
}

type CategoryHandler struct {
	service service
}

func NewCategoryHandler(service service) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) FindById(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	intId, idErr := strconv.Atoi(strId)
	if idErr != nil {
		http.Error(w, idErr.Error(), http.StatusBadRequest)
	}

	// Simple error treatment for simplify the process
	response, searchErr := h.service.FindById(intId)
	if searchErr != nil {
		if errors.Is(searchErr, ErrCategoryNotFound) {
			http.Error(w, searchErr.Error(), http.StatusNotFound)
		} else {
			http.Error(w, searchErr.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonErr := json.NewEncoder(w).Encode(response)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
	}
}

func (h *CategoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	response, searchErr := h.service.FindAll()
	if searchErr != nil {
		if errors.Is(searchErr, ErrNoCategoriesFound) {
			http.Error(w, searchErr.Error(), http.StatusNotFound)
		}
		if errors.Is(searchErr, ErrCategoryConversionFailed) {
			http.Error(w, searchErr.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonErr := json.NewEncoder(w).Encode(response)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
	}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	response, err := h.service.Create(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	jsonErr := json.NewEncoder(w).Encode(response)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
	}
}

func (h *CategoryHandler) UpdateDescription(w http.ResponseWriter, r *http.Request) {
	var dto CategoryRequest
	id, convertErr := strconv.Atoi(r.PathValue("id"))
	if convertErr != nil {
		http.Error(w, convertErr.Error(), http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	response, err := h.service.UpdateDescription(id, &dto)
	if err != nil {
		if errors.Is(err, ErrCategoryFailedToUpdate) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if errors.Is(err, ErrCategoryNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		if errors.Is(err, ErrCategoryConversionFailed) {
			http.Error(w, ErrCategoryConversionFailed.Error(), http.StatusInternalServerError)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonErr := json.NewEncoder(w).Encode(response)
	if jsonErr != nil {
		http.Error(w, jsonErr.Error(), http.StatusInternalServerError)
	}
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, convertErr := strconv.Atoi(r.PathValue("id"))
	if convertErr != nil {
		http.Error(w, convertErr.Error(), http.StatusBadRequest)
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, ErrCategoryNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
