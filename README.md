# Go-Diagrams

## Fast and easy application diagrams

Go-Diagrams is a loose port of [diagrams](https://github.com/mingrammer/diagrams).

## Contents

- [Features](#features)
- [Usage](#usage)

## Features

Turn this:

```golang
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
```

Into this:

![app-diagram](images/app-diagram.png)

## Usage

```sh
go get github.com/blushft/go-diagrams
```

Create a diagram:

```golang
d, err := diagram.New(diagram.Label("my-diagram"), diagram.Filename("my-diagram"))
if err != nil {
    log.Fatal(err)
}

fw := generic.Network.Firewall().Label("fw")
sw := generic.Network.Switch().Label("sw")

d.Connect(fw, sw)
```

Render the output:

```golang
if err := d.Render(); err != nil {
    log.Fatal(err)
}
```

Go-Diagrams will create a folder in the current working directory with the graphviz DOT file and any image assets.

```sh
go build -o diagram-example .
./diagram-example
```

Create an ouput image with any graphviz compatible renderer:

```sh
dot -Tpng go-diagrams/my-diagram.dot > diagram-example.png
```
