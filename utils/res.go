package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON (w http.ResponseWriter, p interface{}, status int){
	changeToByte, err := json.Marshal(p)
	w.Header().Set("Content-Type","application/json")

	if err != nil{
		http.Error(w, "Error", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(changeToByte))
}