package trimmedmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
	// Test case 1: Int data slice that can be trimmed
	intdata := make([]int32, 100)
	for i := 0; i < 100; i++ {
		intdata[i] = int32(i + 1)
	}

	floatdata := make([]float64, 100)
	for i := 0; i < 100; i++ {
		floatdata = append(floatdata, float64(i+1)+0.5)
	}

	emptydata := make([]interface{}, 0)
	cases := []struct {
		name       string
		data       interface{}
		proportion float64
		expected   float64
	}{
		{"int check", intdata, 0.05, 50.5},
		{"float check", floatdata, 0.05, 51},
		{"empty check", emptydata, 0.05, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := TrimmedMean(c.data, c.proportion)
			if err != nil {
				t.Errorf("TrimmedMean(%v) = %v; want %v", c.data, err, c.expected)
			}

		})
	}
}
