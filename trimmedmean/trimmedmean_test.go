package trimmedmean

import (
	"testing"
)

func TestTrimmedMean(t *testing.T) {
    // Test case 1: Int data slice that can be trimmed
    data := make([]float64, 100)
    proportion := 0.05
    for i := 0; i < 100; i++ {
        data[i] = float64(i + 1)
    }
    expected := 50.5
    result, err := TrimmedMean(data, proportion)
    if err != nil {
        t.Errorf("Error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }

    // Test case 2: Float data slice that can be trimmed
    data = make([]float64, 100)
	for i := 0; i < 100; i++ {
		data = append(data, float64(i+1)+0.5)
	}
    expected = 23
    _, err = TrimmedMean(data, proportion)
    if err != nil {
        t.Errorf("Error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }

    // Test case 2: Empty data slice, should return error
    data = []float64{}
    _, err = TrimmedMean(data, proportion)
    if err == nil {
        t.Error("Expected error for empty data slice")
    }
}

