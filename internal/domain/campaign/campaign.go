package campaign

import (
	"emailn/internal/domain"
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	Id        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	CreatedOn time.Time `validate:"required"`
}

type Contact struct {
	Value string `validate:"email|e164"`
}

func NewCampaign(name string, content string, rawContacts []string) (*Campaign, error) {
	contacts := make([]Contact, len(rawContacts))
	for index, rawContact := range rawContacts {
		contacts[index].Value = rawContact
	}

	campaign := &Campaign{
		Id:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedOn: time.Now(),
	}

	err := domain.Validate(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}
