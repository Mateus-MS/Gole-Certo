package user

import (
	"errors"

	"github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model/fields"
)

type BaseUser struct {
	Type         string         `json:"Type"        bson:"type"` // I'm not a fan of using a `type` field here, but till i think in something better, will be like this
	Email        fields.Email   `json:"Email"       bson:"email"`
	Phone        []fields.Phone `json:"Phone"       bson:"phone"`
	Address      []string       `json:"Address"     bson:"address"`
	ContactNames []string       `json:"ContactName" bson:"contactName"` // The first contact in the list is the main one
}

var ErrInvalidName = errors.New("invalid name")

func NewBaseUser(emailRaw string, phonesRaw, address, contactNamesRaw []string) (usr BaseUser, err error) {
	var (
		email        fields.Email
		phones       []fields.Phone
		contactNames []string
	)

	// Validate email
	if email, err = fields.NewEmail(emailRaw); err != nil {
		return usr, fields.ErrInvalidEmail
	}

	// Validate all phones
	for _, phoneRaw := range phonesRaw {
		var phone fields.Phone
		if phone, err = fields.NewPhone(phoneRaw); err != nil {
			return usr, fields.ErrInvalidPhone
		}
		phones = append(phones, phone)
	}

	// Validate all contact names
	for _, contactNameRaw := range contactNamesRaw {
		var contactName string
		if len(contactNameRaw) < 10 {
			return usr, ErrInvalidName
		}
		contactNames = append(contactNames, contactName)
	}

	// Validate address
	for _, adr := range address {
		if len(adr) < 10 {
			return usr, fields.ErrInvalidAddress
		}
	}

	return BaseUser{
		Email:        email,
		Phone:        phones,
		Address:      address,
		ContactNames: contactNames,
	}, nil
}

func (bc *BaseUser) GetMainContactName() string {
	return bc.ContactNames[0]
}
