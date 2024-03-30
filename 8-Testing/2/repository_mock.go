package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(amount float64) error {
	args := m.Called(amount)
	return args.Error(0)
}
