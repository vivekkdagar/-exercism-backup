package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    switch {
        case balance >= 5000 :
            return 2.475
        case balance >= 1000:
        	return 1.621
    	case balance < 0:
        	return 3.213
        default:
        	return 0.5
    }
}

func Interest(balance float64) float64 {
	// Get the interest rate as a float32, then convert it to float64
	// and divide by 100 to get the decimal value.
	rate := float64(InterestRate(balance)) / 100.0
	// Multiply the balance by the corrected rate.
	return balance * rate
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	currentBalance := balance

	for currentBalance < targetBalance {
		// AnnualBalanceUpdate calls InterestRate, which dynamically checks the balance
		// at the start of each year, applying the correct rate.
		currentBalance = AnnualBalanceUpdate(currentBalance)
		years++
	}

	return years
}
