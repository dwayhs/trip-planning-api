package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

type ListTripsHandler struct{}

func (h *ListTripsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List trips")
}

type GetTripHandler struct{}

func (h *GetTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Get trip: %s", params["id"])
}

type PutTripHandler struct{}

func (h *PutTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Update trip: %s", params["id"])
}

type PostTripHandler struct{}

func (h *PostTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create trip")
}

type DeleteTripHandler struct{}

func (h *DeleteTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Delete trip: %s", params["id"])
}

func main() {
	addr := "127.0.0.1:8080"
	router := httptreemux.NewContextMux()

	router.Handler(http.MethodGet, "/trip", &ListTripsHandler{})
	router.Handler(http.MethodPost, "/trip", &PostTripHandler{})
	router.Handler(http.MethodGet, "/trip/:id", &GetTripHandler{})
	router.Handler(http.MethodPut, "/trip/:id", &PutTripHandler{})
	router.Handler(http.MethodDelete, "/trip/:id", &DeleteTripHandler{})

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
