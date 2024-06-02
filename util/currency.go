package util

// Constants for all supported currencies/
const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
	CAD = "CAD"
	ZMW = "ZMW"
)

// IsSupportedCurrency returns true if the currency si supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP, CAD, ZMW:
		return true
	}
	return false
}
