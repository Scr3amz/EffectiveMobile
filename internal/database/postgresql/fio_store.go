package postgresql

import (
	"log"

	"github.com/Scr3amz/EffectiveMobile/internal/database/models"
	"gorm.io/gorm"
)

type FioStore struct {
	DB gorm.DB
}

func (s FioStore) Add(fio models.FIO) (int, error) {
	resault := s.DB.Create(&fio)
	if resault.Error != nil {
		log.Println("Unable to add FIO to postgresql table")
		return -1, resault.Error
	}
	return fio.ID, nil
}

func (s FioStore) List() ([]models.FIO, error) {
	fios := []models.FIO{}
	resault := s.DB.Find(&fios)
	if resault.Error != nil {
		log.Println("Unable to show FIOs from table")
		return nil, resault.Error
	}
	if resault.RowsAffected == 0 {
		log.Println("FIOs table is empty")
		return nil, nil
	}
	return fios, nil

}

func (s FioStore) Update(fio models.FIO) (int, error) {
	resault := s.DB.First(&models.FIO{ID: fio.ID})
	if resault.Error != nil {
		log.Println("Unable to find account in postgresql table")
		return -1, resault.Error
	}
	resault = s.DB.Save(&fio)
	if resault.Error != nil {
		log.Println("Unable to edit account in postgresql table")
		return -1, resault.Error
	}
	log.Printf("FIO with id:%v updated", fio.ID)
	return fio.ID, nil
}

func (s FioStore) Remove(fioID int) error {
	resault := s.DB.First(&models.FIO{ID: fioID})
	if resault.Error != nil {
		log.Println("Unable to find account in postgresql table")
		return resault.Error
	}
	resault = s.DB.Delete(&models.FIO{}, fioID)
	if resault.Error != nil {
		log.Printf("Unable to delete account with id %v in postgresql table\n", fioID)
		return resault.Error
	}
	log.Printf("FIO with id:%v deleted", fioID)
	return nil
}