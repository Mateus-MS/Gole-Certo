package client

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"

type BaseClient struct {
	// I'm not a fan of using a `type` field here, but till i think in something better, will be like this
	Type    string       `json:"Type"    bson:"type"`
	Email   fields.Email `json:"Email"   bson:"email"`
	Phone   fields.Phone `json:"Phone"   bson:"phone"`
	Address string       `json:"Address" bson:"address"`
}

type Client interface {
	GetIdentifier() string
}

func NewBaseClient(emailRaw, phoneRaw, address string) (BaseClient, error) {
	// Validate all validatable fields
	var (
		client BaseClient
		err    error

		email fields.Email
		phone fields.Phone
	)

	if email, err = fields.NewEmail(emailRaw); err != nil {
		return client, err
	}

	if phone, err = fields.NewPhone(phoneRaw); err != nil {
		return client, err
	}

	return BaseClient{
		Email:   email,
		Phone:   phone,
		Address: address,
	}, nil
}
