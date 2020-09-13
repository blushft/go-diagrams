package aws

import "github.com/blushft/go-diagrams/diagram"

type blockchainContainer struct {
	path string
	opts []diagram.NodeOption
}

var Blockchain = &blockchainContainer{
	opts: diagram.OptionSet{diagram.Provider("aws"), diagram.NodeShape("none")},
	path: "assets/aws/blockchain",
}

func (c *blockchainContainer) ManagedBlockchain(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/blockchain/managed-blockchain.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}

func (c *blockchainContainer) QuantumLedgerDatabaseQldb(opts ...diagram.NodeOption) *diagram.Node {
	nopts := diagram.MergeOptionSets(diagram.OptionSet{diagram.Icon("assets/aws/blockchain/quantum-ledger-database-qldb.png")}, c.opts, opts)
	return diagram.NewNode(nopts...)
}
