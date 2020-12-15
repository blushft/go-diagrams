package attr

import (
	graphviz "github.com/awalterschulze/gographviz"
	"github.com/blushft/go-diagrams/attr/color"
)

func WithAttributes(attrs ...Attribute) Attribute {
	return newAttrSet(attrs...)
}

func Damping(f float64) Attribute {
	return newAttr("Damping", f, Graph)
}

func K(f float64) Attribute {
	return newAttr("K", f, Graph, Clusters)
}

func URL(s string) Attribute {
	return newAttr("URL", s, Edges, Nodes, Graph, Clusters)
}

func Background(s string) Attribute {
	return newAttr("_background", s, Graph)
}

func Area(area float64) Attribute {
	return newAttr("area", area, Nodes, Clusters)
}

func Arrowhead(a Arrow) Attribute {
	return newAttr("arrowhead", string(a), Edges)
}

func ArrowSize(f float64) Attribute {
	return newAttr("arrowsize", f, Edges)
}

func ArrowTail(a Arrow) Attribute {
	return newAttr("arrowtail", string(a), Edges)
}

func BoundingBox(r Rect) Attribute {
	return newAttr("bb", r, Graph)
}

func BGColor(c string) Attribute {
	return newAttr("bgcolor", c, Graph, Clusters)
}

func Center(b bool) Attribute {
	return newAttr("center", b, Graph)
}

func Charset(s string) Attribute {
	return newAttr("charset", s, Graph)
}

func Class(s string) Attribute {
	return newAttr("class", s, Edges, Nodes, Clusters, Graph)
}

func ClusterRank(c ClusterMode) Attribute {
	return newAttr("clusterrank", c, Graph)
}

func Color(c color.Color) Attribute {
	return newAttr("color", c, Edges, Nodes, Clusters)
}

func ColorScheme(s string) Attribute {
	return newAttr("colorscheme", s, Edges, Nodes, Clusters, Graph)
}

func Comment(s string) Attribute {
	return newAttr("comment", s, Edges, Nodes, Graph)
}

func Compound(b bool) Attribute {
	return newAttr("compound", b, Graph)
}

func Concentrate(b bool) Attribute {
	return newAttr("concentrate", b, Graph)
}

func Constraint(b bool) Attribute {
	return newAttr("constraint", b, Edges)
}

func Decorate(b bool) Attribute {
	return newAttr("decorate", b, Edges)
}

func DefaultDist(f float64) Attribute {
	return newAttr("defaultdist", f, Graph)
}

func Dim(i int) Attribute {
	return newAttr("dim", i, Graph)
}

func Dimen(i int) Attribute {
	return newAttr("dimen", i, Graph)
}

func Direction(d EdgeDirection) Attribute {
	return newAttr("dir", d, Edges)
}

func Directed(d bool) Attribute {
	return newAttr("dir", d, Graph)
}

func DirEdgeConstraints(b bool) Attribute {
	return newAttr("diredgeconstraints", b, Graph)
}

func Distortion(f float64) Attribute {
	return newAttr("distortion", f, Nodes)
}

func DPI(f float64) Attribute {
	return newAttr("dpi", f, Graph)
}

func EdgeURL(s string) Attribute {
	return newAttr("edgeURL", s, Edges)
}

func EdgeHREF(s string) Attribute {
	return newAttr("edgehref", s, Edges)
}

func EdgeTarget(s string) Attribute {
	return newAttr("edgetarget", s, Edges)
}

func EdgeTooltip(s string) Attribute {
	return newAttr("edgetooltip", s, Edges)
}

func Epsilon(f float64) Attribute {
	return newAttr("epsilon", f, Graph)
}

func ESep(i int) Attribute {
	return newAttr("esep", i, Graph)
}

func FillColor(c color.Color) Attribute {
	return newAttr("fillcolor", c, Nodes, Edges, Clusters)
}

func FixedSize(b bool) Attribute {
	return newAttr("fixedsize", b, Nodes)
}

func FontColor(c color.Color) Attribute {
	return newAttr("fontcolor", c, Edges, Nodes, Graph, Clusters)
}

func FontName(s string) Attribute {
	return newAttr("fontname", s, Edges, Nodes, Graph, Clusters)
}

func FontNames(s ...string) Attribute {
	return newAttr("fontnames", s, Graph)
}

func FontPath(s string) Attribute {
	return newAttr("fontpath", s, Graph)
}

func FontSize(f float64) Attribute {
	return newAttr("fontsize", f, Edges, Nodes, Graph, Clusters)
}

func ForceLabels(b bool) Attribute {
	return newAttr("forcelabels", b, Graph)
}

func GradientAngle(i int) Attribute {
	return newAttr("gradientangle", i, Nodes, Clusters, Graph)
}

func Group(s string) Attribute {
	return newAttr("group", s, Nodes)
}

func HeadURL(s string) Attribute {
	return newAttr("headURL", s, Edges)
}

//HeadLabelPosition is the position of an edge's head label in points. Position indicates the center of the label.
func HeadLabelPosition(p Point) Attribute {
	return newAttr("head_lp", p, Edges)
}

func HeadClip(b bool) Attribute {
	return newAttr("headclip", b, Edges)
}

func HeadHREF(s string) Attribute {
	return newAttr("headhref", s, Edges)
}

func HeadLabel(s string) Attribute {
	return newAttr("headlabel", s, Edges)
}

func HeadPort(p Point) Attribute {
	return newAttr("headport", p, Edges)
}

func HeadTarget(s string) Attribute {
	return newAttr("headtarget", s, Edges)
}

func HeadTooltip(s string) Attribute {
	return newAttr("headtooltip", s, Edges)
}

func Height(f float64) Attribute {
	return newAttr("height", f, Nodes)
}

func HREF(s string) Attribute {
	return newAttr("href", s, Graph, Clusters, Nodes, Edges)
}

func ID(s string) Attribute {
	return newAttr("id", s, Graph, Clusters, Nodes, Edges)
}

func Image(s string) Attribute {
	return newAttr("image", s, Nodes)
}

func AssetImage(s string) Attribute {
	return newAttr(graphviz.Image, s, Nodes)
}

func ImagePath(s string) Attribute {
	return newAttr("imagepath", s, Graph)
}

func ImagePosition(p Position) Attribute {
	return newAttr("imagepos", p, Nodes)
}

func ImageScale(s string) Attribute {
	return newAttr("imagescale", s, Nodes)
}

func InputScale(f float64) Attribute {
	return newAttr("inputscale", f, Graph)
}

func Label(s string) Attribute {
	return newAttr("label", s, Edges, Nodes, Graph, Clusters)
}

func LabelURL(s string) Attribute {
	return newAttr("labelURL", s, Edges)
}

func LabelScheme(i int) Attribute {
	return newAttr("label_scheme", i, Graph)
}

func LabelAngle(f float64) Attribute {
	return newAttr("labelangle", f, Edges)
}

func LabelDistance(f float64) Attribute {
	return newAttr("labeldistance", f, Edges)
}

func LabelFloat(b bool) Attribute {
	return newAttr("labelfloat", b, Edges)
}

func LabelFontColor(c color.Color) Attribute {
	return newAttr("labelfontcolor", c, Edges)
}

func LabelFontName(s string) Attribute {
	return newAttr("labelfontname", s, Edges)
}

func LabelFontSize(f float64) Attribute {
	return newAttr("labelfontsize", f, Edges)
}

func LabelHREF(s string) Attribute {
	return newAttr("labelhref", s, Edges)
}

func LabelJustify(j Justify) Attribute {
	return newAttr("labeljust", j, Graph, Clusters)
}

func LabelLocation(l Location) Attribute {
	return newAttr("labelloc", l, Nodes, Graph, Clusters)
}

func LabelTarget(s string) Attribute {
	return newAttr("labeltarget", s, Edges)
}

func LabelTooltip(s string) Attribute {
	return newAttr("labeltooltip", s, Edges)
}

func Landscape(b bool) Attribute {
	return newAttr("landscape", b, Graph)
}

//Layer is not yet implemented and will panic
//TODO: figure out how to manage layers
func Layer(s ...string) Attribute {
	panic("layers not yet implemented")
}

func LayerListSep(s ...string) Attribute {
	panic("layers not yet implemented")
}

func Layers(s ...string) Attribute {
	panic("layers not yet implemented")
}

func LayerSelect(s ...string) Attribute {
	panic("layers not yet implemented")
}

func LayerSep(s ...string) Attribute {
	panic("layers not yet implemented")
}

func Layout(s string) Attribute {
	return newAttr("layout", s, Graph)
}

func Length(f float64) Attribute {
	return newAttr("len", f, Edges)
}

func Levels(i int) Attribute {
	return newAttr("levels", i, Graph)
}

func LevelsGap(f float64) Attribute {
	return newAttr("levelsgap", f, Graph)
}

func LogicalHead(s string) Attribute {
	return newAttr("lhead", s, Edges)
}

func LabelHeight(f float64) Attribute {
	return newAttr("lheight", f, Graph, Clusters)
}

func LabelPostion(p Point) Attribute {
	return newAttr("lp", p, Edges, Graph, Clusters)
}

func LogicalTail(s string) Attribute {
	return newAttr("ltail", s, Edges)
}

func LabelWidth(f float64) Attribute {
	return newAttr("lwidth", f, Graph, Clusters)
}

func Margin(f float64) Attribute {
	return newAttr("margin", f, Nodes, Graph, Clusters)
}

func MaxIter(i int) Attribute {
	return newAttr("maxiter", i, Graph)
}

func MCLimit(f float64) Attribute {
	return newAttr("mclimit", f, Graph)
}

func MinDist(f float64) Attribute {
	return newAttr("mindist", f, Graph)
}

func MinLen(i int) Attribute {
	return newAttr("minlen", i, Edges)
}

func Mode(s string) Attribute {
	return newAttr("mode", s, Graph)
}

func Model(s string) Attribute {
	return newAttr("model", s, Graph)
}

func Mosek(b bool) Attribute {
	return newAttr("mosek", b, Graph)
}

func NewRank(b bool) Attribute {
	return newAttr("newrank", b, Graph)
}

func NodeSeparation(f float64) Attribute {
	return newAttr("nodesep", f, Graph)
}

func NoJustify(b bool) Attribute {
	return newAttr("nojustify", b, Graph, Clusters, Nodes, Edges)
}

func Normalize(b bool) Attribute {
	return newAttr("normalize", b, Graph)
}

func NoTranslate(b bool) Attribute {
	return newAttr("notranslate", b, Graph)
}

func NSLimit(f float64) Attribute {
	return newAttr("nslimit", f, Graph)
}

func NSLimit1(f float64) Attribute {
	return newAttr("nslimit1", f, Graph)
}

func Ordering(s string) Attribute {
	return newAttr("ordering", s, Graph, Nodes)
}

func Orientation(f float64) Attribute {
	return newAttr("orientation", f, Nodes)
}

func OutputOrder(o OutputMode) Attribute {
	return newAttr("outputorder", o, Graph)
}

func Overlap(b bool) Attribute {
	return newAttr("overlap", b, Graph)
}

func OverlapScaling(f float64) Attribute {
	return newAttr("overlap_scaling", f, Graph)
}

func OverlapShrink(b bool) Attribute {
	return newAttr("overlap_shrink", b, Graph)
}

func Pack(b bool) Attribute {
	return newAttr("pack", b, Graph)
}

func PackMode(p PackModeAttr) Attribute {
	return newAttr("packmode", p, Graph)
}

func Pad(f float64) Attribute {
	return newAttr("pad", f, Graph)
}

func Page(f float64) Attribute {
	return newAttr("page", f, Graph)
}

func PageDir(p PageDirection) Attribute {
	return newAttr("pagedir", p, Graph)
}

func PenColor(c color.Color) Attribute {
	return newAttr("pencolor", c, Clusters)
}

func PenWidth(f float64) Attribute {
	return newAttr("penwidth", f, Clusters, Nodes, Edges)
}

func Peripheries(i int) Attribute {
	return newAttr("peripheries", i, Nodes, Clusters)
}

func Pin(b bool) Attribute {
	return newAttr("pin", b, Nodes)
}

func Pos(p Point) Attribute {
	return newAttr("pos", p, Edges, Nodes)
}

func QuadTree(b bool) Attribute {
	return newAttr("quadtree", b, Graph)
}

func Quantum(f float64) Attribute {
	return newAttr("quantum", f, Graph)
}

func Rank(r RankType) Attribute {
	return newAttr("rank", r, Subs)
}

func RankDir(r RankDirection) Attribute {
	return newAttr("rankdir", r, Graph)
}

func RankSeparation(f float64) Attribute {
	return newAttr("ranksep", f, Graph)
}

func Ration(f float64) Attribute {
	return newAttr("ration", f, Graph)
}

func Ratio(s string) Attribute {
	return newAttr("ratio", s, Graph)
}

func Rects(r Rect) Attribute {
	return newAttr("rects", r, Nodes)
}

func Regular(b bool) Attribute {
	return newAttr("regular", b, Nodes)
}

func ReMinCross(b bool) Attribute {
	return newAttr("remincross", b, Graph)
}

func RepulsiveForce(f float64) Attribute {
	return newAttr("repulsiveforce", f, Graph)
}

func Resolution(f float64) Attribute {
	return newAttr("resolution", f, Graph)
}

func Root(s string) Attribute {
	return newAttr("root", s, Graph, Nodes)
}

func Rotate(i int) Attribute {
	return newAttr("rotate", i, Graph)
}

func Rotation(f float64) Attribute {
	return newAttr("rotation", f, Graph)
}

func SameHead(s string) Attribute {
	return newAttr("samehead", s, Edges)
}

func SameTail(s string) Attribute {
	return newAttr("sametail", s, Edges)
}

func SamplePoints(i int) Attribute {
	return newAttr("samplepoints", i, Nodes)
}

func Scale(f float64) Attribute {
	return newAttr("scale", f, Graph)
}

func SearchSize(f float64) Attribute {
	return newAttr("searchsize", f, Graph)
}

func Sep(f float64) Attribute {
	return newAttr("sep", f, Graph)
}

func Shape(s ShapeType) Attribute {
	return newAttr("shape", s, Nodes)
}

func ShapeFile(s string) Attribute {
	return newAttr("shapefile", s, Nodes)
}

func ShowBoxes(i int) Attribute {
	return newAttr("showboxes", i, Edges, Nodes, Graph)
}

func Sides(i int) Attribute {
	return newAttr("sides", i, Nodes)
}

func Size(f float64) Attribute {
	return newAttr("size", f, Graph)
}

func Skew(f float64) Attribute {
	return newAttr("skew", f, Nodes)
}

func Smoothing(s SmoothType) Attribute {
	return newAttr("smoothing", s, Graph)
}

func SortV(i int) Attribute {
	return newAttr("sortv", i, Graph, Clusters, Nodes)
}

func Splines(s SplineType) Attribute {
	return newAttr("splines", s, Graph)
}

func Start(s StartType) Attribute {
	return newAttr("start", s, Graph)
}

func Style(s StyleType) Attribute {
	return newAttr("style", s, Edges, Nodes, Clusters, Graph)
}

func Stylesheet(s string) Attribute {
	return newAttr("stylesheet", s, Graph)
}

func TailURL(s string) Attribute {
	return newAttr("tailURL", s, Edges)
}

func TailLabelPosition(l Position) Attribute {
	return newAttr("tail_lp", l, Edges)
}

func TailClip(b bool) Attribute {
	return newAttr("tailclip", b, Edges)
}

func TailHREF(s string) Attribute {
	return newAttr("tailhref", s, Edges)
}

func TailLabel(s string) Attribute {
	return newAttr("taillabel", s, Edges)
}

func TailPort(p PortPosition) Attribute {
	return newAttr("tailport", p, Edges)
}

func TailTarget(s string) Attribute {
	return newAttr("tailtarget", s, Edges)
}

func TailTooltip(s string) Attribute {
	return newAttr("tailtooltip", s, Edges)
}

func Target(s string) Attribute {
	return newAttr("target", s, Edges, Nodes, Graph, Clusters)
}

func Tooltip(s string) Attribute {
	return newAttr("tooltip", s, Nodes, Edges, Clusters)
}

func TrueColor(b bool) Attribute {
	return newAttr("truecolor", b, Graph)
}

func Vertices(p ...Point) Attribute {
	return newAttr("vertices", p, Nodes)
}

func Viewport(v ViewportDim) Attribute {
	return newAttr("viewport", v, Graph)
}

func VoroMargion(f float64) Attribute {
	return newAttr("voro_margin", f, Graph)
}

func Weight(f float64) Attribute {
	return newAttr("weight", f, Edges)
}

func Width(f float64) Attribute {
	return newAttr("width", f, Nodes)
}

func XDotVersion(s string) Attribute {
	return newAttr("xdotversion", s, Graph)
}

func XLabel(s string) Attribute {
	return newAttr("xlabel", s, Nodes, Edges)
}

func XLP(p Point) Attribute {
	return newAttr("xlp", p, Nodes, Edges)
}

func Z(f float64) Attribute {
	return newAttr("z", f, Nodes)
}
