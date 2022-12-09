/*
package imlements mathematition statistics functions
*/
package statistics

import (
	"math"
	"sort"
)
/*
	calculates  the arithmetic average. The result roundes to integer
*/
func Average(nmbs []int) int {
	if len(nmbs) == 0 {
		return 0
	}
	s := 0
	for _, n := range nmbs {
		s += n
	}
	return int(math.Round(float64(s) / float64(len(nmbs))))
}

/*
	calculates the median of a sets of integer numbers. The result roundes to integer
*/
func Median(nmbs []int) int {
	sort.Ints(nmbs)
	l := len(nmbs)
	if l%2 == 0 {
		return int(math.Round(float64(nmbs[l/2-1]+nmbs[l/2]) / 2.0))
	} else {
		return int(math.Round(float64(nmbs[l/2]) / 2.0))
	}
}

/*
	calculates the variance of a sets of integer numbers. The result roundes to integer
*/
func Variance(nmbs []int) int {
	if len(nmbs) == 0 {
		return 0
	}
	s := 0
	av := Average(nmbs)
	for _, n := range nmbs {
		s += (n - av) * (n - av)
	}
	return int(math.Round(float64(s) / float64(len(nmbs))))
}

/*
	calculates the standard deviation of a sets of integer numbers. The result roundes to integer
*/
func StandardDeviation(nmbs []int) int {
	return int(math.Round(math.Sqrt(float64(Variance(nmbs)))))
}
