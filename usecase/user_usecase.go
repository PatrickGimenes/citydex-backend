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
		Id:        userId,
		Username:  user.Username,
		Name:      user.Name,
		Home_city: user.Home_city,
	}
	return userResponse, nil
}

func (uu *UserUsecase) GetUserById(id_user string) (*model.UserResponse, error) {
	user, err := uu.repository.GetUserById(id_user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
