package models

type User struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

type Register struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
