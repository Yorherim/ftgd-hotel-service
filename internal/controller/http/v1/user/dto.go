package user

type CreateUserDTO struct {
	FirstName string `bson:"firstName" json:"firstName" validate:"required" example:"Vasya"`
	LastName  string `bson:"lastName" json:"lastName" validate:"required" example:"Pupkin"`
	Email     string `bson:"email" json:"email" validate:"required,email" example:"Vasya@gmail.com"`
	Password  string `bson:"password" json:"password" validate:"required,min=6" example:"123456"`
}
