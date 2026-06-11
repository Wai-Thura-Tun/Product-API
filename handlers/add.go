package handlers

import (
	"net/http"

	"github.com/Wai-Thura-Tun/microservices-ecomm/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle post product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
