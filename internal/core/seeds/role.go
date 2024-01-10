package seeds

import (
	repository "github.com/lMikadal/POS_Backend_golang.git/repositories"
	"gorm.io/gorm"
)

func SeedRole(db *gorm.DB) {
	roles := []repository.Role{
		{RoleName: "admin"},
		{RoleName: "employee"},
	}

	db.Save(&roles)
}
