package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Storage = &storageContainer{path: "assets/gcp/storage"}

func (c *storageContainer) Filestore(opts ...attr.Attribute) *node.Node {
	return node.New("filestore", attr.AssetImage("assets/gcp/storage/filestore.png"), attr.Shape(attr.None))
}

func (c *storageContainer) PersistentDisk(opts ...attr.Attribute) *node.Node {
	return node.New("persistent-disk", attr.AssetImage("assets/gcp/storage/persistent-disk.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Storage(opts ...attr.Attribute) *node.Node {
	return node.New("storage", attr.AssetImage("assets/gcp/storage/storage.png"), attr.Shape(attr.None))
}


func init() {
	const namespace = "gcp.storage"
	Register(namespace, "Filestore", Storage.Filestore)
	Register(namespace, "PersistentDisk", Storage.PersistentDisk)
	Register(namespace, "Storage", Storage.Storage)
}