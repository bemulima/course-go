package interfaces

import "course-go/modules/user/models"

type UserService interface {
	CreateUser(user *models.User) (models.User, error)
	GetUser(id int) (models.User, error)
	UpdateUser(user *models.User, id int) (models.User, error)
}
