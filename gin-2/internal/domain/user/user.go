package user

type Roles string

const (
	RoleAdmin      Roles = "admin"
	RoleSuperAdmin Roles = "super-admin"
)

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email" `
	Password string `json:"password" `
	Role     Roles  `json:"role" `
}
