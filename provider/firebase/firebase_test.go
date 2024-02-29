package firebase

import (
	"testing"

	"github.com/kaijuci/apk-distributor/distributor"
)

func TestFirebaseDistribute(t *testing.T) {
	expected := "expected"

	p, err := NewFirebaseDistributor()
	if err != nil {
		t.Fatal(err)
	}

	result, err := p.Distribute(&distributor.APKRelease{})
	if err != nil {
		t.Fatal(err)
	}

	if result != expected {
		t.Errorf("got: %s, want: %s", result, expected)
	}
}
