package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type gcpStorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var GCPStorage = &gcpStorageContainer{path: "assets/gcp/storage"}

func (c *gcpStorageContainer) Filestore(opts ...attr.Attribute) *node.Node {
	return node.New("filestore", attr.AssetImage("assets/gcp/storage/filestore.png"), attr.Shape(attr.None))
}

func (c *gcpStorageContainer) PersistentDisk(opts ...attr.Attribute) *node.Node {
	return node.New("persistent-disk", attr.AssetImage("assets/gcp/storage/persistent-disk.png"), attr.Shape(attr.None))
}

func (c *gcpStorageContainer) Storage(opts ...attr.Attribute) *node.Node {
	return node.New("storage", attr.AssetImage("assets/gcp/storage/storage.png"), attr.Shape(attr.None))
}


func init() {
	const namespace = "gcp.storage"
	Register(namespace, "Filestore", GCPStorage.Filestore)
	Register(namespace, "PersistentDisk", GCPStorage.PersistentDisk)
	Register(namespace, "Storage", GCPStorage.Storage)
}