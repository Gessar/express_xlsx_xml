package declaration

import "strconv"

func String(s string) *string {
	return &s
}

// FormatCurrencyValue форматирует значение валюты с двумя десятичными знаками
func FormatCurrencyValue(value float64) *string {
	formatted := strconv.FormatFloat(value, 'f', 2, 64)
	return &formatted
}

func FormatGrossMassValue(value float64) *string {
	formatted := strconv.FormatFloat(value, 'f', 3, 64)
	return &formatted
}
