package encription

import "golang.org/x/crypto/bcrypt"

type CustomEncryption interface {
	Encrypt(password []byte, cost int) ([]byte, error)
}

type DebtEncryption struct{}

func (e *DebtEncryption) Encrypt(password []byte, cost int) ([]byte, error) {
	result, encError := bcrypt.GenerateFromPassword(password, cost)
	if encError != nil {
		return nil, encError
	}
	return result, nil
}

func NewDebtEncryption() CustomEncryption {
	return &DebtEncryption{}
}
