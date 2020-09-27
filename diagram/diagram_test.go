package diagram_test

import (
	"github.com/blushft/go-diagrams/diagram"
	"os"
	"testing"
)

func TestDiagramParse(t *testing.T){
	file, err := os.Open("../test/assets/parse/ymls/1.yml")
	if err != nil {
		t.Fatal(err)
	}
	d, err := diagram.New(diagram.Label("my-diagram"),
		diagram.DirName("/tmp/diagrams-1"),
		diagram.Filename("1"))
	if err != nil {
		t.Fatal(err)
	}

	err = d.Parse(file)
	if err != nil {
		t.Fatal(err)
	}

	err = d.Render()
	if err != nil {
		t.Fatal(err)
	}
}
