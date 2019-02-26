package utils

import (
	"assignment/user/app/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJsonRes(w http.ResponseWriter, res interface{}) {
	// Marshal provided interface into JSON structure
	uj, err := json.Marshal(res)
	if err != nil {
		WriteJsonErr(w, err)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(uj)
}

func WriteJsonErr(w http.ResponseWriter, err error) {
	uj, err := json.Marshal(model.Error{Error: err.Error()})
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "application/json")
	w.Write(uj)
}
