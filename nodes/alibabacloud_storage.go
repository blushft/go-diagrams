package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Storage = &storageContainer{path: "assets/alibabacloud/storage"}

func (c *storageContainer) ObjectStorageService(opts ...attr.Attribute) *node.Node {
	return node.New("object-storage-service", attr.AssetImage("assets/alibabacloud/storage/object-storage-service.png"), attr.Shape(attr.None))
}

func (c *storageContainer) ObjectTableStore(opts ...attr.Attribute) *node.Node {
	return node.New("object-table-store", attr.AssetImage("assets/alibabacloud/storage/object-table-store.png"), attr.Shape(attr.None))
}

func (c *storageContainer) CloudStorageGateway(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-storage-gateway", attr.AssetImage("assets/alibabacloud/storage/cloud-storage-gateway.png"), attr.Shape(attr.None))
}

func (c *storageContainer) FileStorageHdfs(opts ...attr.Attribute) *node.Node {
	return node.New("file-storage-hdfs", attr.AssetImage("assets/alibabacloud/storage/file-storage-hdfs.png"), attr.Shape(attr.None))
}

func (c *storageContainer) FileStorageNas(opts ...attr.Attribute) *node.Node {
	return node.New("file-storage-nas", attr.AssetImage("assets/alibabacloud/storage/file-storage-nas.png"), attr.Shape(attr.None))
}

func (c *storageContainer) HybridBackupRecovery(opts ...attr.Attribute) *node.Node {
	return node.New("hybrid-backup-recovery", attr.AssetImage("assets/alibabacloud/storage/hybrid-backup-recovery.png"), attr.Shape(attr.None))
}

func (c *storageContainer) HybridCloudDisasterRecovery(opts ...attr.Attribute) *node.Node {
	return node.New("hybrid-cloud-disaster-recovery", attr.AssetImage("assets/alibabacloud/storage/hybrid-cloud-disaster-recovery.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Imm(opts ...attr.Attribute) *node.Node {
	return node.New("imm", attr.AssetImage("assets/alibabacloud/storage/imm.png"), attr.Shape(attr.None))
}
