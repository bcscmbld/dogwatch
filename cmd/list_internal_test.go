package cmd

import "testing"
import "github.com/gugahoi/dogwatch/pkg/dognzb"

func TestList(t *testing.T) {
	testCases := []struct {
		desc string
		kind dognzb.Type
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			list(tC.kind)
		})
	}
}
