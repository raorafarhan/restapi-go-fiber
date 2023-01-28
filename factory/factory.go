package factory

import (
	"github.com/gofiber/fiber/v2"
	userControllers "restapi-gofiber/features/users/controllers"
	userRepository "restapi-gofiber/features/users/repository"
	userServices "restapi-gofiber/features/users/services"

	"gorm.io/gorm"
)

func InitFactory(e *fiber.App, db *gorm.DB) {
	userRepositoryFactory := userRepository.New(db)
	userServiceFactory := userServices.NewUsersServices(userRepositoryFactory)
	userControllers.New(e, userServiceFactory)

}
