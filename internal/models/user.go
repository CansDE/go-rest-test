package models

type User struct {
	ID       int32
	Name     string
	Email    string
	Password string
	Admin    bool
}

func NewUser(name, email, password string, admin bool) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Admin:    admin,
	}
}
