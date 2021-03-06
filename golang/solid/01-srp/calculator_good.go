package srp

import (
	"fmt"
	"io"
)

// Calculator calculates the test coverage for a directory and it's sub-directories
type Calculator struct {
	// coverage data populated by `Calculate()` method
	data map[string]float64
}

// Calculate will calculate the coverage
func (c *Calculator) Calculate(_ string) error {
	// run `go test -cover ./[path]/...` and store the results
	return nil
}

func (c *Calculator) getData() map[string]float64 {
	// copy and return the map
	return nil
}

type Printer interface {
	Output(data map[string]float64)
}

type DefaultPrinter struct {
	Writer io.Writer
}

// Output implements Printer
func (d *DefaultPrinter) Output(data map[string]float64) {
	for path, result := range data {
		fmt.Fprintf(d.Writer, "%s -> %.1f\n", path, result)
	}
}

type CSVPrinter struct {
	Writer io.Writer
}

// Output implements Printer
func (d *CSVPrinter) Output(data map[string]float64) {
	for path, result := range data {
		fmt.Fprintf(d.Writer, "%s,%.1f\n", path, result)
	}
}
