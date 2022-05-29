package auth_service

import "github.com/jamesluo111/gin-blog/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckoutAuth(a.Username, a.Password)
}
