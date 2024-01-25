package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handlers) CreateFio(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "creating fio...")
}

func (h *Handlers) ListFio(w http.ResponseWriter, req *http.Request)   {
	fmt.Fprint(w, "showing all fios...")
}

func (h *Handlers) UpdateFio(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "updating fio...")
}

func (h *Handlers) RemoveFio(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "removing fio...")
}
