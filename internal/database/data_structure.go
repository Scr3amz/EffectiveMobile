package database

import (
	"net/http"

	"github.com/Scr3amz/EffectiveMobile/internal/database/models"
)

type FioStorer interface {
	Add(models.FIO) (int, error)
	List() ([]models.FIO, error)
	Update(models.FIO) (int, error)
	Remove(fioID int) error
	ListWithPagination(req *http.Request) ([]models.FIO, error)
}

type Store struct {
	FioStorer
}
