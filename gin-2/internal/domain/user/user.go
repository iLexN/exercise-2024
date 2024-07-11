package user

import (
	"payment-portal/internal/model"
	"payment-portal/internal/paginator"
)

type Roles string

const (
	RoleAdmin      Roles = "admin"
	RoleSuperAdmin Roles = "super-admin"
	Roleuser       Roles = "super-admin"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email" `
	Password string `json:"password" `
	Role     Roles  `json:"role" `
}

type ListResult struct {
	Users     []model.User `json:"users"`
	Paginator paginator.LengthAwarePaginator
}
