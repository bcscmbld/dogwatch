package dognzb_test

import "testing"
import "github.com/gugahoi/dogwatch/pkg/dognzb"

func TestNew(t *testing.T) {
	testCases := []struct {
		desc string
		api  string
	}{
		{
			desc: "with api", api: "some-api",
		},
		{
			desc: "empty api", api: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			d := dognzb.New(tC.api)
			if d == nil {
				t.Errorf("Expected a dognzb object, got nil")
			}
		})
	}
}
