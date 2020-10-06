package color

import (
	"fmt"
	"image/color"
)

type Color interface {
	color.Color
	String() string
}

type colorAttr struct {
	color.Color
}

func newColorAttr(c color.RGBA) colorAttr {
	return colorAttr{c}
}

func (c colorAttr) String() string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf("#%d%d%d%d", r, g, b, a)
}
