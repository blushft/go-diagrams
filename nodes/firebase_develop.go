package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type developContainer struct {
	path  string
	attrs []attr.Attribute
}

var Develop = &developContainer{path: "assets/firebase/develop"}

func (c *developContainer) Functions(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/develop/functions.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("functions", opts...)
}

func (c *developContainer) Hosting(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/develop/hosting.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("hosting", opts...)
}

func (c *developContainer) MlKit(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/develop/ml-kit.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("ml-kit", opts...)
}

func (c *developContainer) RealtimeDatabase(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/develop/realtime-database.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("realtime-database", opts...)
}

func (c *developContainer) Storage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/develop/storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage", opts...)
}

func (c *developContainer) Authentication(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/develop/authentication.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("authentication", opts...)
}

func (c *developContainer) Firestore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/firebase/develop/firestore.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("firestore", opts...)
}
