package usecase

import (
	"citydex/model"
	"citydex/repository"
)

type UserUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return UserUsecase{
		repository: repo,
	}
}

func (uu *UserUsecase) GetUsers() ([]model.User, error) {
	return uu.repository.GetUsers()
}
