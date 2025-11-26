package user_test

import (
	"testing"

	user "github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model"
	"github.com/Mateus-MS/Gole-Certo/dev/backend/modules/user/model/fields"
	"github.com/stretchr/testify/assert"
)

func TestNewCompanyValid(t *testing.T) {
	cnpj := "12345678000195"
	email := "company@example.com"
	fantasyName := "My Business"
	legalName := "My Business LTDA"
	phones := []string{"911911911"}
	addresses := []string{"123 Business St"}
	contactNames := []string{"John Doe doe"}

	company, err := user.NewCompany(cnpj, email, fantasyName, legalName, phones, addresses, contactNames)
	assert.NoError(t, err, "Expected no error on valid company")

	assert.Equal(t, cnpj, company.CNPJ.Get(), "CNPJ mismatch")
	assert.Equal(t, "company", company.Type, "Expected type to be 'company'")
	assert.Equal(t, email, company.Email.Get(), "Email mismatch")
	assert.Equal(t, fantasyName, company.FantasyName, "Fantasy name mismatch")
	assert.Equal(t, legalName, company.LegalName, "Legal name mismatch")
}

func TestNewCompanyInvalidCNPJ(t *testing.T) {
	cnpj := "invalid-cnpj"
	email := "company@example.com"
	fantasyName := "My Business"
	legalName := "My Business LTDA"
	phones := []string{"911911911"}
	addresses := []string{"123 Business St"}
	contactNames := []string{"John Doe doe"}

	_, err := user.NewCompany(cnpj, email, fantasyName, legalName, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidCNPJ, "Expected invalid CNPJ")
}

func TestNewCompanyInvalidEmail(t *testing.T) {
	cnpj := "12345678000195"
	email := "invalid-email"
	fantasyName := "My Business"
	legalName := "My Business LTDA"
	phones := []string{"911911911"}
	addresses := []string{"123 Business St"}
	contactNames := []string{"John Doe doe"}

	_, err := user.NewCompany(cnpj, email, fantasyName, legalName, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidEmail, "Expected invalid email")
}

func TestNewCompanyInvalidPhone(t *testing.T) {
	cnpj := "12345678000195"
	email := "company@example.com"
	fantasyName := "My Business"
	legalName := "My Business LTDA"
	phones := []string{"invalid-phone"}
	addresses := []string{"123 Business St"}
	contactNames := []string{"John Doe doe"}

	_, err := user.NewCompany(cnpj, email, fantasyName, legalName, phones, addresses, contactNames)
	assert.ErrorIs(t, err, fields.ErrInvalidPhone, "Expected invalid phone")
}

func TestNewCompanyInvalidContactName(t *testing.T) {
	cnpj := "12345678000195"
	email := "company@example.com"
	fantasyName := "My Business"
	legalName := "My Business LTDA"
	phones := []string{"911911911"}
	addresses := []string{"123 Business St"}
	contactNames := []string{"ja"}

	_, err := user.NewCompany(cnpj, email, fantasyName, legalName, phones, addresses, contactNames)
	assert.ErrorIs(t, err, user.ErrInvalidName, "Expected invalid contact name")
}

func TestNewCompanyInvalidFantasyName(t *testing.T) {
	cnpj := "12345678000195"
	email := "company@example.com"
	fantasyName := ""
	legalName := "My Business LTDA"
	phones := []string{"911911911"}
	addresses := []string{"123 Business St"}
	contactNames := []string{"John Doe doe"}

	_, err := user.NewCompany(cnpj, email, fantasyName, legalName, phones, addresses, contactNames)
	assert.ErrorIs(t, err, user.ErrInvalidName, "Expected invalid fantasy name")
}

func TestNewCompanyInvalidLegalName(t *testing.T) {
	cnpj := "12345678000195"
	email := "company@example.com"
	fantasyName := "My Business"
	legalName := ""
	phones := []string{"911911911"}
	addresses := []string{"123 Business St"}
	contactNames := []string{"John Doe doe"}

	_, err := user.NewCompany(cnpj, email, fantasyName, legalName, phones, addresses, contactNames)
	assert.ErrorIs(t, err, user.ErrInvalidName, "Expected invalid legal name")
}
