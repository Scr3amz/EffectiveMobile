package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"strconv"

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
		default:
			http.Error(w, fmt.Sprintf("expect method POST, GET or PUT at /fios/, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	} else {
		path := strings.Trim(req.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(w, "expect /fios/<id>", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathParts[1])
		if err != nil {
			// TODO: испарвить
			return
		}
		switch req.Method {
		case http.MethodDelete:
			h.RemoveFio(w, req, id)
		}
	}
}
