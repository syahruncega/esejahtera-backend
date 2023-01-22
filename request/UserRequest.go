package request

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"require"`
	NoHp     string `json:"noHp" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	NoHp     string `json:"noHp"`
	Role     string `json:"role"`
}
