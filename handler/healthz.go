package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
//
//	type HealthzHandler struct {
//		Message string `json:"message"`
//	}
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{
		message :"OK"
	}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// message := model.Message{
	// 	message: "OK"
	// }
	m := &model.HealthzResponse{
		Message: "OK",
	}
	err := json.NewEncoder(w)
	log.Printf(err)
	// if err != nil {
	// 	// http.Error(w, err.Error(), 500)
	// 	log.Printf(err)
	// 	return
	// }
	http.HandleFunc("/healthz", handler)
	w.WriteHeader(http.StatusOK)
	w.Write(m)
}
