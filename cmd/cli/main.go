package main

import (
	"flag"
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/node"
	"github.com/blushft/go-diagrams/node/k8s"

	"github.com/pkg/errors"
)

var (
	out = "diagram.dot"
)

func init() {
	flag.StringVar(&out, "outfile", "diagram", "output file")
}

func main() {
	flag.Parse()

	d, err := diagram.New(diagram.Label("System"), diagram.Filename(out))
	if err != nil {
		log.Fatal(errors.Wrap(err, "new diagram"))
	}

	g := d.AddNode(k8s.Clusterconfig.Quota(node.Label("quota"))).
		Group("kubernetes", diagram.Label("kubernetes")).
		Connect(
			k8s.Clusterconfig.Hpa(node.Label("my hpa")),
			k8s.Clusterconfig.Limits(node.Label("some limits")),
		)

	if err := g.Render(); err != nil {
		log.Fatal(errors.Wrap(err, "render diagram"))
	}
}
