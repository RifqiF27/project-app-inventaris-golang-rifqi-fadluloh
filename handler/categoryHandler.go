package handler

import (
	"encoding/json"
	"inventaris/model"
	"inventaris/service"
	"inventaris/utils"
	"inventaris/validation"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CategoryHandler struct {
	Service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: service}
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.Service.GetAllCategories()
	if err != nil {

		utils.SendJSONResponse(w, false, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "", categories)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {

		utils.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validation.ValidateCategory(&category); err != nil {

		utils.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := h.Service.CreateCategory(&category); err != nil {

		utils.SendJSONResponse(w, false, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	w.WriteHeader(http.StatusCreated)

	utils.SendJSONResponse(w, true, http.StatusOK, "category added success", category)

}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id_int, _ := strconv.Atoi(id)
	category, err := h.Service.GetCategoryByID(id_int)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusNotFound, "Category not found", nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "", category)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id_int, _ := strconv.Atoi(id)

	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {

		utils.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := validation.ValidateCategory(&category); err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}

	existingCategory, err := h.Service.GetCategoryByID(id_int)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusNotFound, "Category not found", nil)
		return
	}

	category.ID = existingCategory.ID

	if err := h.Service.UpdateCategory(&category); err != nil {
		utils.SendJSONResponse(w, false, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "Category updated successfully", category)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	id_int, _ := strconv.Atoi(id)

	_, err := h.Service.GetCategoryByID(id_int)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusNotFound, "Category not found", nil)
		return
	}

	if err := h.Service.DeleteCategory(id_int); err != nil {
		utils.SendJSONResponse(w, false, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "Delete Success", nil)

}
