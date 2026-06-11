// Package classification Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Wai-Thura-Tun/microservices-ecomm/data"
)

// A list of products returns in the response
// swagger:response productsResponse
type ProductsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type ProductsNoContent struct {
}

// swagger:parameters deleteProduct
type ProductIDParameterWrapper struct {
	// The ID of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l: l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		p.l.Println("PUT", r.URL.Path)
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		matches := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(matches) != 1 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		if len(matches[0]) != 2 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		idString := matches[0][1]
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			http.Error(rw, "Something went wrong", http.StatusInternalServerError)
// 			return
// 		}

// 		p.l.Println("got id", id)

// 		p.updateProduct(id, rw, r)
// 		return
// 	}

// 	// catch all
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

// Middleware Section

type KeyProduct struct{}

func (p Products) MiddlewareProductValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)

		// deserializing product
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
			return
		}

		// add the product to the product list
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
