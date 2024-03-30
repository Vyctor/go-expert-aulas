package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectec := 5.0

	result := CalculateTax(amount)

	if result != expectec {
		t.Errorf("Expected %.2f, got %.2f", expectec, result)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{0.0, 0.0},
		{500.0, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, test := range table {
		result := CalculateTax(test.amount)
		if result != test.expect {
			t.Errorf("Expected %.2f, got %.2f", test.expect, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}

	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 5 {
			t.Errorf("Passed %f and eceived %f but expected 5", amount, result)
		}
		if amount >= 1000 && result != 10 {
			t.Errorf("Passed %f and eceived %f but expected 10", amount, result)
		}
		if amount >= 0 && amount < 1000 && result != 5 {
			t.Errorf("Passed %f and eceived %f but expected 5", amount, result)
		}
	})
}
