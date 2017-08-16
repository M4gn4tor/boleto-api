package citibank

import (
	"errors"
	"fmt"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/validations"
)

func citiValidateAgency(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		err := t.Agreement.IsAgencyValid()
		if err != nil {
			return err
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}

func citiValidateAccount(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		if len(t.Agreement.Account) != 9 {
			return errors.New(fmt.Sprintf("A conta deve conter somente 9 digítos."))
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}

func citiValidateAccountDigit(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		if len(t.Agreement.AccountDigit) < 1 || len(t.Agreement.AccountDigit) > 2 {
			return errors.New(fmt.Sprintf("O digito da conta precisa ser preenchido."))
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}

func citiValidateWallet(b interface{}) error {
	switch t := b.(type) {
	case *models.BoletoRequest:
		if t.Agreement.Wallet < 100 || t.Agreement.Wallet > 999 {
			return errors.New(fmt.Sprintf("A wallet deve conter somente 3 digítos."))
		}
		return nil
	default:
		return validations.InvalidType(t)
	}
}
