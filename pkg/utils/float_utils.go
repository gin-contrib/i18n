package utils

import (
	"math"
	"strconv"
)

// Round --
func Round(val float64, roundOn float64, places int) float64 {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	return round / pow
}

// Float64ToString --
func Float64ToString(num float64, precision int) string {
	return strconv.FormatFloat(num, 'f', precision, 64)
}

// Float64InSlice --
func Float64InSlice(value float64, list []float64) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
