package seeds

import (
	"github.com/lMikadal/POS_Backend_golang.git/repositories"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	users := []repository.User{
		{UserName: "admin", UserEmail: "admin@gmail.com", UserPassword: "admin", RoleID: 1},
		{UserName: "employee", UserEmail: "employee@gmail.com", UserPassword: "employee", RoleID: 2},
	}

	db.Save(&users)
}
