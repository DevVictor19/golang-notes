package math

import (
	"fmt"
	"strconv"
)

func Average(n ...float64) float64 {
	total := 0.0
	for _, num := range n {
		total += num
	}
	average := total / float64(len(n))
	rounded, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", average), 64)
	return rounded
}
