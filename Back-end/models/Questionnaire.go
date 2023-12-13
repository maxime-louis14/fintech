package models

import "time"

type Questionnaire struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string    `json:"Title" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Answer    string    `json:"answer" validate:"required"`
	Question  string    `json:"question" validate:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
