package models

import "time"

type Response struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string    `json:"firstName" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Answer    string    `json:"answer" validate:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
