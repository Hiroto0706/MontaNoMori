package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func GetUserById(id int) (user *User, err error) {
	user = &User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}

func GetUserByEmailAndPassword(email, password string) (user User, err error) {
	cmd := `select id, uuid, name, email, password, created_at from users where email = $1 and password = $2`
	err = Db.QueryRow(cmd, email, password).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt)

	return user, err
}
