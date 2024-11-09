package handler

import (
	"fmt"
	"inventaris/model"
	"inventaris/service"
	"inventaris/utils"
	"inventaris/validation"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type ItemHandler struct {
	Service *service.ItemService
}

func NewItemHandler(service *service.ItemService) *ItemHandler {
	return &ItemHandler{Service: service}
}

func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	categoryName := r.URL.Query().Get("categoryName")
	maxUsageDays := r.URL.Query().Get("maxUsageDays") == "true"
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}

	items, totalItems, totalPages, err := h.Service.GetAllItems(name, categoryName, maxUsageDays, limit, page)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if len(items) == 0 {
		utils.SendJSONResponse(w, false, http.StatusNotFound, "No items found", nil)
		return
	}

	utils.SendJSONResponsePagination(w, true, page, limit, totalItems, totalPages, http.StatusOK, "", items)
}

func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item

	item.Name = r.FormValue("name")
	if categoryID := r.FormValue("category_id"); categoryID != "" {
		if cid, err := strconv.Atoi(categoryID); err == nil {
			item.CategoryID = cid
		}
	}
	if price := r.FormValue("price"); price != "" {
		if p, err := strconv.ParseFloat(price, 64); err == nil {
			item.Price = p
		}
	}

	item.PurchaseDate = r.FormValue("purchase_date")
	purchaseDate, err := time.Parse("2006-01-02", item.PurchaseDate)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "Invalid purchase date format", nil)
		return
	}

	currentDate := time.Now()
	usageDays := int(currentDate.Sub(purchaseDate).Hours() / 24)

	item.UsageDays = usageDays

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "Unable to parse form", nil)
		return
	}

	file, _, err := r.FormFile("photo_url")
	var photoPath string
	if err == nil {
		defer file.Close()

		photoFileName := "temp_photo.jpg"
		photoPath, err = utils.SaveUploadedFile(file, photoFileName)
		if err != nil {
			utils.SendJSONResponse(w, false, http.StatusInternalServerError, "Failed to save file", nil)
			return
		}

		item.Photo = photoPath
	}

	if err := validation.ValidateItem(&item); err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := h.Service.CreateItem(&item); err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "ID not found", nil)
		return
	}

	if photoPath != "" {
		newPhotoFileName := fmt.Sprintf("photo_%d.jpg", item.ID)
		newPhotoPath := filepath.Join("uploads", newPhotoFileName)
		if err := os.Rename(photoPath, newPhotoPath); err != nil {
			utils.SendJSONResponse(w, false, http.StatusInternalServerError, "Failed to rename photo file", nil)
			return
		}
		item.Photo = newPhotoPath

		if err := h.Service.UpdateItem(&item); err != nil {
			utils.SendJSONResponse(w, false, http.StatusInternalServerError, "Failed to update item photo", nil)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	utils.SendJSONResponse(w, true, http.StatusOK, "Item created successfully", item)
}

func (h *ItemHandler) GetItemByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	item, err := h.Service.GetItemByID(id)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusNotFound, "Item not found", nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "", item)
}

func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {

	itemID := chi.URLParam(r, "id")
	id, err := strconv.Atoi(itemID)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "Invalid item ID", nil)
		return
	}
	var item model.Item

	item.Name = r.FormValue("name")

	if categoryID := r.FormValue("category_id"); categoryID != "" {
		if cid, err := strconv.Atoi(categoryID); err == nil {
			item.CategoryID = cid
		}
	}

	if price := r.FormValue("price"); price != "" {
		if p, err := strconv.ParseFloat(price, 64); err == nil {
			item.Price = p
		}
	}
	item.PurchaseDate = r.FormValue("purchase_date")

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "Unable to parse form", nil)
		return
	}

	file, _, err := r.FormFile("photo_url")
	var photoPath string
	if err == nil {
		defer file.Close()

		photoFileName := fmt.Sprintf("photo_%d.jpg", id)

		photoPath, err = utils.SaveUploadedFile(file, photoFileName)
		if err != nil {
			utils.SendJSONResponse(w, false, http.StatusInternalServerError, "Failed to save file", nil)
			return
		}
	} else {
		fmt.Println("No file uploaded or error opening file:", err)
	}
	item.Photo = photoPath

	if err := validation.ValidateItem(&item); err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}
	purchaseDate, err := time.Parse("2006-01-02", item.PurchaseDate)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "Invalid purchase date format", nil)
		return
	}
	currentDate := time.Now()
	usageDays := int(currentDate.Sub(purchaseDate).Hours() / 24)
	item.UsageDays = usageDays

	item.ID = id

	if err := h.Service.UpdateItem(&item); err != nil {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "ID not found", nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "Item updated successfully", item)
}

func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	_, err := h.Service.GetItemByID(id)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusNotFound, "Item not found", nil)
		return
	}

	if err := h.Service.DeleteItem(id); err != nil {
		utils.SendJSONResponse(w, false, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "Item deleted successfully", nil)
}

func (h *ItemHandler) GetTotalInvestment(w http.ResponseWriter, r *http.Request) {

	items, err := h.Service.Repo.GetAllInvestment()
	if err != nil {

		utils.SendJSONResponse(w, false, http.StatusInternalServerError, "Failed calculate total investment", nil)
		return
	}

	var totalInvestment float64
	var totalDepreciatedValue float64

	for _, item := range items {

		totalInvestment += item.Price

		depreciatedValue := item.Price
		if item.UsageDays > 30 && item.DepreciationRate > 0 {

			periods := item.UsageDays / 30
			for i := 0; i < periods; i++ {
				depreciatedValue = depreciatedValue * (1 - item.DepreciationRate/100)
			}
		}

		totalDepreciatedValue += depreciatedValue
	}

	investment := model.Investment{
		TotalInvestment:  totalInvestment,
		DepreciatedValue: totalDepreciatedValue,
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "", investment)
}

func (h *ItemHandler) GetItemInvestmentByID(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil || id <= 0 {
		utils.SendJSONResponse(w, false, http.StatusBadRequest, "Invalid item ID", nil)
		return
	}

	item, err := h.Service.Repo.GetByID(id)
	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusNotFound, "Item not found", nil)
		return
	}

	var depreciatedValue float64
	if item.UsageDays > 30 && item.DepreciationRate > 0 {
		depreciatedValue = item.Price

		periods := item.UsageDays / 30
		for i := 0; i < periods; i++ {
			depreciatedValue = depreciatedValue * (1 - item.DepreciationRate/100)
		}
	} else {

		depreciatedValue = item.Price
	}

	itemInvestment := model.ItemInvestment{
		ItemID:           item.ID,
		Name:             item.Name,
		InitialPrice:     item.Price,
		DepreciatedValue: depreciatedValue,
		DepreciationRate: item.DepreciationRate,
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "", itemInvestment)
}

func (h *ItemHandler) GetItemsReplacementNeeded(w http.ResponseWriter, r *http.Request) {

	items, err := h.Service.GetAllItemsReplacementNeeded()

	if err != nil {
		utils.SendJSONResponse(w, false, http.StatusInternalServerError, "Gagal mengambil data barang yang perlu diganti", nil)
		return
	}

	var replacementItems []model.ItemReplacement

	for _, item := range items {

		replacementItems = append(replacementItems, model.ItemReplacement{
			ID:                  item.ID,
			Name:                item.Name,
			Category:            item.CategoryName,
			PurchaseDate:        item.PurchaseDate,
			TotalUsageDays:      item.UsageDays,
			ReplacementRequired: true,
		})

	}

	if len(replacementItems) == 0 {
		utils.SendJSONResponse(w, false, http.StatusNotFound, "No items need replacement", nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "", replacementItems)
}
