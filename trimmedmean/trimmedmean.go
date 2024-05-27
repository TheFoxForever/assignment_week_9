package trimmedmean

import (
	"errors"
	"sort"
)

//takes a slice of float64 and a proportion to be trimmed from both ends
func TrimmedMean(data []float64, proportion float64) (float64, error) {
    // checking incorrect parameters and that data is not empty
    if proportion < 0 || proportion >= 1.0 {
        return 0, errors.New("proportion must be between 0 and 1.0")
    }

    dataLength := len(data)
    if dataLength < 100 {
        return 0, errors.New("must provide at least 100 data points")
    }

    // sorting data
    sort.Float64s(data)

    // determining the number of elements to trim from each end
    trimAmount := int(proportion * float64(dataLength))

    // calculating the trimmed mean
    sum := 0.0
    for i := trimAmount; i < dataLength-trimAmount; i++ {
        sum += data[i]
    }

    // calculating the length of the trimmed data
    trimLength := dataLength - 2*trimAmount

    //returning the sum divided by the length of the trimmed data, or nil if there is no error
    return sum / float64(trimLength), nil
}
