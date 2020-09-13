package diagram

type Font struct {
	Name  string
	Size  float64
	Color string
}

func defaultFont() Font {
	return Font{
		Name:  "Sans-Serif",
		Size:  13,
		Color: "#2D3436",
	}
}

func trimAttrs(attrs map[string]string) map[string]string {
	res := map[string]string{}

	for k, v := range attrs {
		if v != "" {
			res[k] = v
		}
	}

	return res
}
