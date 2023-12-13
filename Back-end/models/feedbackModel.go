package models

import "time"

type Feedback struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Nom       string    `json:"nom" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Reponse   string    `json:"reponse" validate:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
