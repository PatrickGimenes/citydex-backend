package repository

import (
	"citydex/model"
	"citydex/utils"
	"database/sql"
	"log"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	query := "SELECT * FROM users"
	rows, err := ur.connection.Query(query)
	if err != nil {
		log.Print(err)
		return []model.User{}, nil
	}

	var userList []model.User
	var userObj model.User

	for rows.Next() {
		err = rows.Scan(
			&userObj.Id,
			&userObj.Username,
			&userObj.Password,
			&userObj.Name,
			&userObj.Home_city,
			&userObj.Create_At,
			&userObj.Update_At)

		if err != nil {
			log.Println(err)
			return []model.User{}, err
		}

		userList = append(userList, userObj)
	}
	rows.Close()
	return userList, nil
}

func (ur *UserRepository) CreateUser(user model.User) (string, error) {

	hashedPassword, err := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	if err != nil {
		log.Println(err)
		return "", err
	}

	var id string

	query, err := ur.connection.Prepare("INSERT INTO USERS (username, password, name, home_city, create_At, update_At) VALUES ($1, $2, $3, (SELECT id FROM CITIES WHERE name = $4), CURRENT_DATE, CURRENT_DATE) RETURNING id")

	if err != nil {
		log.Println(err)
		return "", err
	}

	err = query.QueryRow(user.Username, user.Password, user.Name, user.Home_city).Scan(&id)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return id, nil
}
