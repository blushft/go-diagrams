package font

type Options struct {
	Name  string
	Size  int
	Color string
}

func DefaultOptions() Options {
	return Options{
		Name:  "Sans-Serif",
		Size:  15,
		Color: "#2D3436",
	}
}
