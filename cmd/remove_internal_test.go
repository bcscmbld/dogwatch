package cmd

import "testing"
import "github.com/gugahoi/dogwatch/pkg/dognzb"

func TestRemove(t *testing.T) {
	testCases := []struct {
		desc string
		kind dognzb.Type
		ids  []string
	}{
		{
			desc: "",
			kind: dognzb.Movies,
			ids:  []string{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			remove(tC.kind, tC.ids)
		})
	}
}
