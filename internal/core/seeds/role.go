package seeds

import (
	"github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"
	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []domain.Role{
		{RoleName: "admin"},
		{RoleName: "employee"},
	}

	db.Save(&roles)
}
