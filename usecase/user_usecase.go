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

func (uu *UserUsecase) CreateUser(user model.User) (model.UserResponse, error) {
	userId, err := uu.repository.CreateUser(user)

	if err != nil {
		return model.UserResponse{}, err
	}

	userResponse := model.UserResponse{
		ID:       userId,
		Username: user.Username,
		Name:     user.Name,
		HomeCity: user.Home_city,
	}
	return userResponse, nil
}
