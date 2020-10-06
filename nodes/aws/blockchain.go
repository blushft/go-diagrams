package aws

import (
	attr "github.com/blushft/go-diagrams/attr"
	"github.com/blushft/go-diagrams/node"
)

type blockchainContainer struct {
	path  string
	attrs []attr.Attribute
}

var Blockchain = &blockchainContainer{path: "assets/aws/blockchain"}

func (c *blockchainContainer) ManagedBlockchain(opts ...attr.Attribute) *node.Node {
	return node.New("managed-blockchain", attr.AssetImage("assets/aws/blockchain/managed-blockchain.png"), attr.Shape(attr.None))
}

func (c *blockchainContainer) QuantumLedgerDatabaseQldb(opts ...attr.Attribute) *node.Node {
	return node.New("quantum-ledger-database-qldb", attr.AssetImage("assets/aws/blockchain/quantum-ledger-database-qldb.png"), attr.Shape(attr.None))
}
