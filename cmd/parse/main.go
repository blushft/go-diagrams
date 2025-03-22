package main

import (
	"github.com/alexflint/go-arg"
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/gcp"
	"github.com/blushft/go-diagrams/nodes/generic"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func diag1(){
	d, err := diagram.New(diagram.Filename("app"), diagram.Label("App"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	dns := gcp.Network.Dns(diagram.NodeLabel("DNS"))
	lb := gcp.Network.LoadBalancing(diagram.NodeLabel("NLB"))
	cache := gcp.Database.Memorystore(diagram.NodeLabel("Cache"))
	db := gcp.Database.Sql(diagram.NodeLabel("Database"))

	dc := diagram.NewGroup("GCP")
	dc.NewGroup("services").
		Label("Service Layer").
		Add(
			gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 1")),
			gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 2")),
			gcp.Compute.ComputeEngine(diagram.NodeLabel("Server 3")),
		).
		ConnectAllFrom(lb.ID(), diagram.Forward()).
		ConnectAllTo(cache.ID(), diagram.Forward())

	dc.NewGroup("data").Label("Data Layer").Add(cache, db).Connect(cache, db)

	d.Connect(dns, lb, diagram.Forward()).Group(dc)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}

func diag2(){
	d, err := diagram.New(diagram.Label("my-diagram"), diagram.Filename("diagram"))
	if err != nil {
		log.Fatal(err)
	}

	fw := generic.Network.Firewall().Label("fw")
	sw := generic.Network.Switch().Label("sw")

	d.Connect(fw, sw)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}

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
