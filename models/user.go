package models

import (
	"time"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

//User: Datos del usuario
type User struct {
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"ID, omitempty"`
	Name      string    `json:"Name"`
	Email     string    `json:"Email"`
	CreatedAt time.Time `bson:"created_At" json:"CreateAt"`
	UpdatedAt time.Time `bson:"updated_At" json:"UpdatedAt,omitempty"`
}

//Users: Lista de usuarios
type Users []*User
