package attr

import "fmt"

type EdgeDirection string

func (d EdgeDirection) String() string {
	return string(d)
}

var (
	Forward  EdgeDirection = "forward"
	Back     EdgeDirection = "back"
	Both     EdgeDirection = "both"
	EdgeNone EdgeDirection = "none"
)

type PageDirection string

func (d PageDirection) String() string {
	return string(d)
}

var (
	PageBottomLeft  PageDirection = "BL"
	PageBottomRight PageDirection = "BR"
	PageTopLeft     PageDirection = "TL"
	PageTopRight    PageDirection = "TR"
	PageRightBottom PageDirection = "RB"
	PageRightTop    PageDirection = "RT"
	PageLeftBottom  PageDirection = "LB"
	PageLeftTop     PageDirection = "LT"
)

type Position string

func (p Position) String() string {
	return string(p)
}

var (
	TopLeft        Position = "tl"
	TopCentered    Position = "tc"
	TopRight       Position = "tr"
	MiddleLeft     Position = "ml"
	MiddleCentered Position = "mc"
	MiddleRight    Position = "mr"
	BottomLeft     Position = "bl"
	BottomCentered Position = "bc"
	BottonRight    Position = "br"
)

type PortPosition string

func (p PortPosition) String() string {
	return string(p)
}

const (
	North      PortPosition = "n"
	NorthEast  PortPosition = "ne"
	East       PortPosition = "e"
	SouthEast  PortPosition = "se"
	South      PortPosition = "s"
	SoutWest   PortPosition = "sw"
	West       PortPosition = "w"
	NorthWest  PortPosition = "nw"
	CenterPort PortPosition = "c"
)

type Justify string

func (j Justify) String() string {
	return string(j)
}

var (
	JustifyLeft   = "l"
	JustifyCenter = "c"
	JustifyRight  = "r"
)

type Location string

func (l Location) String() string {
	return string(l)
}

var (
	Top       Location = "t"
	Bottom    Location = "b"
	LocCenter Location = "c"
)

type OutputMode string

func (o OutputMode) String() string {
	return string(o)
}

var (
	BreadthFirst OutputMode = "breadthfirst"
	NodesFirst   OutputMode = "nodesfirst"
	EdgesFirst   OutputMode = "edgesfirst"
)

type PackModeAttr string

func (p PackModeAttr) String() string {
	return string(p)
}

var (
	PackModeNode    PackModeAttr = "node"
	PackModeCluster PackModeAttr = "cluster"
	PackModeGraph   PackModeAttr = "graph"
)

type StartType string

func (s StartType) String() string {
	return string(s)
}

const (
	StartRegular StartType = "regular"
	StartSelf    StartType = "self"
	StartRandom  StartType = "random"
)

type ViewportDim struct {
	Width  float64
	Height float64
	Zoom   float64
	X      float64
	Y      float64
}

func (v ViewportDim) String() string {
	return fmt.Sprintf("%f,%f,%f,%f,%f", v.Width, v.Height, v.Zoom, v.X, v.Y)
}
