package tax

import "errors"

type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0.0 {
		return 0.0, errors.New("amount should be greater than 0")
	}

	if amount >= 1000.0 && amount < 20000 {
		return 10.0, nil
	}

	if amount >= 20000 {
		return 20.0, nil
	}

	return 5.0, nil
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0.0 {
		return 0
	}

	if amount >= 1000.0 && amount < 20000 {
		return 10
	}

	if amount >= 20000 {
		return 20
	}

	return 5
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTax2(amount)
	return repository.SaveTax(tax)
}
