package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type alibabaCloudStorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var AlibabacloudStorage =&alibabaCloudStorageContainer{path: "assets/alibabacloud/storage"}

func (c *alibabaCloudStorageContainer) ObjectStorageService(opts ...attr.Attribute) *node.Node {
	return node.New("object-storage-service", attr.AssetImage("assets/alibabacloud/storage/object-storage-service.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudStorageContainer) ObjectTableStore(opts ...attr.Attribute) *node.Node {
	return node.New("object-table-store", attr.AssetImage("assets/alibabacloud/storage/object-table-store.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudStorageContainer) CloudStorageGateway(opts ...attr.Attribute) *node.Node {
	return node.New("cloud-storage-gateway", attr.AssetImage("assets/alibabacloud/storage/cloud-storage-gateway.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudStorageContainer) FileStorageHdfs(opts ...attr.Attribute) *node.Node {
	return node.New("file-storage-hdfs", attr.AssetImage("assets/alibabacloud/storage/file-storage-hdfs.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudStorageContainer) FileStorageNas(opts ...attr.Attribute) *node.Node {
	return node.New("file-storage-nas", attr.AssetImage("assets/alibabacloud/storage/file-storage-nas.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudStorageContainer) HybridBackupRecovery(opts ...attr.Attribute) *node.Node {
	return node.New("hybrid-backup-recovery", attr.AssetImage("assets/alibabacloud/storage/hybrid-backup-recovery.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudStorageContainer) HybridCloudDisasterRecovery(opts ...attr.Attribute) *node.Node {
	return node.New("hybrid-cloud-disaster-recovery", attr.AssetImage("assets/alibabacloud/storage/hybrid-cloud-disaster-recovery.png"), attr.Shape(attr.None))
}

func (c *alibabaCloudStorageContainer) Imm(opts ...attr.Attribute) *node.Node {
	return node.New("imm", attr.AssetImage("assets/alibabacloud/storage/imm.png"), attr.Shape(attr.None))
}

func init() {
  const namespace = "alibabacloud.storage"
  Register(namespace, "ObjectStorageService", AlibabacloudStorage.ObjectStorageService)
  Register(namespace, "ObjectTableStore", AlibabacloudStorage.ObjectTableStore)
  Register(namespace, "CloudStorageGateway", AlibabacloudStorage.CloudStorageGateway)
  Register(namespace, "FileStorageHdfs", AlibabacloudStorage.FileStorageHdfs)
  Register(namespace, "FileStorageNas", AlibabacloudStorage.FileStorageNas)
  Register(namespace, "HybridBackupRecovery", AlibabacloudStorage.HybridBackupRecovery)
  Register(namespace, "HybridCloudDisasterRecovery", AlibabacloudStorage.HybridCloudDisasterRecovery)
  Register(namespace, "Imm", AlibabacloudStorage.Imm)
}
