package postgresql

import (
	"fmt"

	"github.com/Scr3amz/EffectiveMobile/config"
	"github.com/Scr3amz/EffectiveMobile/internal/database"
	"github.com/Scr3amz/EffectiveMobile/internal/database/models"
	"github.com/Scr3amz/EffectiveMobile/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresqlStore struct {
	Fios FioStore
}

// NewStore creates postgreSQL database and returns Store structure.
func NewStore(config config.Config, logger logger.Logger) *database.Store {
	db := createDB(config, logger)

	postgresqlStore := postgresqlStore{
		Fios: FioStore{
			DB:     *db,
			Logger: logger,
		},
	}
	store := database.Store{
		FioStorer: postgresqlStore.Fios,
	}

	return &store
}

// CreateDB creates, setup and returns DB structure by dsn.
// It will automigrate all the structures to tables in DB.
// To create DB you need a Config file host of DB, port of DB,
// name of DB, username and password from DB. This function will
// panic if failed to create or migrate.
func createDB(config config.Config, logger logger.Logger) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.HostDB,
		config.UserDB,
		config.PasswordDB,
		config.NameDB,
		config.PortDB,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ErrorLog.Println("failed to connect to postgresql database")
		panic(err)
	}
	setupDB(db, logger)
	return db
}

// setupDB migrates database. Panic if failed
func setupDB(db *gorm.DB, logger logger.Logger) {
	err := db.AutoMigrate(&models.FIO{})

	if err != nil {
		logger.ErrorLog.Println("failed to migrate tables")
		panic(err)
	}
}
