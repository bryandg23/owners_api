package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Owner struct {
	Id               primitive.ObjectID
	FirstName        string
	LastName         string
	Email            string
	Birthday         string
	AppartmentNumber int
	Contract         string
	PhoneNumber      string
}
