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
	// return &HealthzHandler{
	// 	message :"OK"
	// }
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// message := model.Message{
	// 	message: "OK"
	// }
	m := &model.HealthzResponse{
		Message: "OK",
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(m); err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// m := &model.HealthzResponse{
	// 	Message: "OK",
	// }
	// err := json.NewEncoder(w)
	// log.Printf(err)

	// http.HandleFunc("/healthz", handler)
	// w.WriteHeader(http.StatusOK)
	// w.Write(m)
}
