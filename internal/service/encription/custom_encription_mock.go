package encription

import "github.com/stretchr/testify/mock"

type CustomEncryptionMock struct {
	mock.Mock
}

func (e *CustomEncryptionMock) Encrypt(password []byte, cost int) ([]byte, error) {
	args := e.Called(password, cost)
	bytes := args.Get(0)
	value, _ := bytes.(string)
	err := args.Error(1)
	if bytes == nil {
		return nil, err
	}
	return []byte(value), err
}
