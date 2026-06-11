package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Wai-Thura-Tun/microservices-ecomm/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)

	p := handlers.NewProduct(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// create subrouter with get method
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", p.GetProducts)

	// create subrouter with put method
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", p.UpdateProduct)
	putRouter.Use(p.MiddlewareProductValidate)

	// create subrouter with post method
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", p.AddProduct)
	postRouter.Use(p.MiddlewareProductValidate)

	// create subrouter with delete method
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", p.DeleteProduct)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	sh := middleware.Redoc(ops, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yml", http.FileServer(http.Dir("./")))

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
