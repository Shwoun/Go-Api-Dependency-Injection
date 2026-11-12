package service

import (
	"Di-Test/internal/model"
	"Di-Test/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) CreateUser(user model.User) error {
	return s.Repo.CreateUser(user)
}
