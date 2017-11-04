package dognzb_test

import "testing"
import "github.com/gugahoi/dogwatch/pkg/dognzb"

func TestGetIDTV(t *testing.T) {
	// arrange
	i := &dognzb.Item{
		TVdbID: "7331",
	}

	// assign
	id := i.GetID()

	// assert
	if id != "7331" {
		t.Errorf("expected id to be '7331', got '%v'", id)
	}
}

func TestGetIDMovie(t *testing.T) {
	// arrange
	i := &dognzb.Item{
		ImdbID: "800813",
	}

	// assing
	id := i.GetID()

	// assert
	if id != "tt800813" {
		t.Errorf("expected id to be 'tt800813', got '%v'", id)
	}
}
