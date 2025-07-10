package client

import "github.com/Mateus-MS/Gole-Certo/dev/backend/domain/client/fields"

type BaseClient struct {
	Email   fields.Email
	Phone   fields.Phone
	Address string
}

type Client interface {
}
