package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Role string
type Status string

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	Email        string
	Name         string
	LastName     string
	FullName     string
	Status       Status
	Role         []Role
	CreationTime time.Time
}
