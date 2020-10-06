package firebase

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
	return node.New("functions", attr.AssetImage("assets/firebase/develop/functions.png"), attr.Shape(attr.None))
}

func (c *developContainer) Hosting(opts ...attr.Attribute) *node.Node {
	return node.New("hosting", attr.AssetImage("assets/firebase/develop/hosting.png"), attr.Shape(attr.None))
}

func (c *developContainer) MlKit(opts ...attr.Attribute) *node.Node {
	return node.New("ml-kit", attr.AssetImage("assets/firebase/develop/ml-kit.png"), attr.Shape(attr.None))
}

func (c *developContainer) RealtimeDatabase(opts ...attr.Attribute) *node.Node {
	return node.New("realtime-database", attr.AssetImage("assets/firebase/develop/realtime-database.png"), attr.Shape(attr.None))
}

func (c *developContainer) Storage(opts ...attr.Attribute) *node.Node {
	return node.New("storage", attr.AssetImage("assets/firebase/develop/storage.png"), attr.Shape(attr.None))
}

func (c *developContainer) Authentication(opts ...attr.Attribute) *node.Node {
	return node.New("authentication", attr.AssetImage("assets/firebase/develop/authentication.png"), attr.Shape(attr.None))
}

func (c *developContainer) Firestore(opts ...attr.Attribute) *node.Node {
	return node.New("firestore", attr.AssetImage("assets/firebase/develop/firestore.png"), attr.Shape(attr.None))
}
