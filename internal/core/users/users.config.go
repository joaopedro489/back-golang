package users

import (
	"github.com/joaopedro489/back-golang/internal/core/users/handlers"
	"github.com/joaopedro489/back-golang/internal/core/users/handlers/controllers"
	"github.com/joaopedro489/back-golang/internal/core/users/repository"
	usecases "github.com/joaopedro489/back-golang/internal/core/users/use-cases"
	"gorm.io/gorm"
)

func Config(db *gorm.DB) *handlers.UserHandlers {
	userRepo := repository.NewUserRepository(db)

	addUserUC := usecases.NewAddUser(userRepo)
	addUser := controllers.NewUserController(*addUserUC)

	readUserUC := usecases.NewReadUser(userRepo)
	readUser := controllers.NewReadUserController(*readUserUC)

	deleteUserUC := usecases.NewDeleteUser(userRepo)
	deleteUser := controllers.NewDeleteUserController(*deleteUserUC)

	browseUserUC := usecases.NewBrowseUser(userRepo)
	browseUser := controllers.NewBrowseUserController(*browseUserUC)

	editUserUC := usecases.NewEditUser(userRepo)
	editUser := controllers.NewEditUserController(*editUserUC)

	getDetailsUC := usecases.NewGetDetails(userRepo)
	getDetails := controllers.NewGetDetailsController(*getDetailsUC)

	userHandler := &handlers.UserHandlers{
		Add:        addUser,
		Read:       readUser,
		Delete:     deleteUser,
		Browse:     browseUser,
		Edit:       editUser,
		GetDetails: getDetails,
	}

	return userHandler
}
