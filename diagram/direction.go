package diagram

type DiagramDirection string

const (
	TopToBottom DiagramDirection = "TB"
	BottomToTop DiagramDirection = "BT"
	LeftToRight DiagramDirection = "LR"
	RightToLeft DiagramDirection = "RL"
)

func Directions() []DiagramDirection {
	return []DiagramDirection{TopToBottom, BottomToTop, LeftToRight, RightToLeft}
}

func validateDirection(d DiagramDirection) bool {
	dirs := Directions()

	for _, vd := range dirs {
		if vd == d {
			return true
		}
	}

	return false
}

func curveSytles() []string {
	return []string{"ortho", "curved"}
}

func outputFormats() []string {
	return []string{"png", "jpg", "svg", "pdf"}
}
