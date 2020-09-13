package main

import (
	"fmt"
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/gcp"
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

	lb := gcp.Network.LoadBalancing(diagram.NodeLabel("nlb"))
	d.Add(lb)

	db := gcp.Database.Sql(diagram.NodeLabel("db"))
	d.Add(db)

	workers := make([]*diagram.Node, workerCount)

	for i := 0; i < workerCount; i++ {
		label := fmt.Sprintf("worker %d", i+1)
		workers[i] = gcp.Compute.ComputeEngine(diagram.NodeLabel(label))
	}

	d.Group(diagram.NewGroup("workers").
		Add(workers...).
		ConnectAllTo(db.ID()).
		ConnectAllFrom(lb.ID()),
	)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
