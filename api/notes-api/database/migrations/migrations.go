package migrations

import (
	"github.com/matheussASM/pre-atividade-02/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Note{})
}
