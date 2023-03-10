package handlers

import "net/http"

// ProductIndex displays all products in the system
func (h *handler) ProductIndex(w http.ResponseWriter, r *http.Request) {
	p := h.repo.GetAllProducts()
	// Send an HTTP status & send the slice
	writeResponse(w, http.StatusOK, p, nil)
}
