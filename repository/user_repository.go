package repository

import (
	"citydex/model"
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
