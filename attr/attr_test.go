package attr

import (
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestRenderHook(t *testing.T) {
	attrs := NewAttributes(URL("http://testing.com"))

	h := func(a Attribute) (Attribute, error) {
		v := a.Value().String()

		_, err := url.Parse(v)
		if err != nil {
			return nil, err
		}

		return a, nil
	}

	_, err := attrs.Render(h)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRenderAsset(t *testing.T) {
	attrs := NewAttributes(AssetImage("assets/apps/analytics/beam.png"))

	out := filepath.Join(os.TempDir(), "diagram")
	t.Logf("writing to %s", out)

	m, err := attrs.Render(AssetLoader(out))
	if err != nil {
		t.Fatal(err)
	}

	spew.Dump(m)
}
