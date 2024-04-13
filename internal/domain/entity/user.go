package entity

type User struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty" example:"660c9b35208b7ec8f81791e6"`
	FirstName string `bson:"firstName" json:"firstName" example:"Vasya"`
	LastName  string `bson:"lastName" json:"lastName" example:"Pupkin"`
	Email     string `bson:"email" json:"email" example:"vasya@gmail.com"`
	Password  string `bson:"password" json:"-"`
}
