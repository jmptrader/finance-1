package finance

import "math"

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
	hasTrendedDownwards := false
	for i := 0.0; i < irrPercentageUpperBound; i = i + irrPercentageStepPrecision {
		npv := netPresentValue(periodicIncomes, i)
		if lowestNpv == 0 {
			lowestNpv = npv
		}
		if math.Abs(npv) < math.Abs(lowestNpv) {
			hasTrendedDownwards = true
			lowestNpv = npv
			bestTestedIrr = i
		} else {
			if hasTrendedDownwards {
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
	return pv
}

// CompoundInterest takes a principle amount, a nominal interest rate, and a number of periods
// and returns the compounded value
func CompoundInterest(principleAmount float64, nominalInterestRate float64, numPeriods int, numTimesCompoundedPerPeriod int) float64 {
	exponent := float64(numTimesCompoundedPerPeriod * numPeriods)
	s := principleAmount * math.Pow(1+nominalInterestRate/float64(numTimesCompoundedPerPeriod), exponent)
	return s
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
