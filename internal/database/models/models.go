package models

type FIO struct {
	ID          int      `gorm:"primary"`
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}
