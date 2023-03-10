package compaign

import "time"

type Campaign struct {
	Id        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contact
}

type Contact struct {
	Email string
}
