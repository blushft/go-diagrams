package main

import (
	"log"

	"github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/gcp"
)

func main() {
	d := diagram.New("app", attr.Label("App"), attr.Forward())

	dns := gcp.Network.Dns(attr.Label("DNS"))
	lb := gcp.Network.LoadBalancing(attr.Label("NLB"))
	cache := gcp.Database.Memorystore(attr.Label("Cache"))
	db := gcp.Database.Sql(attr.Label("Database"))

	dc := d.Group("GCP")
	dc.AddNodes(dns, lb, cache, db)
	dc.Group("services").
		Label("Service Layer").
		AddNodes(
			gcp.Compute.ComputeEngine(attr.Label("Server 1")),
			gcp.Compute.ComputeEngine(attr.Label("Server 2")),
			gcp.Compute.ComputeEngine(attr.Label("Server 3")),
		).
		ConnectAllFrom(lb, attr.Forward()).
		ConnectAllTo(cache, attr.Forward())

	dc.Group("data").Label("Data Layer").AddNodes(cache, db).Connect(cache, db)

	d.Connect(dns, lb, attr.Forward())

	if err := diagram.Render(d, diagram.OutputPath("./tmp")); err != nil {
		log.Fatal(err)
	}
}
