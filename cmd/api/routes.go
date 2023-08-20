package main

import (
	"net/http"
	"quiz3/handlers"
	"quiz3/repositories"

	"github.com/julienschmidt/httprouter"
)

type APIHandler struct {
	CategoryAPIHandler    handlers.CategoryAPI
	BangunDatarAPIHandler    handlers.BangunDatarAPI
}

func Router() http.Handler {
	router := httprouter.New()

	categoryRepo := repositories.NewCategoryRepo()

	categoryAPIHandler := handlers.NewCategoryAPI(categoryRepo)
	bangunDatarAPIHandler := handlers.NewBangunDatarAPI()

	apiHandler := APIHandler{
		CategoryAPIHandler:    categoryAPIHandler,
		BangunDatarAPIHandler:    bangunDatarAPIHandler,
	}

	// Bangun Datar
	router.GET("/bangun-datar/segitiga-sama-sisi", apiHandler.BangunDatarAPIHandler.SegitigaSamaSisi )
	router.GET("/bangun-datar/persegi", apiHandler.BangunDatarAPIHandler.Persegi )
	router.GET("/bangun-datar/persegi-panjang", apiHandler.BangunDatarAPIHandler.PersegiPanjang )
	router.GET("/bangun-datar/lingkaran", apiHandler.BangunDatarAPIHandler.Lingkaran )
	router.GET("/bangun-datar/jajar-genjang", apiHandler.BangunDatarAPIHandler.JajarGenjang)
	
	// Category
	router.GET("/categories", apiHandler.CategoryAPIHandler.GetAllCategories)
	router.POST("/categories", apiHandler.CategoryAPIHandler.InsertCategory)
	router.PUT("/categories/:id", apiHandler.CategoryAPIHandler.UpdateCategory)
	router.DELETE("/categories/:id", apiHandler.CategoryAPIHandler.DeleteCategory)

	return router
}