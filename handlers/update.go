package handlers

import (
	"net/http"
	"strconv"

	"github.com/Wai-Thura-Tun/microservices-ecomm/data"
	"github.com/gorilla/mux"
)

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT request")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Request URL is invalid", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}

}
