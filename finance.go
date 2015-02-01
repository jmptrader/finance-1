package finance

import (
	"fmt"
	"math"
	"strconv"
)

// IRR (Internal Rate of Return) is the interest rate that makes the Net Present Value zero
func IRR(yearlyPayments map[int][]float64) float64 {
	var lowestNpv float64
	var bestTestedIrr float64
	for i := 0.0; i < 1.001; i = i + 0.001 {
		npv := netPresentValue(yearlyPayments, i)
		if lowestNpv == 0 {
			lowestNpv = npv
		}
		if math.Abs(npv) < math.Abs(lowestNpv) {
			lowestNpv = npv
			bestTestedIrr = i
		}
	}
	return bestTestedIrr
}

func presentValue(futureValue float64, interestRate float64, numYears int) float64 {
	n := float64(numYears)
	pv := futureValue / math.Pow(1+interestRate, n)
	return float64(roundTo2DecimalPlaces(pv))
}

func roundTo2DecimalPlaces(value float64) float64 {
	roundedValue, err := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 32)
	if err != nil {
		panic(err)
	}
	return roundedValue
}

func netPresentValue(yearlyPayments map[int][]float64, interestRate float64) float64 {
	var presentValues []float64
	for year, value := range yearlyPayments {
		for _, v := range value {
			pv := presentValue(v, interestRate, year)
			presentValues = append(presentValues, pv)
		}
	}
	var netPresentValue float64
	for _, v := range presentValues {
		netPresentValue = netPresentValue + v
	}
	return netPresentValue
}
