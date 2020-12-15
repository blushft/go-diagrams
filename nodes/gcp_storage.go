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
	opts = append(opts, attr.AssetImage("assets/gcp/storage/filestore.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("filestore", opts...)
}

func (c *gcpStorageContainer) PersistentDisk(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/storage/persistent-disk.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("persistent-disk", opts...)
}

func (c *gcpStorageContainer) Storage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/gcp/storage/storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage", opts...)
}


func init() {
	const namespace = "gcp.storage"
	Register(namespace, "Filestore", GCPStorage.Filestore)
	Register(namespace, "PersistentDisk", GCPStorage.PersistentDisk)
	Register(namespace, "Storage", GCPStorage.Storage)
}