package main

import (
	"api/config"
	"api/models"
	"api/product"
	"api/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db, e := config.MySQL()
	if e != nil{
		log.Fatal(e)
	}
	eb := db.Ping()
	if eb != nil{
		panic(eb.Error())
	}
	fmt.Println("Connection Succeded")
	http.HandleFunc("/products/insert", PostProduct)
	http.HandleFunc("/products/update", UpdateProduct)
	http.HandleFunc("/products/delete", DeleteProduct)
	http.HandleFunc("/products", GetProduct)

	err := http.ListenAndServe(":8000", nil)

	if err != nil{
		log.Fatal(err)
	}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var prd models.Product

		id := r.URL.Query().Get("id")

		if id == "" {
			utils.ResponseJSON(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}
		prd.ID, _ = strconv.Atoi(id)

		if err := product.Delete(ctx, prd); err != nil {

			kesalahan := map[string]string{
				"error": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Succesfully",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT"{
		if r.Header.Get("Content-Type") != "application/json"{
			http.Error(w, "Gunakan content type application /json",
				http.StatusBadRequest)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var prd models.Product

		if err := json.NewDecoder(r.Body).Decode(&prd);err!= nil{
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		fmt.Println(prd)
		if err := product.Update(ctx, prd); err != nil{
			utils.ResponseJSON(w, err,http.StatusInternalServerError)
			return
		}
		res := map [string]string{
			"status" : "Succesfully",
		}
		utils.ResponseJSON(w, res,http.StatusCreated)
		return
	}
	http.Error(w, "Tidak diizinkan!", http.StatusMethodNotAllowed)
	return
}

//Menambahkan produk baru
func PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		if r.Header.Get("Content-Type") != "application/json"{
			http.Error(w, "Gunakan content type application /json",
				http.StatusBadRequest)
			return
		}
		ctx,cancel := context.WithCancel(context.Background())
		defer cancel()

		var prd models.Product

		if err:= json.NewDecoder(r.Body).Decode(&prd); err != nil{
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}
		if err := product.Insert(ctx,prd); err != nil{
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}
		res := map[string]string{
			"status":"succesfully",
		}
		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}
	http.Error(w, "Tidak diizinkan!", http.StatusMethodNotAllowed)
	return
}

//Menampilkan seluruh data product
func GetProduct(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{

		//untuk membatalkan seluruh proses ketika ada kesalahan
		ctx, cancel := context.WithCancel(context.Background())
		//Tutup koneksi database ketika proses selesai
		defer cancel()
		products, err := product.GetAll(ctx)
		if err != nil{
			fmt.Println(err)
		}
		utils.ResponseJSON(w, products, http.StatusOK)
		return
	}
	http.Error(w,"Tidak diizinkan!", http.StatusNotFound)
	return
}
