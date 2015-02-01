package finance

import (
	"fmt"
	"testing"
)

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

	answer1 := "0.124"
	answer2 := "0.334"
	answer3 := "0.243"

	runIrrTest(t, payments1, answer1)
	runIrrTest(t, payments2, answer2)
	runIrrTest(t, payments3, answer3)
}

func runIrrTest(t *testing.T, payments map[int][]float64, correctAnswer string) {
	result := IRR(payments)
	if fmt.Sprintf("%.3f", result) != correctAnswer {
		t.Errorf("Expected %v, got %v", correctAnswer, result)
	}
}

func TestPresentValue(t *testing.T) {
	result1 := PresentValue(40000.00, 0.023, 4)
	correctAnswer1 := "36522.24"
	runPVTest(t, result1, correctAnswer1)

	result2 := PresentValue(489000.13, 0.053, 43)
	correctAnswer2 := "53074.88"
	runPVTest(t, result2, correctAnswer2)
}

func runPVTest(t *testing.T, result float64, correctAnswer string) {
	if fmt.Sprintf("%.2f", result) != correctAnswer {
		t.Errorf("Expected %v, got %v", correctAnswer, result)
	}
}
