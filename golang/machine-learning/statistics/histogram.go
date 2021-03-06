package statistics

import (
	"fmt"
	"io"
	"log"

	"github.com/kniren/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func VisualizeHistogram(b io.Reader) {
	// Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(b)

	// Create a histogram for each of the feature columns in the dataset.
	for _, colName := range irisDF.Names() {

		// If the column is one of the feature columns,
		// let's create a histogram of the values.
		if colName != "species" {
			// Create a plotter.Values value and fill it with the
			// values from the respective column of the dataframe.
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// Make a plot and set its title.
			p := plot.New()
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

			// Create a histogram of our values drawn
			// from the standard normal.
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			// Normalize the histogram.
			h.Normalize(1)

			// Add the histogram to the plot.
			p.Add(h)
			// Save the plot to a PNG file.
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}
}
