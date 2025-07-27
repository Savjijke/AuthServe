package models

type User struct {
	ID           int64  `db:id`
	Username     string `db:username`
	Email        string `db:email`
	PasswordHash string `db:password_hash`
}
