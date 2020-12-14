package nodes

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type awsStorageContainer struct {
	path  string
	attrs []attr.Attribute
}

var AWSStorage = &awsStorageContainer{path: "assets/aws/storage"}

func (c *awsStorageContainer) S3Glacier(opts ...attr.Attribute) *node.Node {
	return node.New("s3-glacier", attr.AssetImage("assets/aws/storage/s3-glacier.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) EfsInfrequentaccessPrimaryBg(opts ...attr.Attribute) *node.Node {
	return node.New("efs-infrequentaccess-primary-bg", attr.AssetImage("assets/aws/storage/efs-infrequentaccess-primary-bg.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) EfsStandardPrimaryBg(opts ...attr.Attribute) *node.Node {
	return node.New("efs-standard-primary-bg", attr.AssetImage("assets/aws/storage/efs-standard-primary-bg.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) ElasticFileSystemEfs(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-file-system-efs", attr.AssetImage("assets/aws/storage/elastic-file-system-efs.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) FsxForWindowsFileServer(opts ...attr.Attribute) *node.Node {
	return node.New("fsx-for-windows-file-server", attr.AssetImage("assets/aws/storage/fsx-for-windows-file-server.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) ElasticBlockStoreEbs(opts ...attr.Attribute) *node.Node {
	return node.New("elastic-block-store-ebs", attr.AssetImage("assets/aws/storage/elastic-block-store-ebs.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) Storage(opts ...attr.Attribute) *node.Node {
	return node.New("storage", attr.AssetImage("assets/aws/storage/storage.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) Backup(opts ...attr.Attribute) *node.Node {
	return node.New("backup", attr.AssetImage("assets/aws/storage/backup.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) Fsx(opts ...attr.Attribute) *node.Node {
	return node.New("fsx", attr.AssetImage("assets/aws/storage/fsx.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) StorageGateway(opts ...attr.Attribute) *node.Node {
	return node.New("storage-gateway", attr.AssetImage("assets/aws/storage/storage-gateway.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) Snowball(opts ...attr.Attribute) *node.Node {
	return node.New("snowball", attr.AssetImage("assets/aws/storage/snowball.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) Snowmobile(opts ...attr.Attribute) *node.Node {
	return node.New("snowmobile", attr.AssetImage("assets/aws/storage/snowmobile.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) CloudendureDisasterRecovery(opts ...attr.Attribute) *node.Node {
	return node.New("cloudendure-disaster-recovery", attr.AssetImage("assets/aws/storage/cloudendure-disaster-recovery.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) FsxForLustre(opts ...attr.Attribute) *node.Node {
	return node.New("fsx-for-lustre", attr.AssetImage("assets/aws/storage/fsx-for-lustre.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) SimpleStorageServiceS3(opts ...attr.Attribute) *node.Node {
	return node.New("simple-storage-service-s3", attr.AssetImage("assets/aws/storage/simple-storage-service-s3.png"), attr.Shape(attr.None))
}

func (c *awsStorageContainer) SnowballEdge(opts ...attr.Attribute) *node.Node {
	return node.New("snowball-edge", attr.AssetImage("assets/aws/storage/snowball-edge.png"), attr.Shape(attr.None))
}

func init() {
  const namespace= "aws.storage"
  Register(namespace, "S3Glacier", AWSStorage.S3Glacier)
  Register(namespace, "EfsInfrequentaccessPrimaryBg", AWSStorage.EfsInfrequentaccessPrimaryBg)
  Register(namespace, "EfsStandardPrimaryBg", AWSStorage.EfsStandardPrimaryBg)
  Register(namespace, "ElasticFileSystemEfs", AWSStorage.ElasticFileSystemEfs)
  Register(namespace, "FsxForWindowsFileServer", AWSStorage.FsxForWindowsFileServer)
  Register(namespace, "ElasticBlockStoreEbs", AWSStorage.ElasticBlockStoreEbs)
  Register(namespace, "AWSStorage", AWSStorage.Storage)
  Register(namespace, "Backup", AWSStorage.Backup)
  Register(namespace, "Fsx", AWSStorage.Fsx)
  Register(namespace, "StorageGateway", AWSStorage.StorageGateway)
  Register(namespace, "Snowball", AWSStorage.Snowball)
  Register(namespace, "Snowmobile", AWSStorage.Snowmobile)
  Register(namespace, "CloudendureDisasterRecovery", AWSStorage.CloudendureDisasterRecovery)
  Register(namespace, "FsxForLustre", AWSStorage.FsxForLustre)
  Register(namespace, "SimpleStorageServiceS3", AWSStorage.SimpleStorageServiceS3)
  Register(namespace, "SnowballEdge", AWSStorage.SnowballEdge)
}
