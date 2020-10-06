package attr

type ClusterMode string

func (c ClusterMode) String() string {
	return string(c)
}

var (
	ClusterLocal  ClusterMode = "local"
	ClusterGlobal ClusterMode = "global"
	ClusterNone   ClusterMode = "none"
)
