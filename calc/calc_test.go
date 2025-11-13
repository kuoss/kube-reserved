package calc_test

import (
	"testing"

	"github.com/kuoss/kube-reserved/calc"
	"github.com/stretchr/testify/assert"
)

func TestCalculateCPUReservation(t *testing.T) {
	testCases := []struct {
		name  string
		input int64
		want  int64
	}{
		{"one cpu core", 1000, 60},
		{"two cpu cores", 2000, 70},
		{"three cpu cores", 3000, 75},
		{"four cpu cores", 4000, 80},
		{"five cpu cores", 5000, 82},
		{"six cpu cores", 6000, 85},
		{"ten cpu cores", 10000, 95},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := calc.CalculateCPUReservation(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestCalculateMemoryReservation(t *testing.T) {
	const Mi = calc.Mi
	const Gi = calc.Gi

	testCases := []struct {
		name  string
		input int64
		want  int64
	}{
		{"500Mi", 500 * Mi, 255 * Mi},
		{"1Gi", 1 * Gi, 256 * Mi},
		{"2Gi", 2 * Gi, 512 * Mi},
		{"4Gi", 4 * Gi, 1 * Gi},
		{"8Gi", 8 * Gi, Gi + 4*Gi/5},
		{"16Gi", 16 * Gi, Gi + 4*Gi/5 + 8*Gi/10},
		{"32Gi", 32 * Gi, Gi + 4*Gi/5 + 8*Gi/10 + 16*Gi/100*6},
		{"64Gi", 64 * Gi, Gi + 4*Gi/5 + 8*Gi/10 + 48*Gi/100*6},
		{"128Gi", 128 * Gi, Gi + 4*Gi/5 + 8*Gi/10 + 112*Gi/100*6},
		{"256Gi", 256 * Gi, Gi + 4*Gi/5 + 8*Gi/10 + 112*Gi/100*6 + 128*Gi/100*2},
		{"18Gi", 18 * Gi, Gi + 4*Gi/5 + 8*Gi/10 + 2*Gi/100*6},
		{"42Gi", 42 * Gi, Gi + 4*Gi/5 + 8*Gi/10 + 26*Gi/100*6},
		{"1TiB", 1024 * Gi, Gi + 4*Gi/5 + 8*Gi/10 + 112*Gi/100*6 + 896*Gi/100*2},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := calc.CalculateMemoryReservation(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
