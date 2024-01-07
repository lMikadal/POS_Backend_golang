package seeds

import (
	"github.com/lMikadal/POS_Backend_golang.git/models"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	users := []models.User{
		{UserName: "admin", UserEmail: "admin@gmail.com", UserPassword: "admin", Role: []models.Role{{RoleName: "admin"}}},
		{UserName: "employee", UserEmail: "employee@gmail.com", UserPassword: "employee", Role: []models.Role{{RoleName: "employee"}}},
	}

	db.Save(&users)
}
