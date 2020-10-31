package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/dgrijalva/jwt-go"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"adress"`
	Token    string `json:"token"`
}

type User struct {
	id       primitive.ObjectID `bson:"_id"`
	Login    string             `bson:"login"`
	Password string             `bson:"password"`
	Token    string             `bson:"token"`
	Address  string             `bson:"adress"`
	Phone    string             `bson:"phone"`
	Email    string             `bson:"email"`
	Manager  string             `bson:"manager"`
	Role     string             `bson:"role"`
}

type Voc struct {
	id    primitive.ObjectID `bson:"_id"`
	Login string             `bson:"login"`
	Name  string             `bson:"name"`
	Voc   []string           `bson:"voc"`
}

type Rvoc struct {
	Request []string `json:"request"`
}
