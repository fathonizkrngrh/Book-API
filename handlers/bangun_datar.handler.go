package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"quiz3/utils"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BangunDatarAPI interface {
	SegitigaSamaSisi(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Persegi(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	PersegiPanjang(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Lingkaran(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	JajarGenjang(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type bangunDatarAPI struct {
}

func NewBangunDatarAPI() *bangunDatarAPI {
	return &bangunDatarAPI{}
}

func (c *bangunDatarAPI) SegitigaSamaSisi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	alasStr := r.URL.Query().Get("alas")
	tinggiStr := r.URL.Query().Get("tinggi")
	hitung := r.URL.Query().Get("hitung")

	alas, err := strconv.ParseFloat(alasStr, 64)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid 'alas' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	tinggi, err := strconv.ParseFloat(tinggiStr, 64)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid 'tinggi' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	var result float64
	if hitung == "luas" {
		result = 0.5 * alas * tinggi
	} else if hitung == "keliling" {
		result = 3 * alas
	} else{
		utils.ErrorJSON(w, errors.New("Invalid 'hitung' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success hitung %s segitiga sama sisi dengan Alas %s dan Tinggi %s", hitung, alasStr, tinggiStr),
		Data: result,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (c *bangunDatarAPI)Persegi(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sisiStr := r.URL.Query().Get("sisi")
	hitung := r.URL.Query().Get("hitung")

	sisi, err := strconv.ParseFloat(sisiStr, 64)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid 'sisi' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	var result float64
	if hitung == "luas" {
		result = sisi * sisi
	} else if hitung == "keliling" {
		result = 4 * sisi
	} else {
		utils.ErrorJSON(w, errors.New("Invalid 'hitung' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success hitung %s Persegi dengan Sisi %s", hitung, sisiStr),
		Data: result,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (c *bangunDatarAPI)PersegiPanjang(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panjangStr := r.URL.Query().Get("panjang")
	lebarStr := r.URL.Query().Get("lebar")
	hitung := r.URL.Query().Get("hitung")

	panjang, err := strconv.ParseFloat(panjangStr, 64)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid 'panjang' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	lebar, err := strconv.ParseFloat(lebarStr, 64)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid 'lebar' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	var result float64
	if hitung == "luas" {
		result = panjang * lebar
	} else if hitung == "keliling" {
		result = 2*(panjang+lebar)
	} else {
		utils.ErrorJSON(w, errors.New("Invalid 'hitung' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success hitung %s Persegi Panjang dengan Panjang %s dan Lebar %s", hitung, panjangStr, lebarStr),
		Data: result,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (c *bangunDatarAPI) Lingkaran(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jariJariStr := r.URL.Query().Get("jariJari")
	hitung := r.URL.Query().Get("hitung")

	jariJari, err := strconv.ParseFloat(jariJariStr, 64)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid 'jariJari' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	var result float64
	if hitung == "luas" {
		result = 3.14 * jariJari * jariJari
	} else if hitung == "keliling" {
		result = 2 * 3.14 * jariJari
	} else {
		utils.ErrorJSON(w, errors.New("Invalid 'hitung' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success hitung %s Lingkaran dengan Jari-jari %s", hitung, jariJariStr),
		Data: result,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (c *bangunDatarAPI) JajarGenjang(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sisiStr := r.URL.Query().Get("sisi")
	alasStr := r.URL.Query().Get("alas")
	tinggiStr := r.URL.Query().Get("tinggi")
	hitung := r.URL.Query().Get("hitung")

	sisi, err := strconv.ParseFloat(sisiStr, 64)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid 'sisi' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	alas, err := strconv.ParseFloat(alasStr, 64)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid 'alas' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	tinggi, err := strconv.ParseFloat(tinggiStr, 64)
	if err != nil {
		utils.ErrorJSON(w, errors.New("Invalid 'tinggi' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	var result float64
	if hitung == "luas" {
		result = alas * tinggi
	} else if hitung == "keliling" {
		result = 2*(alas+sisi)
	} else {
		utils.ErrorJSON(w, errors.New("Invalid 'hitung' parameter"), "BAD REQUEST",http.StatusBadRequest)
	  	return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success hitung %s Jajar genjang dengan Sisi %s, Alas %s dan Tinggi %s", hitung, sisiStr, alasStr, tinggiStr),
		Data: result,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}
