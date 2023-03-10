package compaign

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
	var _error error

	if *name == "" {
		_error = errors.New("Name is required")
	}

	if *content == "" {
		_error = errors.New("Content is required")
	}

	if len(*rawContacts) == 0 {
		_error = errors.New("Contacts are required")
	}

	return _error
}
