package services

import (
	"course-go/modules/user/models"
	"course-go/modules/user/repositories"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(id int) (models.User, error)
	UpdateUser(user models.User, id int) (models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetUser() (user models.User, err error) {
	return s.userRepo.GetUser()
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUser(id int) (models.User, error) {
	return s.userRepo.GetUser(id)
}

func (s *userService) UpdateUser(user models.User, id int) (models.User, error) {
	return s.userRepo.UpdateUser(user, id)
}
