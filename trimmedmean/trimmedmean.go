package trimmedmean

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// converts data to a slice of float64
func convertToFloat64Slice(data interface{}) ([]float64, error) {
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		return nil, errors.New("data must be a slice")
	}

	converted := make([]float64, val.Len())
	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)
		switch elem.Kind() {
		case reflect.Float64:
			converted[i] = elem.Float()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			converted[i] = float64(elem.Int())
		case reflect.Float32:
			converted[i] = float64(elem.Float())
		default:
			return nil, fmt.Errorf("unsupported element type: %s", elem.Kind())
		}
	}
	return converted, nil
}

// takes a slice of float64 and a proportion to be trimmed from both ends
func TrimmedMean(data interface{}, proportion float64) (float64, error) {
	// convert to float64 slice
	dataslice, err := convertToFloat64Slice(data)
	if err != nil {
		return 0, err
	}

	if proportion < 0 || proportion >= 1.0 {
		return 0, errors.New("proportion must be between 0 and 1.0")
	}

	dataLength := len(dataslice)
	if dataLength < 0 {
		return 0, errors.New("cannot provide empty slice of data")
	}

	sort.Float64s(dataslice)

	// determining the number of elements to trim from each end
	trimAmount := int(float64(len(dataslice)) * proportion)
	trimmedSlice := dataslice[trimAmount : len(dataslice)-trimAmount]

	// calculating the trimmed mean
	sum := 0.0
	for _, item := range trimmedSlice {
		sum += item

	}

	//returning the sum divided by the length of the trimmed data, or nil if there is no error
	return sum / float64(len(trimmedSlice)), nil
}
