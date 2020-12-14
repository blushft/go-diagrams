package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type ociStorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var OciStorage = &ociStorageContainer{path: "assets/oci/storage"}

func (c *ociStorageContainer) FileStorageWhite(opts ...attr.Attribute) *node.Node {
	return node.New("file-storage-white", attr.AssetImage("assets/oci/storage/file-storage-white.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) FileStorage(opts ...attr.Attribute) *node.Node {
	return node.New("file-storage", attr.AssetImage("assets/oci/storage/file-storage.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) ObjectStorage(opts ...attr.Attribute) *node.Node {
	return node.New("object-storage", attr.AssetImage("assets/oci/storage/object-storage.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) StorageGateway(opts ...attr.Attribute) *node.Node {
	return node.New("storage-gateway", attr.AssetImage("assets/oci/storage/storage-gateway.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) BlockStorageCloneWhite(opts ...attr.Attribute) *node.Node {
	return node.New("block-storage-clone-white", attr.AssetImage("assets/oci/storage/block-storage-clone-white.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) BlockStorageClone(opts ...attr.Attribute) *node.Node {
	return node.New("block-storage-clone", attr.AssetImage("assets/oci/storage/block-storage-clone.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) BlockStorage(opts ...attr.Attribute) *node.Node {
	return node.New("block-storage", attr.AssetImage("assets/oci/storage/block-storage.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) ElasticPerformanceWhite(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-performance-white", attr.AssetImage("assets/oci/storage/elastic-performance-white.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) ElasticPerformance(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-performance", attr.AssetImage("assets/oci/storage/elastic-performance.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) BackupRestore(opts ...attr.Attribute) *node.Node {
	return node.New("backup-restore", attr.AssetImage("assets/oci/storage/backup-restore.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) BlockStorageWhite(opts ...attr.Attribute) *node.Node {
	return node.New("block-storage-white", attr.AssetImage("assets/oci/storage/block-storage-white.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) BucketsWhite(opts ...attr.Attribute) *node.Node {
	return node.New("buckets-white", attr.AssetImage("assets/oci/storage/buckets-white.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) Buckets(opts ...attr.Attribute) *node.Node {
	return node.New("buckets", attr.AssetImage("assets/oci/storage/buckets.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) StorageGatewayWhite(opts ...attr.Attribute) *node.Node {
	return node.New("storage-gateway-white", attr.AssetImage("assets/oci/storage/storage-gateway-white.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) BackupRestoreWhite(opts ...attr.Attribute) *node.Node {
	return node.New("backup-restore-white", attr.AssetImage("assets/oci/storage/backup-restore-white.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) DataTransferWhite(opts ...attr.Attribute) *node.Node {
	return node.New("data-transfer-white", attr.AssetImage("assets/oci/storage/data-transfer-white.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) DataTransfer(opts ...attr.Attribute) *node.Node {
	return node.New("data-transfer", attr.AssetImage("assets/oci/storage/data-transfer.png"), attr.Shape(attr.None))
}

func (c *ociStorageContainer) ObjectStorageWhite(opts ...attr.Attribute) *node.Node {
	return node.New("object-storage-white", attr.AssetImage("assets/oci/storage/object-storage-white.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "oci.storage"
  Register(namespace, "FileStorageWhite", OciStorage.FileStorageWhite)
  Register(namespace, "FileStorage", OciStorage.FileStorage)
  Register(namespace, "ObjectStorage", OciStorage.ObjectStorage)
  Register(namespace, "StorageGateway", OciStorage.StorageGateway)
  Register(namespace, "BlockStorageCloneWhite", OciStorage.BlockStorageCloneWhite)
  Register(namespace, "BlockStorageClone", OciStorage.BlockStorageClone)
  Register(namespace, "BlockStorage", OciStorage.BlockStorage)
  Register(namespace, "ElasticPerformanceWhite", OciStorage.ElasticPerformanceWhite)
  Register(namespace, "ElasticPerformance", OciStorage.ElasticPerformance)
  Register(namespace, "BackupRestore", OciStorage.BackupRestore)
  Register(namespace, "BlockStorageWhite", OciStorage.BlockStorageWhite)
  Register(namespace, "BucketsWhite", OciStorage.BucketsWhite)
  Register(namespace, "Buckets", OciStorage.Buckets)
  Register(namespace, "StorageGatewayWhite", OciStorage.StorageGatewayWhite)
  Register(namespace, "BackupRestoreWhite", OciStorage.BackupRestoreWhite)
  Register(namespace, "DataTransferWhite", OciStorage.DataTransferWhite)
  Register(namespace, "DataTransfer", OciStorage.DataTransfer)
  Register(namespace, "ObjectStorageWhite", OciStorage.ObjectStorageWhite)
}
