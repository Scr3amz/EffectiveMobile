package handlers

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"

	"github.com/Scr3amz/EffectiveMobile/internal/api"
	"github.com/Scr3amz/EffectiveMobile/internal/database/models"
)

func (h *Handlers) CreateFio(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json content-type", http.StatusUnsupportedMediaType)
		return
	}
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var fio models.FIO
	if err := dec.Decode(&fio); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = api.FillTheMessage(&fio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fioID, err := h.store.FioStorer.Add(fio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "fio with id:%v added successfully", fioID)
}

func (h *Handlers) ListFio(w http.ResponseWriter, req *http.Request) {
	fios, err := h.store.FioStorer.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, fios)
}

func (h *Handlers) UpdateFio(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json content-type", http.StatusUnsupportedMediaType)
		return
	}
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var fio models.FIO
	if err := dec.Decode(&fio); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fioID, err := h.store.FioStorer.Update(fio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "fio with id:%v updated successfully", fioID)
}

func (h *Handlers) RemoveFio(w http.ResponseWriter, req *http.Request, id int) {
	err := h.store.FioStorer.Remove(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "fio with id:%v deleted successfully", id)
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
