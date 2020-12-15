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
	opts = append(opts, attr.AssetImage("assets/oci/storage/file-storage-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("file-storage-white", opts...)
}

func (c *ociStorageContainer) FileStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/file-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("file-storage", opts...)
}

func (c *ociStorageContainer) ObjectStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/object-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("object-storage", opts...)
}

func (c *ociStorageContainer) StorageGateway(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/storage-gateway.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage-gateway", opts...)
}

func (c *ociStorageContainer) BlockStorageCloneWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/block-storage-clone-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("block-storage-clone-white", opts...)
}

func (c *ociStorageContainer) BlockStorageClone(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/block-storage-clone.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("block-storage-clone", opts...)
}

func (c *ociStorageContainer) BlockStorage(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/block-storage.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("block-storage", opts...)
}

func (c *ociStorageContainer) ElasticPerformanceWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/elastic-performance-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-performance-white", opts...)
}

func (c *ociStorageContainer) ElasticPerformance(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/elastic-performance.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("elastic-performance", opts...)
}

func (c *ociStorageContainer) BackupRestore(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/backup-restore.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("backup-restore", opts...)
}

func (c *ociStorageContainer) BlockStorageWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/block-storage-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("block-storage-white", opts...)
}

func (c *ociStorageContainer) BucketsWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/buckets-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("buckets-white", opts...)
}

func (c *ociStorageContainer) Buckets(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/buckets.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("buckets", opts...)
}

func (c *ociStorageContainer) StorageGatewayWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/storage-gateway-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("storage-gateway-white", opts...)
}

func (c *ociStorageContainer) BackupRestoreWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/backup-restore-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("backup-restore-white", opts...)
}

func (c *ociStorageContainer) DataTransferWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/data-transfer-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-transfer-white", opts...)
}

func (c *ociStorageContainer) DataTransfer(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/data-transfer.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("data-transfer", opts...)
}

func (c *ociStorageContainer) ObjectStorageWhite(opts ...attr.Attribute) *node.Node {
	opts = append(opts, attr.AssetImage("assets/oci/storage/object-storage-white.png"))
	opts = append(opts, attr.Shape(attr.None))
	return node.New("object-storage-white", opts...)
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
