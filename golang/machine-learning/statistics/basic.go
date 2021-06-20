package statistics

import (
	"fmt"
	"io"
	"log"

	"github.com/gonum/floats"
	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"
)

func Basics(b io.Reader) {
	// Create a dataframe from the CSV source.
	irisDF := dataframe.ReadCSV(b)

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	sepalLength := irisDF.Col("petal_length").Float()

	// Calculate the Mean of the variable.
	meanVal := stat.Mean(sepalLength, nil)

	// Calculate the Mode of the variable.
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	// Calculate the Median of the variable.
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Median value: %0.2f\n\n", medianVal)
}

func Spread(b io.Reader) {
	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(b)

	// Get the float values from the "sepal_length" column as
	// we will be looking at the measures for this variable.
	sepalLength := irisDF.Col("petal_length").Float()

	// Calculate the Max of the variable.
	minVal := floats.Min(sepalLength)

	// Calculate the Max of the variable.
	maxVal := floats.Max(sepalLength)

	// Calculate the Median of the variable.
	rangeVal := maxVal - minVal

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(sepalLength, nil)

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(sepalLength, nil)

	// Sort the values.
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Get the Spread.
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n", rangeVal)
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n\n", quant75)
}