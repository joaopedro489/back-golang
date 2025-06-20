package handlers

import "github.com/joaopedro489/back-golang/internal/core/users/handlers/controllers"

type UserHandlers struct {
	Add        *controllers.AddUserController
	Read       *controllers.ReadUserController
	Delete     *controllers.DeleteUserController
	Edit       *controllers.EditUserController
	Browse     *controllers.BrowseUserController
	GetDetails *controllers.GetDetailsController
}

func NewUserHandlers(
	add *controllers.AddUserController,
	read *controllers.ReadUserController,
	delete *controllers.DeleteUserController,
	edit *controllers.EditUserController,
	browse *controllers.BrowseUserController,
	getDetails *controllers.GetDetailsController,
) *UserHandlers {
	return &UserHandlers{
		Add:        add,
		Read:       read,
		Delete:     delete,
		Edit:       edit,
		Browse:     browse,
		GetDetails: getDetails,
	}
}
