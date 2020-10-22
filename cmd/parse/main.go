package main

import (
	"github.com/alexflint/go-arg"
	"github.com/sirupsen/logrus"
	"os"
)

var args struct {
	Input string `arg:"-i" help:"Input file path (e.g: /tmp/1.yml)"`
	Output string `arg:"-o" help:"Output directory"`
}

func main(){
	arg.MustParse(&args)

	if args.Input == "" {
		logrus.Fatal("Input file not specified")
	}

	_, err := os.Open(args.Input)
	if err != nil {
		logrus.Fatalf("unable to open input file: %v", err)
	}

	_, err = os.Stat(args.Output)
	if err != nil {
		logrus.Fatal(err)
	}
}
