package diagram

type NodeData struct {
	Name            string     `yaml:"name"`
	Label           string     `yaml:"label"`
	Type            string     `yaml:"type"`
	Nodes           []NodeData `yaml:"nodes"`
	ConnectTo       []string   `yaml:"connectTo"`
	ConnectFrom     []string   `yaml:"connectFrom"`
	ConnectAllTo    []string   `yaml:"connectAllTo"`
	ConnectAllFrom  []string   `yaml:"connectAllFrom"`
	BackgroundColor string     `yaml:"backgroundColor"`
}

type DiagramData struct {
	Direction string     `yaml:"direction" default:"LTR"`
	Label     string     `yaml:"label"`
	Nodes     []NodeData `yaml:"nodes"`
}
