package database

import "github.com/Scr3amz/EffectiveMobile/internal/database/models"

type FioStorer interface {
	Add(models.FIO) (int, error)
	List() ([]models.FIO, error)
	Update(models.FIO) (int, error)
	Remove(fioID int) error
}

type Store struct {
	FioStorer
}
