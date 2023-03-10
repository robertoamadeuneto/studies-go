package campaign

import (
	"errors"
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

func NewCampaign(name string, content string, rawContacts []string) (*Campaign, error) {
	if _error := validateCampaignProperties(&name, &content, &rawContacts); _error != nil {
		return nil, _error
	}

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
	}, nil
}

func validateCampaignProperties(name *string, content *string, rawContacts *[]string) error {
	var err error

	if *name == "" {
		err = errors.New("Name is required")
	} else if *content == "" {
		err = errors.New("Content is required")
	} else if len(*rawContacts) == 0 {
		err = errors.New("Contacts are required")
	}

	return err
}
