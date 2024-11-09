package repository

import (
	"database/sql"
	"inventaris/model"
)


type UserRepository interface {
	Create(user *model.User) error
	GetAll(user *[]model.User) error
	GetUserLogin(user model.User) (*model.User, error)
	UsernameExists(username string) (bool, error)
}

type UserRepoDb struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &UserRepoDb{DB: db}
}

func (r *UserRepoDb) Create(user *model.User) error {
	query := `INSERT INTO "Users" (username, password, role) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, user.Username, user.Password, user.Role).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepoDb) GetAll(users *[]model.User) error {
	query := `SELECT user_id, username, password, role FROM "User" WHERE role != 'admin'`
	rows, err := r.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role); err != nil {
			return err
		}
		*users = append(*users, user)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func (r *UserRepoDb) GetUserLogin(user model.User) (*model.User, error) {
	query := `SELECT id, username, password, role FROM "Users" WHERE username=$1 AND password=$2`
	var userResponse model.User
	err := r.DB.QueryRow(query, user.Username, user.Password).Scan(&userResponse.ID, &userResponse.Username, &userResponse.Password, &userResponse.Role)

	if err != nil {
		return nil, err
	}

	return &userResponse, nil
}

func (r *UserRepoDb) UsernameExists(username string) (bool, error) {
    query := `SELECT COUNT(*) FROM "Users" WHERE username = $1`
    var count int
    err := r.DB.QueryRow(query, username).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

