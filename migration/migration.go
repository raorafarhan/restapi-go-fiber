package migration

import (
	userModel "restapi-gofiber/features/users/repository"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})

}
