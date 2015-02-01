package finance

import (
	"fmt"
	"math"
	"strconv"
)

const (
	irrPercentageUpperBound    = 1.0
	irrPercentageStepPrecision = 0.001
)

// IRR Calculates the internal rate of return of a series of periodic incomes (positive or negative).
// The input is a map of float64 arrays keyed by period number. A given period can have multiple positive and/or
// negative incomes.
func IRR(periodicIncomes map[int][]float64) float64 {
	var lowestNpv float64
	var bestTestedIrr float64
	isTrendingDownwards := false
	for i := 0.0; i < irrPercentageUpperBound; i = i + irrPercentageStepPrecision {
		npv := netPresentValue(periodicIncomes, i)
		if lowestNpv == 0 {
			lowestNpv = npv
		}
		if math.Abs(npv) < math.Abs(lowestNpv) {
			isTrendingDownwards = true
			lowestNpv = npv
			bestTestedIrr = i
		} else {
			if isTrendingDownwards {
				break
			}
		}
	}
	return bestTestedIrr
}

// PresentValue calculates the present value for a given future value, interest rate, and number of periods.
func PresentValue(futureValue float64, interestRate float64, numPeriods int) float64 {
	n := float64(numPeriods)
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

func netPresentValue(periodicIncomes map[int][]float64, interestRate float64) float64 {
	var presentValues []float64
	for period, value := range periodicIncomes {
		for _, v := range value {
			pv := PresentValue(v, interestRate, period)
			presentValues = append(presentValues, pv)
		}
	}
	var netPresentValue float64
	for _, v := range presentValues {
		netPresentValue = netPresentValue + v
	}
	return netPresentValue
}
