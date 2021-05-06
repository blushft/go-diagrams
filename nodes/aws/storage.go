package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type storageContainer struct {
	path  string
	attrs []attr.Attribute
}

var Storage = &storageContainer{path: "assets/aws/storage"}

func (c *storageContainer) Storage(opts ...attr.Attribute) *node.Node {
	return node.New("storage", attr.AssetImage("assets/aws/storage/storage.png"), attr.Shape(attr.None))
}

func (c *storageContainer) CloudendureDisasterRecovery(opts ...attr.Attribute) *node.Node {
	return node.New("cloudendure-disaster-recovery", attr.AssetImage("assets/aws/storage/cloudendure-disaster-recovery.png"), attr.Shape(attr.None))
}

func (c *storageContainer) ElasticFileSystemEfs(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-file-system-efs", attr.AssetImage("assets/aws/storage/elastic-file-system-efs.png"), attr.Shape(attr.None))
}

func (c *storageContainer) FsxForWindowsFileServer(opts ...attr.Attribute) *node.Node {
	return node.New("fsx-for-windows-file-server", attr.AssetImage("assets/aws/storage/fsx-for-windows-file-server.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Snowball(opts ...attr.Attribute) *node.Node {
	return node.New("snowball", attr.AssetImage("assets/aws/storage/snowball.png"), attr.Shape(attr.None))
}

func (c *storageContainer) StorageGateway(opts ...attr.Attribute) *node.Node {
	return node.New("storage-gateway", attr.AssetImage("assets/aws/storage/storage-gateway.png"), attr.Shape(attr.None))
}

func (c *storageContainer) EfsInfrequentaccessPrimaryBg(opts ...attr.Attribute) *node.Node {
	return node.New("efs-infrequentaccess-primary-bg", attr.AssetImage("assets/aws/storage/efs-infrequentaccess-primary-bg.png"), attr.Shape(attr.None))
}

func (c *storageContainer) ElasticBlockStoreEbs(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-block-store-ebs", attr.AssetImage("assets/aws/storage/elastic-block-store-ebs.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Fsx(opts ...attr.Attribute) *node.Node {
	return node.New("fsx", attr.AssetImage("assets/aws/storage/fsx.png"), attr.Shape(attr.None))
}

func (c *storageContainer) S3Glacier(opts ...attr.Attribute) *node.Node {
	return node.New("s3-glacier", attr.AssetImage("assets/aws/storage/s3-glacier.png"), attr.Shape(attr.None))
}

func (c *storageContainer) FsxForLustre(opts ...attr.Attribute) *node.Node {
	return node.New("fsx-for-lustre", attr.AssetImage("assets/aws/storage/fsx-for-lustre.png"), attr.Shape(attr.None))
}

func (c *storageContainer) SimpleStorageServiceS3(opts ...attr.Attribute) *node.Node {
	return node.New("simple-storage-service-s3", attr.AssetImage("assets/aws/storage/simple-storage-service-s3.png"), attr.Shape(attr.None))
}

func (c *storageContainer) SnowballEdge(opts ...attr.Attribute) *node.Node {
	return node.New("snowball-edge", attr.AssetImage("assets/aws/storage/snowball-edge.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Snowmobile(opts ...attr.Attribute) *node.Node {
	return node.New("snowmobile", attr.AssetImage("assets/aws/storage/snowmobile.png"), attr.Shape(attr.None))
}

func (c *storageContainer) Backup(opts ...attr.Attribute) *node.Node {
	return node.New("backup", attr.AssetImage("assets/aws/storage/backup.png"), attr.Shape(attr.None))
}

func (c *storageContainer) EfsStandardPrimaryBg(opts ...attr.Attribute) *node.Node {
	return node.New("efs-standard-primary-bg", attr.AssetImage("assets/aws/storage/efs-standard-primary-bg.png"), attr.Shape(attr.None))
}
