package assets

import (
	"testing"
)

func TestAssetList(t *testing.T) {
	di, err := ReadDir("assets/apps")
	if err != nil {
		t.Fatal(err)
	}

	if len(di) != 20 {
		t.Logf("unexpected dir info count: expected %d, got %d", 20, len(di))
	}

	_, err = ReadFile("assets/apps/client/client.png")
	if err != nil {
		t.Fatal(err)
	}

}
