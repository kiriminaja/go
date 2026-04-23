package volumetric_test

import (
	"testing"

	"github.com/kiriminaja/go/utils/volumetric"
)

func TestEmpty(t *testing.T) {
	got := volumetric.Calculate(nil)
	if got != (volumetric.Dimensions{}) {
		t.Fatalf("expected zero dimensions, got %+v", got)
	}
}

func TestSingleItem(t *testing.T) {
	got := volumetric.Calculate([]volumetric.Item{
		{Qty: 1, Length: 10, Width: 5, Height: 3},
	})
	want := volumetric.Dimensions{Length: 10, Width: 5, Height: 3}
	if got != want {
		t.Fatalf("got %+v want %+v", got, want)
	}
}

func TestPicksSmallestStacking(t *testing.T) {
	// Five flat-and-wide items default to vertical stacking (ties go to vertical).
	got := volumetric.Calculate([]volumetric.Item{
		{Qty: 2, Length: 10, Width: 10, Height: 2},
	})
	want := volumetric.Dimensions{Length: 10, Width: 10, Height: 4}
	if got != want {
		t.Fatalf("vertical: got %+v want %+v", got, want)
	}

	// Tall stack of thin items + a long thin filler — horizontal wins.
	got = volumetric.Calculate([]volumetric.Item{
		{Qty: 5, Length: 2, Width: 10, Height: 10},
		{Qty: 1, Length: 10, Width: 1, Height: 1},
	})
	want = volumetric.Dimensions{Length: 20, Width: 10, Height: 10}
	if got != want {
		t.Fatalf("horizontal: got %+v want %+v", got, want)
	}

	// Mirror of the above on the width axis — side-by-side wins.
	got = volumetric.Calculate([]volumetric.Item{
		{Qty: 5, Length: 10, Width: 2, Height: 10},
		{Qty: 1, Length: 1, Width: 10, Height: 1},
	})
	want = volumetric.Dimensions{Length: 10, Width: 20, Height: 10}
	if got != want {
		t.Fatalf("side: got %+v want %+v", got, want)
	}
}

func TestQtyZeroTreatedAsOne(t *testing.T) {
	got := volumetric.Calculate([]volumetric.Item{
		{Qty: 0, Length: 5, Width: 5, Height: 5},
	})
	want := volumetric.Dimensions{Length: 5, Width: 5, Height: 5}
	if got != want {
		t.Fatalf("got %+v want %+v", got, want)
	}
}
