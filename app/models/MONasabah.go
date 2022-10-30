package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Nasabah struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Cif      int                `json:"cif,omitempty" validate:"required"`
	Nama     string             `json:"nama,omitempty" validate:"required"`
	NoHp     int                `json:"noHp,omitempty" validate:"required"`
	Email    string             `json:"email,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
