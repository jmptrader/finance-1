package finance

import (
  "testing"
  "fmt"
)

func TestIRR(t *testing.T) {
  PAYMENTS_1 := map[int][]float64{
    0: []float64{-2000.00},
    1: []float64{100.00},
    2: []float64{100.00},
    3: []float64{100.00, 2500.00},
  }
  PAYMENTS_2 := map[int][]float64{
    0: []float64{-2390.00},
    1: []float64{1000.00},
    2: []float64{2910.00},
    3: []float64{9.00},
  }
  PAYMENTS_3 := map[int][]float64{
    0: []float64{-300000.00},
    1: []float64{150000.00},
    2: []float64{150000.00},
    3: []float64{150000.00},
    4: []float64{10000.00},
  }

  ANSWER_1 := "0.124"
  ANSWER_2 := "0.334"
  ANSWER_3 := "0.243"

  runTest(t, PAYMENTS_1, ANSWER_1)
  runTest(t, PAYMENTS_2, ANSWER_2)
  runTest(t, PAYMENTS_3, ANSWER_3)
}

func runTest(t *testing.T, payments map[int][]float64, correctAnswer string) {
  result := IRR(payments)
  if fmt.Sprintf("%.3f", result) != correctAnswer {
    t.Errorf("Expected %v, got %v", correctAnswer, result)
  }
}
