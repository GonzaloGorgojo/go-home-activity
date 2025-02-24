package users

import (
	"database/sql"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{Db: db}
}

func (r *UserRepositoryImpl) GetAllUsers() ([]User, error) {
	rows, err := r.Db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Type); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
