package main

import (
	"log"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
)

func main() {
	d, err := diagram.New(diagram.Label("Web Service"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	inet := apps.Network.Internet().Label("Internet")
	proxy := apps.Network.Caddy().Label("Caddy")

	d.Connect(inet, proxy)

	ss := apps.Inmemory.Redis().Label("session")
	rs := apps.Inmemory.Redis().Label("replica")

	cache := diagram.NewGroup("cache").Label("Sessions").
		Connect(ss, rs, diagram.Bidirectional())

	dbmain := apps.Database.Postgresql().Label("Master DB")
	repls := []*diagram.Node{
		apps.Database.Postgresql().Label("DB Replica 1"),
		apps.Database.Postgresql().Label("DB Replica 2"),
	}

	db := diagram.NewGroup("database").Label("Database").
		Add(dbmain).Add(repls...)

	svcs := diagram.NewGroup("services").Label("Services").
		Add(
			apps.Container.Docker().Label("Replica 1"),
			apps.Container.Docker().Label("Replica 2"),
		).
		ConnectAllFrom(proxy.ID()).
		ConnectAllTo(dbmain.ID()).
		ConnectAllTo(ss.ID())

	for _, r := range repls {
		db.Connect(dbmain, r)
	}

	d.Group(svcs).
		Group(cache).
		Group(db)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}

}
