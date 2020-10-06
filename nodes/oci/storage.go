package oci

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Storage = &storageContainer{path: "assets/oci/storage"}

func (c *storageContainer) FileStorageWhite(opts ...attr.Attribute) *node.Node {
	return node.New("file-storage-white", attr.AssetImage("assets/oci/storage/file-storage-white.png"), attr.Shape(attr.None))
}

func (c *storageContainer) FileStorage(opts ...attr.Attribute) *node.Node {
	return node.New("file-storage", attr.AssetImage("assets/oci/storage/file-storage.png"), attr.Shape(attr.None))
}

func (c *storageContainer) ObjectStorage(opts ...attr.Attribute) *node.Node {
	return node.New("object-storage", attr.AssetImage("assets/oci/storage/object-storage.png"), attr.Shape(attr.None))
}

func (c *storageContainer) StorageGateway(opts ...attr.Attribute) *node.Node {
	return node.New("storage-gateway", attr.AssetImage("assets/oci/storage/storage-gateway.png"), attr.Shape(attr.None))
}

func (c *storageContainer) BlockStorageCloneWhite(opts ...attr.Attribute) *node.Node {
	return node.New("block-storage-clone-white", attr.AssetImage("assets/oci/storage/block-storage-clone-white.png"), attr.Shape(attr.None))
}

func (c *storageContainer) BlockStorageClone(opts ...attr.Attribute) *node.Node {
	return node.New("block-storage-clone", attr.AssetImage("assets/oci/storage/block-storage-clone.png"), attr.Shape(attr.None))
}

func (c *storageContainer) BlockStorage(opts ...attr.Attribute) *node.Node {
	return node.New("block-storage", attr.AssetImage("assets/oci/storage/block-storage.png"), attr.Shape(attr.None))
}

func (c *storageContainer) ElasticPerformanceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-performance-white", attr.AssetImage("assets/oci/storage/elastic-performance-white.png"), attr.Shape(attr.None))
}

func (c *storageContainer) ElasticPerformance(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-performance", attr.AssetImage("assets/oci/storage/elastic-performance.png"), attr.Shape(attr.None))
}

func (c *storageContainer) BackupRestore(opts ...attr.Attribute) *node.Node {
	return node.New("backup-restore", attr.AssetImage("assets/oci/storage/backup-restore.png"), attr.Shape(attr.None))
}

func (c *storageContainer) BlockStorageWhite(opts ...attr.Attribute) *node.Node {
	return node.New("block-storage-white", attr.AssetImage("assets/oci/storage/block-storage-white.png"), attr.Shape(attr.None))
}

func (c *storageContainer) BucketsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("buckets-white", attr.AssetImage("assets/oci/storage/buckets-white.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Buckets(opts ...attr.Attribute) *node.Node {
	return node.New("buckets", attr.AssetImage("assets/oci/storage/buckets.png"), attr.Shape(attr.None))
}

func (c *storageContainer) StorageGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("storage-gateway-white", attr.AssetImage("assets/oci/storage/storage-gateway-white.png"), attr.Shape(attr.None))
}

func (c *storageContainer) BackupRestoreWhite(opts ...attr.Attribute) *node.Node {
	return node.New("backup-restore-white", attr.AssetImage("assets/oci/storage/backup-restore-white.png"), attr.Shape(attr.None))
}

func (c *storageContainer) DataTransferWhite(opts ...attr.Attribute) *node.Node {
	return node.New("data-transfer-white", attr.AssetImage("assets/oci/storage/data-transfer-white.png"), attr.Shape(attr.None))
}

func (c *storageContainer) DataTransfer(opts ...attr.Attribute) *node.Node {
	return node.New("data-transfer", attr.AssetImage("assets/oci/storage/data-transfer.png"), attr.Shape(attr.None))
}

func (c *storageContainer) ObjectStorageWhite(opts ...attr.Attribute) *node.Node {
	return node.New("object-storage-white", attr.AssetImage("assets/oci/storage/object-storage-white.png"), attr.Shape(attr.None))
}
