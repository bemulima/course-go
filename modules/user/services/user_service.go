package services

import (
	"course-go/modules/user/models"
	"course-go/modules/user/repositories"
)

type UserService interface {
	CreateUser(user *models.User) (models.User, error)
	GetUser(id int) (models.User, error)
	UpdateUser(user *models.User, id int) (models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(user *models.User) (models.User, error) {
	if err := s.repo.Create(user); err != nil {
		return models.User{}, err // Return empty user and the error
	}
	return *user, nil
}

func (s *userService) GetUser(id int) (models.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (s *userService) UpdateUser(user *models.User, id int) (models.User, error) {
	if err := s.repo.Update(user); err != nil {
		return models.User{}, err // Return empty user and the error
	}
	return *user, nil
}
