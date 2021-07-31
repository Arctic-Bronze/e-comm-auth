package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserSignupModel struct {
	Email    string
	Password string
	Name     string
	LastName string
}

func (u *UserSignupModel) ToUser() *User {
	return &User{
		Id:           primitive.NewObjectID(),
		Email:        u.Email,
		Name:         u.Name,
		LastName:     u.LastName,
		FullName:     u.Name + " " + u.LastName,
		Status:       "EMAIL_NOT_VERIFIED",
		Role:         []Role{"CONSUMER"},
		CreationTime: time.Now(),
	}
}
