package user

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/user/fields"

type BaseUser struct {
	Type         string         `json:"Type"        bson:"type"` // I'm not a fan of using a `type` field here, but till i think in something better, will be like this
	Email        fields.Email   `json:"Email"       bson:"email"`
	Phone        []fields.Phone `json:"Phone"       bson:"phone"`
	Address      []string       `json:"Address"     bson:"address"`
	ContactNames []fields.Name  `json:"ContactName" bson:"contactName"` // The first contact in the list is the main one
}

type User interface {
	GetIdentifier() string
	IsValid() error
}

func NewBaseUser(emailRaw string, phonesRaw, address, contactNamesRaw []string) (BaseUser, error) {
	var (
		usr BaseUser
		err error

		email        fields.Email
		phones       []fields.Phone
		contactNames []fields.Name
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
		var contactName fields.Name
		if contactName, err = fields.NewName(contactNameRaw); err != nil {
			return usr, fields.ErrInvalidName
		}
		contactNames = append(contactNames, contactName)
	}

	return BaseUser{
		Email:        email,
		Phone:        phones,
		Address:      address,
		ContactNames: contactNames,
	}, nil
}

func (bc *BaseUser) GetMainContactName() string {
	return bc.ContactNames[0].Get()
}
