package models

import "time"

type Questionnaire struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Answer    string    `json:"answer" validate:"required"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
