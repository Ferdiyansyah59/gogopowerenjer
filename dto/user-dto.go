package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password" binding:"required,min=8,omitempty"`
}

// type UserCreateDTO struct {
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
// 	Password string `json:"password" form:"password" binding:"required" validate:"min:8"`
// }
