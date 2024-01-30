package handlers

import (
	"net/http"
	"regexp"

	"github.com/Scr3amz/EffectiveMobile/config"
	"github.com/Scr3amz/EffectiveMobile/internal/database"
	"github.com/Scr3amz/EffectiveMobile/logger"
)

var (
	FiosReg           = regexp.MustCompile(`^/fios/$`)
	FiosRegWithID     = regexp.MustCompile(`^/fios/[0-9]+$`)
	FiosRegWithParams = regexp.MustCompile(`^/fios\?\w+=\w+(&\w+=\w+)*$`)
)

type Handlers struct {
	store  database.Store
	config config.Config
	logger logger.Logger
}

func NewFioHandler(s database.Store, c config.Config, l logger.Logger) *Handlers {
	return &Handlers{
		store:  s,
		config: c,
		logger: l,
	}
}

func (h *Handlers) FioHandler(w http.ResponseWriter, req *http.Request) {
	switch {
	case req.Method == http.MethodPost && FiosReg.MatchString(req.URL.Path):
		h.CreateFio(w, req)
		return
	case req.Method == http.MethodGet:
		h.ListFioWithPagination(w, req)
		return
	case req.Method == http.MethodPut && FiosReg.MatchString(req.URL.Path):
		h.UpdateFio(w, req)
		return
	case req.Method == http.MethodDelete && FiosRegWithID.MatchString(req.URL.Path):
		h.RemoveFio(w, req)
		return
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
}
