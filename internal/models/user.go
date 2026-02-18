package models

import "github.com/gofiber/fiber/v3"

type User struct {
	ID       int32
	Name     string
	Email    string
	Password string
	Admin    bool
}

func UserFromContext(ctx fiber.Ctx) *User {
	u := new(User)
	err := ctx.Bind().Body(u)
	if err != nil {
		return &User{}
	}
	return u
}
