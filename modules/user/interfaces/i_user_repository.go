package interfaces

import (
	"course-go/modules/user/models"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
	FindById(id int) (*models.User, error)
	Update(user *models.User) error
}
