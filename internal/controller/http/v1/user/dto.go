package user

type CreateUserDTO struct {
	FirstName string `json:"firstName" validate:"required" example:"Vasya"`
	LastName  string `json:"lastName" validate:"required" example:"Pupkin"`
	Email     string `json:"email" validate:"required,email" example:"Vasya@gmail.com"`
	Password  string `json:"password" validate:"required,min=6" example:"123456"`
}
