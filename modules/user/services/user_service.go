package services

import (
	"course-go/modules/user/interfaces"
	"course-go/modules/user/models"
)

type userService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) interfaces.UserService {
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
