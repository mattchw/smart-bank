package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsValidCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	default:
		return false
	}
}
