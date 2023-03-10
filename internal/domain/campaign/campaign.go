package compaign

import (
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	Id        string
	Name      string
	Content   string
	Contacts  []Contact
	CreatedOn time.Time
}

type Contact struct {
	Value string
}

func NewCampaign(name string, content string, rawContacts []string) *Campaign {
	contacts := make([]Contact, len(rawContacts))
	for index, rawContact := range rawContacts {
		contacts[index].Value = rawContact
	}

	return &Campaign{
		Id:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedOn: time.Now(),
	}
}
