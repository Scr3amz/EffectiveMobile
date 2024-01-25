package handlers

import (
	"fmt"
	"net/http"

	"github.com/Scr3amz/EffectiveMobile/config"
	"github.com/Scr3amz/EffectiveMobile/internal/database"
)

type Handlers struct {
	store  database.Store
	config config.Config
}

func NewFioHandler(s database.Store, c config.Config) *Handlers {
	return &Handlers{
		store:  s,
		config: c,
	}
}

func (h *Handlers) FioHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/fios/" {
		switch req.Method {
		case http.MethodPost:
			h.CreateFio(w, req)
		case http.MethodGet:
			h.ListFio(w, req)
		case http.MethodPut:
			h.UpdateFio(w, req)
		case http.MethodDelete:
			h.RemoveFio(w, req)
		default:
			http.Error(w, fmt.Sprintf("expect method POST, GET, PUT or DELETE at /fios/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}
