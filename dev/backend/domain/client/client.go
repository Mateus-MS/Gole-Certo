package client

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"

type BaseClient struct {
	Email   fields.Email `json:"Email"`
	Phone   fields.Phone `json:"Phone"`
	Address string       `json:"Address"`
}

type Client interface {
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
