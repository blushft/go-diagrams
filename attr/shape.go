package attr

import "fmt"

type ShapeType string

func (s ShapeType) String() string {
	return string(s)
}

var (
	Box               ShapeType = "box"
	Ellipse           ShapeType = "ellipse"
	Oval              ShapeType = "oval"
	Circle            ShapeType = "circle"
	Egg               ShapeType = "egg"
	Triangle          ShapeType = "triangle"
	Plain             ShapeType = "plain"
	Diamond           ShapeType = "diamond"
	Trapezium         ShapeType = "trapezium"
	Parallelogram     ShapeType = "parallelogram"
	House             ShapeType = "house"
	Pentagon          ShapeType = "pentagon"
	Hexagon           ShapeType = "hexagon"
	Septagon          ShapeType = "septagon"
	Octagon           ShapeType = "Octagon"
	DoubleCircle      ShapeType = "doublecircle"
	DoubleOctagon     ShapeType = "doubleoctagon"
	TripleOctagon     ShapeType = "tripleOctagon"
	InvertedTriangle  ShapeType = "invtriangle"
	InvertedTrapezium ShapeType = "invtrapezium"
	InvertedHouse     ShapeType = "invhouse"
	MDiamond          ShapeType = "mdiamond"
	MSquare           ShapeType = "msquare"
	MCircle           ShapeType = "mcircle"
	Rectangle         ShapeType = "rect"
	Square            ShapeType = "square"
	Star              ShapeType = "star"
	None              ShapeType = "none"
	Underline         ShapeType = "underline"
	Cylinder          ShapeType = "cylinder"
	Note              ShapeType = "note"
	Tab               ShapeType = "tab"
	Folder            ShapeType = "folder"
	Box3D             ShapeType = "box3d"
	Component         ShapeType = "component"
	Promoter          ShapeType = "promoter"
	CDS               ShapeType = "cds"
	Terminator        ShapeType = "terminator"
	UTR               ShapeType = "utr"
	PrimerSite        ShapeType = "primersite"
	RestrictionSite   ShapeType = "restrictionsite"
	FivePOverhang     ShapeType = "fivepoverhang"
	ThreePOverhang    ShapeType = "threepoverhang"
	NOverhang         ShapeType = "noverhang"
	Assembly          ShapeType = "assembly"
	Signature         ShapeType = "signature"
	Insulator         ShapeType = "insulator"
	Ribosite          ShapeType = "ribosite"
	RNASTab           ShapeType = "rnastab"
	ProteaseSite      ShapeType = "proteasesite"
	ProteinsTab       ShapeType = "proteinstab"
	RPromoter         ShapeType = "rpromoter"
	RArrow            ShapeType = "rarrow"
	LArrow            ShapeType = "larrow"
	LPromoter         ShapeType = "lpromoter"
)

type Point struct {
	X         float64
	Y         float64
	Z         float64
	InputOnly bool
}

func NewPoint(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

func (p Point) String() string {
	s := fmt.Sprintf("%f,%f", p.X, p.Y)

	if p.Z > 0 {
		s = fmt.Sprintf("%s,%f", s, p.Z)
	}

	if p.InputOnly {
		return s + "!"
	}

	return s
}

type Rect struct {
	LL Point
	UR Point
}

func NewRect(ll, ur Point) Rect {
	return Rect{
		LL: ll,
		UR: ur,
	}
}

func (r Rect) String() string {
	return fmt.Sprintf("%f,%f,%f,%f", r.LL.X, r.LL.Y, r.UR.X, r.UR.Y)
}

type SmoothType string

func (s SmoothType) String() string {
	return string(s)
}

const (
	AvgDistSmoothing   SmoothType = "avg_dist"
	GraphDistSmoothing SmoothType = "graph_dist"
	PowerDistSmoothing SmoothType = "power_dist"
	RNGSmoothing       SmoothType = "rng"
	SpringSmoothing    SmoothType = "spring"
	TriangleSmoothing  SmoothType = "triangle"
)

type SplineType string

func (s SplineType) String() string {
	return string(s)
}

const (
	SplineNone     SplineType = "none"
	SplineLines    SplineType = "lines"
	SplinePolyline SplineType = "polyline"
	SplineCurved   SplineType = "curved"
	SplineOrtho    SplineType = "ortho"
	SplineSpline   SplineType = "spline"
)
