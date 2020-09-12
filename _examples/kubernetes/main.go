package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/node"
	"github.com/blushft/go-diagrams/node/k8s"
)

func main() {
	d, err := diagram.New(
		diagram.Label("Kubernetes"),
		diagram.Filename("k8s"),
		diagram.Direction("TB"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ingress := k8s.Network.Ing(node.Label("nginx"))
	svc := k8s.Network.Svc(node.Label("http"))

	d.Connect(ingress, svc)

	d.Group("pods", diagram.Label("app")).Connect(svc, k8s.Compute.Pod(node.Label("web server")))

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}
}
