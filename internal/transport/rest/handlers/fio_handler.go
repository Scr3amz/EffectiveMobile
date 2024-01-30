package handlers

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"strconv"
	"strings"

	"github.com/Scr3amz/EffectiveMobile/internal/api"
	"github.com/Scr3amz/EffectiveMobile/internal/database/models"
)

func (h *Handlers) CreateFio(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		h.logger.WarningLog.Println("occured problem with mediatype")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		h.logger.WarningLog.Println("got wrong mediatype")
		http.Error(w, "expect application/json content-type", http.StatusUnsupportedMediaType)
		return
	}
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var fio models.FIO
	if err := dec.Decode(&fio); err != nil {
		h.logger.WarningLog.Println("failed to decode fio")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if fio.Name == "" || fio.Surname == "" {
		h.logger.WarningLog.Println("got fio without name or surname")
		http.Error(w, `data must contain fields "name" and "surname"`, http.StatusBadRequest)
		return
	}
	err = api.FillTheMessage(&fio, h.logger)
	if err != nil {
		h.logger.WarningLog.Println("failed to get data from external api")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fioID, err := h.store.FioStorer.Add(fio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.logger.InfoLog.Printf("fio with id:%v added successfully\n", fioID)
	fmt.Fprintf(w, "fio with id:%v added successfully\n", fioID)
}

func (h *Handlers) ListFio(w http.ResponseWriter, req *http.Request) {
	fios, err := h.store.FioStorer.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, fios)
}

func (h *Handlers) ListFioWithPagination(w http.ResponseWriter, req *http.Request) {
	fios, err := h.store.FioStorer.ListWithPagination(req)
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
		h.logger.WarningLog.Println("occured problem with mediatype")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		h.logger.WarningLog.Println("got wrong mediatype")
		http.Error(w, "expect application/json content-type", http.StatusUnsupportedMediaType)
		return
	}
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	var fio models.FIO
	if err := dec.Decode(&fio); err != nil {
		h.logger.WarningLog.Println("failed to decode fio")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if fio.Name == "" || fio.Surname == "" {
		h.logger.WarningLog.Println("got fio without name or surname")
		http.Error(w, `data must contain fields "name" and "surname"`, http.StatusBadRequest)
		return
	}
	fioID, err := h.store.FioStorer.Update(fio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.logger.InfoLog.Printf("fio with id:%v updated successfully\n", fioID)
	fmt.Fprintf(w, "fio with id:%v updated successfully\n", fioID)
}

func (h *Handlers) RemoveFio(w http.ResponseWriter, req *http.Request) {
	matches := strings.Split(req.URL.Path, "/")
	id, err := strconv.Atoi(matches[2])
	if err != nil {
		h.logger.WarningLog.Println("failed to get id of fio")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.store.FioStorer.Remove(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.logger.InfoLog.Printf("fio with id:%v deleted successfully\n", id)
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
