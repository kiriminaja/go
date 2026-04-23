// Package volumetric computes the smallest bounding box that fits a set of
// items, evaluating three simple stacking strategies (vertical, horizontal,
// and side-by-side) and returning the one with the smallest volume.
package volumetric

// Item describes a single product to be stacked. Length, width and height
// are expressed in the same unit (typically centimetres). Qty must be at
// least 1.
type Item struct {
	Qty    int
	Length int
	Width  int
	Height int
}

// Dimensions is the resulting bounding box.
type Dimensions struct {
	Length int
	Width  int
	Height int
}

// Calculate returns the smallest bounding box (length × width × height)
// across three stacking strategies. An empty slice yields a zero-value
// Dimensions.
func Calculate(items []Item) Dimensions {
	if len(items) == 0 {
		return Dimensions{}
	}

	var (
		hVert, lVert, wVert int
		hHor, lHor, wHor    int
		hSide, lSide, wSide int
	)

	for _, it := range items {
		qty := it.Qty
		if qty < 1 {
			qty = 1
		}

		// Vertical: stack height, max footprint
		hVert += it.Height * qty
		lVert = max(lVert, it.Length)
		wVert = max(wVert, it.Width)

		// Horizontal: stack length, max cross-section
		hHor = max(hHor, it.Height)
		lHor += it.Length * qty
		wHor = max(wHor, it.Width)

		// Side-by-side: stack width, max cross-section
		hSide = max(hSide, it.Height)
		lSide = max(lSide, it.Length)
		wSide += it.Width * qty
	}

	volVert := hVert * lVert * wVert
	volHor := hHor * lHor * wHor
	volSide := hSide * lSide * wSide

	switch {
	case volVert <= volHor && volVert <= volSide:
		return Dimensions{Length: lVert, Width: wVert, Height: hVert}
	case volHor <= volSide:
		return Dimensions{Length: lHor, Width: wHor, Height: hHor}
	default:
		return Dimensions{Length: lSide, Width: wSide, Height: hSide}
	}
}
