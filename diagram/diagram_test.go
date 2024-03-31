package diagram_test

import (
	"errors"
	"os"
	"testing"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/stretchr/testify/require"
)

func cleanupTestdata(t *testing.T) {
	// remove the auto-generated data from this test execution
	err := os.RemoveAll("./go-diagrams")
	if errors.Is(err, os.ErrNotExist) {
		return
	}

	require.NoError(t, err)
}

func TestRender(t *testing.T) {
	defer cleanupTestdata(t)

	for _, testcase := range []struct {
		name      string
		graphFunc func(diagram *diagram.Diagram)
		err       error
	}{
		{
			name: "Success/Simple",
			graphFunc: func(d *diagram.Diagram) {
				inet := apps.Network.Internet().Label("Internet")
				proxy := apps.Network.Caddy().Label("Caddy")

				d.Connect(inet, proxy)
			},
		},
		{
			name: "Success/RerunOnUpdate",
			graphFunc: func(d *diagram.Diagram) {
				inet := apps.Network.Internet().Label("Internet")
				proxy := apps.Network.Caddy().Label("Caddy")

				d.Connect(inet, proxy)

				ss := apps.Inmemory.Redis().Label("session")
				rs := apps.Inmemory.Redis().Label("replica")

				cache := diagram.NewGroup("cache").Label("Sessions").
					Connect(ss, rs, diagram.Bidirectional())

				d.Group(cache)
			},
		},
	} {
		t.Run(testcase.name, func(t *testing.T) {
			d, err := diagram.New(diagram.Label("test"), diagram.Direction("LR"))

			require.True(t, errors.Is(err, testcase.err))

			if testcase.graphFunc != nil {
				testcase.graphFunc(d)
			}

			require.NoError(t, d.Render())
		})
	}
}
