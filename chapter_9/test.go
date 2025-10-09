package main

import "testing"

func Average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}


func TestAverage(t *testing.T) {
	got := Average([]float64{98, 93, 77, 82, 83})
	want := 86.6

	if got != want {
		t.Errorf("got %.1f, want %.1f", got, want)
	}
}

