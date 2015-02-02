package finance

import (
	"math"
	"testing"
)

const ExpectedPrecision = 0.01

func TestIRR(t *testing.T) {
	payments1 := map[int][]float64{
		0: []float64{-2000.00},
		1: []float64{100.00},
		2: []float64{100.00},
		3: []float64{100.00, 2500.00},
	}
	payments2 := map[int][]float64{
		0: []float64{-2390.00},
		1: []float64{1000.00},
		2: []float64{2910.00},
		3: []float64{9.00},
	}
	payments3 := map[int][]float64{
		0: []float64{-300000.00},
		1: []float64{150000.00},
		2: []float64{150000.00},
		3: []float64{150000.00},
		4: []float64{10000.00},
	}

	answer1 := 0.124
	answer2 := 0.334
	answer3 := 0.243

	result1 := IRR(payments1)
	result2 := IRR(payments2)
	result3 := IRR(payments3)

	compareFloatToPrecision(t, result1, answer1)
	compareFloatToPrecision(t, result2, answer2)
	compareFloatToPrecision(t, result3, answer3)
}

func TestPresentValue(t *testing.T) {
	result1 := PresentValue(40000.00, 0.023, 4)
	correctAnswer1 := 36522.24
	compareFloatToPrecision(t, result1, correctAnswer1)

	result2 := PresentValue(489000.13, 0.053, 43)
	correctAnswer2 := 53074.88
	compareFloatToPrecision(t, result2, correctAnswer2)
}

func TestCompoundInterest(t *testing.T) {
	result := CompoundInterest(1500.00, 0.043, 6, 4)
	correctAnswer := 1938.84
	compareFloatToPrecision(t, result, correctAnswer)
}

func compareFloatToPrecision(t *testing.T, result float64, correctAnswer float64) {
	absRemainder := math.Abs(math.Remainder(correctAnswer, result))
	if absRemainder >= ExpectedPrecision {
		t.Errorf("Expected %v, got %v", correctAnswer, result)
	}
}
