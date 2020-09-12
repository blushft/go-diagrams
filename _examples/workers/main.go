package main

import (
	"fmt"
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/node"
	"github.com/blushft/go-diagrams/node/gcp"
)

func main() {
	workerCount := 5

	d, err := diagram.New(
		diagram.Label("Workers"),
		diagram.Filename("workers"),
		diagram.Direction("TB"),
	)

	if err != nil {
		log.Fatal(err)
	}

	lb := gcp.Network.LoadBalancing(node.Label("nlb"))
	db := gcp.Database.Sql(node.Label("db"))

	workers := make([]*node.Node, workerCount)

	for i := 0; i < workerCount; i++ {
		label := fmt.Sprintf("worker %d", i+1)
		workers[i] = gcp.Compute.ComputeEngine(node.Label(label))
	}

	for _, w := range workers {
		d.Connect(lb, w)
		d.Connect(w, db)
	}

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
