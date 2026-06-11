package handlers

import (
	"net/http"
	"strconv"

	"github.com/Wai-Thura-Tun/microservices-ecomm/data"
	"github.com/gorilla/mux"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Return noContent
// responses:
// 	201: noContent

// DeleteProducts deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle DELETE request")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Request URL is invalid", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "Not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
