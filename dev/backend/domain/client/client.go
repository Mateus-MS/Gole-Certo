package client

import (
	"github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"
)

type BaseClient struct {
	Type         string         `json:"Type"        bson:"type"` // I'm not a fan of using a `type` field here, but till i think in something better, will be like this
	Email        fields.Email   `json:"Email"       bson:"email"`
	Phone        []fields.Phone `json:"Phone"       bson:"phone"`
	Address      []string       `json:"Address"     bson:"address"`
	ContactNames []fields.Name  `json:"ContactName" bson:"contactName"` // The first contact in the list is the main one
}

type Client interface {
	GetIdentifier() string
	IsValid() error
}

func NewBaseClient(emailRaw string, phonesRaw, address, contactNamesRaw []string) (BaseClient, error) {
	var (
		client BaseClient
		err    error

		email        fields.Email
		phones       []fields.Phone
		contactNames []fields.Name
	)

	// Validate email
	if email, err = fields.NewEmail(emailRaw); err != nil {
		return client, fields.ErrInvalidEmail
	}

	// Validate all phones
	for _, phoneRaw := range phonesRaw {
		var phone fields.Phone
		if phone, err = fields.NewPhone(phoneRaw); err != nil {
			return client, fields.ErrInvalidPhone
		}
		phones = append(phones, phone)
	}

	// Validate all contact names
	for _, contactNameRaw := range contactNamesRaw {
		var contactName fields.Name
		if contactName, err = fields.NewName(contactNameRaw); err != nil {
			return client, fields.ErrInvalidName
		}
		contactNames = append(contactNames, contactName)
	}

	return BaseClient{
		Email:        email,
		Phone:        phones,
		Address:      address,
		ContactNames: contactNames,
	}, nil
}

func (bc *BaseClient) GetMainContactName() string {
	return bc.ContactNames[0].Get()
}
