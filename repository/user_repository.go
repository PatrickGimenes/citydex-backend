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

func (ur *UserRepository) UpdateUser(id_user string, user model.User) error {

	hashedPassword, err := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	if err != nil {
		log.Println(err)
		return err
	}
	query := "UPDATE USERS SET username = $1, password = $2, name = $3, home_city = (SELECT id FROM CITIES WHERE name = $4), update_At = CURRENT_DATE WHERE id = $5;"

	_, err = ur.connection.Exec(query, user.Username, user.Password, user.Name, user.Home_city, id_user)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserById(id_user string) (*model.UserResponse, error) {

	query, err := ur.connection.Prepare("SELECT id, username, name, home_city, create_at, update_at FROM users WHERE id = $1")

	if err != nil {
		log.Print(err)
		return nil, err
	}

	var user model.UserResponse

	err = query.QueryRow(id_user).Scan(
		&user.Id,
		&user.Username,
		&user.Name,
		&user.Home_city,
		&user.Create_At,
		&user.Update_At,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	query.Close()
	return &user, nil
}
