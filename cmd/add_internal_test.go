package cmd

import (
	"testing"

	"github.com/gugahoi/dogwatch/pkg/dognzb"
)

func TestAdd(t *testing.T) {
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
			add(tC.kind, tC.ids)
		})
	}
}
