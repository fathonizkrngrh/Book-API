package main

import (
	"net/http"
	"quiz3/handlers"
	"quiz3/repositories"

	"github.com/julienschmidt/httprouter"
)

type APIHandler struct {
	CategoryAPIHandler    	handlers.CategoryAPI
	BookAPIHandler    		handlers.BookAPI
	BangunDatarAPIHandler   handlers.BangunDatarAPI
}

func Router() http.Handler {
	router := httprouter.New()

	categoryRepo := repositories.NewCategoryRepo()
	bookRepo := repositories.NewBookRepo()

	categoryAPIHandler := handlers.NewCategoryAPI(categoryRepo)
	bookAPIHandler := handlers.NewBookAPI(bookRepo)
	bangunDatarAPIHandler := handlers.NewBangunDatarAPI()

	apiHandler := APIHandler{
		CategoryAPIHandler:    	categoryAPIHandler,
		BookAPIHandler:    		bookAPIHandler,
		BangunDatarAPIHandler:  bangunDatarAPIHandler,
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
	router.GET("/categories/:id", apiHandler.CategoryAPIHandler.GetCategoryById)

	// Category
	router.GET("/books", apiHandler.BookAPIHandler.GetAllBook)
	router.POST("/books", apiHandler.BookAPIHandler.InsertBook)
	router.PUT("/books/:id", apiHandler.BookAPIHandler.UpdateBook)
	router.DELETE("/books/:id", apiHandler.BookAPIHandler.DeleteBook)
	router.GET("/books/:id", apiHandler.BookAPIHandler.GetBookById)

	return router
}