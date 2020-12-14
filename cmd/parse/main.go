package main

import (
	"github.com/alexflint/go-arg"
	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/diagram"
	"github.com/sirupsen/logrus"
	"os"
)

var args struct {
	Input string `arg:"-i" help:"Input file path (e.g: /tmp/1.yml)"`
	Output string `arg:"-o" help:"Output directory"`
	AssetsDir string `arg:"-a" help:"Directory where the assets are stored"`
}

func main(){
	arg.MustParse(&args)

	if args.Input == "" {
		logrus.Fatal("Input file not specified")
	}

	_, err := os.Stat(args.Output)
	if err != nil {
		logrus.Fatalf("output not valid: %v", err)
	}

	file, err := os.Open(args.Input)
	if err != nil {
		logrus.Fatalf("unable to open input file: %v", err)
	}

	d := diagram.New("diag", attr.Label("my-diagram"), attr.OutputOrder(attr.NodesFirst))

	err = d.Parse(file)
	if err != nil {
		logrus.Fatalf("unable to parse input: %v", err)
	}

	err = diagram.Render(d)
	if err != nil {
		logrus.Fatalf("unable to render input: %v", err)
	}
}
