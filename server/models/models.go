package models

import (
	"time"

	"github.com/kamva/mgm/v3"
	// "github.com/satori/go.uuid"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	mgm.DefaultModel     `bson:",inline"`
	FirstName            string    `json:"firstName"`
	LastName             string    `json:"lastName"`
	TimeTrackingSessions []Session `json:"timeTrackingSessions"`
}

type Session struct {
	UUID  string    `json:"id"`
	Name  string    `json:"name"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func NewUser(firstName, lastName string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
	}
}
