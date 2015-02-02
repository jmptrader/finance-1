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
	numChans := 0
	c := make(chan npvResult)
	for i := 0.0; i < irrPercentageUpperBound; i = i + irrPercentageStepPrecision {
		numChans++
		go calculateNPV(c, periodicIncomes, i)
	}

	for i := 0; i < numChans; i++ {
		result := <-c
		if lowestNpv == 0 {
			lowestNpv = result.npv
		}
		if math.Abs(result.npv) < math.Abs(lowestNpv) {
			lowestNpv = result.npv
			bestTestedIrr = result.interestRate
		}
	}
	return bestTestedIrr
}

type npvResult struct {
	npv          float64
	interestRate float64
}

func calculateNPV(c chan<- npvResult, periodicIncomes map[int][]float64, percentAttempt float64) {
	npv := NetPresentValue(periodicIncomes, percentAttempt)
	c <- npvResult{npv: npv, interestRate: percentAttempt}
	return
}

// PresentValue calculates the present value for a given future value, interest rate, and number of periods.
func PresentValue(futureValue float64, interestRate float64, numPeriods int) float64 {
	n := float64(numPeriods)
	pv := futureValue / math.Pow(1+interestRate, n)
	return pv
}

// CompoundInterest takes a principle amount, a nominal interest rate, a number of periods,
// and a number of compounding times per period and returns the compounded value.
func CompoundInterest(principleAmount float64, nominalInterestRate float64, numPeriods int, numTimesCompoundedPerPeriod int) float64 {
	exponent := float64(numTimesCompoundedPerPeriod * numPeriods)
	s := principleAmount * math.Pow(1+nominalInterestRate/float64(numTimesCompoundedPerPeriod), exponent)
	return s
}

// NetPresentValue takes periodic incomes (in the same format as IRR) and a nominal interest rate.
func NetPresentValue(periodicIncomes map[int][]float64, interestRate float64) float64 {
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
