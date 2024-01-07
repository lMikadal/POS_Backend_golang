package seeds

import (
	"github.com/lMikadal/POS_Backend_golang.git/models"
	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []models.Role{
		{RoleName: "admin"},
		{RoleName: "employee"},
	}

	db.Create(&roles)
}
