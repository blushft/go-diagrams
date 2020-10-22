package diagram_test

import (
	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/diagram"
	"os"
	"testing"
)

func TestDiagramParse(t *testing.T){
	file, err := os.Open("../test/assets/parse/ymls/1.yml")
	if err != nil {
		t.Fatal(err)
	}
	d := diagram.New("diag", attr.Label("my-diagram"))

	err = d.Parse(file)
	if err != nil {
		t.Fatal(err)
	}

	err = diagram.Render(d)
	if err != nil {
		t.Fatal(err)
	}
}
