package example

import (
	"testing"

	"github.com/budiuno/gocsvparser"
)

type mtcarsFlat struct {
	Model            string  `csv:"model"`
	MilesPerGalon    float64 `csv:"mpg"`
	Cylinder         int     `csv:"cyl"`
	Displacement     float64 `csv:"disp"`
	Horsepower       int     `csv:"hp"`
	DriveShaftRatio  float64 `csv:"drat"`
	Weight           float64 `csv:"wt"`
	QuerterMileTime  float64 `csv:"qsec"`
	VEngine          bool    `csv:"vs"`
	AutoTransmission bool    `csv:"am"`
	Gear             int     `csv:"gear"`
	Carburetors      int     `csv:"carb"`
}

func TestMtcars(t *testing.T) {
	var (
		cars []mtcarsFlat
		err  error
	)

	err = gocsvparser.Unmarshal(MtcarsCsv, &cars)
	if err != nil {
		t.Fatalf("unexpected error %+v", err)
	}

	for _, c := range cars {
		t.Logf("%+v", c)
	}
}
