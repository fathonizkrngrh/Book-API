package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"quiz3/forms"
	"quiz3/repositories"
	"quiz3/utils"

	"github.com/julienschmidt/httprouter"
)

type CategoryAPI interface {
	GetAllCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetCategoryById(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	InsertCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type categoryAPI struct {
	categoryRepo repositories.CategoryRepo
}

func NewCategoryAPI(categoryRepo repositories.CategoryRepo) *categoryAPI {
	return &categoryAPI{categoryRepo}
}

func (c *categoryAPI) GetAllCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	categories, err := c.categoryRepo.GetAll(ctx)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Error fetching categories"), "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Success get all categories",
		Data: categories,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (c *categoryAPI) InsertCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var newCategory forms.InsertCategory
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
	}

	isDuplicate, err := c.categoryRepo.IsDuplicate(ctx, newCategory.Name)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	if isDuplicate {
		utils.ErrorJSON(w, errors.New("Data category already exist"), "CONFLICT", http.StatusConflict)
		return
	}

	if err := c.categoryRepo.Insert(ctx, newCategory); err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success insert new category %v", newCategory.Name),
		Data: newCategory,
	}
	
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (c *categoryAPI) GetCategoryById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idCategory = utils.ParseInt(ps.ByName("id"))
	
	category, err := c.categoryRepo.GetById(ctx, idCategory)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success get category with id %v", idCategory),
		Data: category,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (c *categoryAPI) UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var newCategory forms.UpdateCategory

    if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
        utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
    }

	isDuplicate, err := c.categoryRepo.IsDuplicate(ctx, newCategory.Name)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	if isDuplicate {
		utils.ErrorJSON(w, errors.New("Data category already exist"), "CONFLICT", http.StatusConflict)
		return
	}

    var idCategory = utils.ParseInt(ps.ByName("id"))

    if err := c.categoryRepo.UpdateByID(ctx, newCategory, idCategory); err != nil {
        utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
    }

    payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success update category %v", newCategory.Name),
		Data: newCategory,
	}
  
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var idCategory = utils.ParseInt(ps.ByName("id"))

    if err := c.categoryRepo.DeleteByID(ctx, idCategory); err != nil {
        utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
    }

    payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success delete category %v", idCategory),
	}
  
	utils.WriteJSON(w, http.StatusCreated, payload)
}