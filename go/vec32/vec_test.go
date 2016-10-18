package vec32

import (
	"math"
	"testing"
)

const (
	e = MISSING_DATA_SENTINEL
)

func near(a, b float32) bool {
	return math.Abs(float64(a-b)) < 0.001
}

func vecNear(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if !near(x, b[i]) {
			return false
		}
	}
	return true
}

func TestNorm(t *testing.T) {
	testCases := []struct {
		In  []float32
		Out []float32
	}{
		{
			In:  []float32{1.0, -1.0, e},
			Out: []float32{1.0, -1.0, e},
		},
		{
			In:  []float32{e, 2.0, -2.0},
			Out: []float32{e, 1.0, -1.0},
		},
		{
			In:  []float32{e},
			Out: []float32{e},
		},
		{
			In:  []float32{},
			Out: []float32{},
		},
		{
			In:  []float32{2.0},
			Out: []float32{0.0},
		},
		{
			In:  []float32{0.0, 0.1},
			Out: []float32{-0.05, 0.05},
		},
	}
	for _, tc := range testCases {
		Norm(tc.In, 0.1)
		if got, want := tc.In, tc.Out; !vecNear(tc.Out, tc.In) {
			t.Errorf("Norm: Got %#v Want %#v", got, want)
		}
	}
}

func TestFill(t *testing.T) {
	testCases := []struct {
		In  []float32
		Out []float32
	}{
		{
			In:  []float32{e, e, 2, 3, e, 5},
			Out: []float32{2, 2, 2, 3, 5, 5},
		},
		{
			In:  []float32{e, 3, e},
			Out: []float32{3, 3, 3},
		},
		{
			In:  []float32{e, e},
			Out: []float32{0, 0},
		},
		{
			In:  []float32{e},
			Out: []float32{0},
		},
		{
			In:  []float32{},
			Out: []float32{},
		},
	}
	for _, tc := range testCases {
		Fill(tc.In)
		if got, want := tc.In, tc.Out; !vecNear(tc.Out, tc.In) {
			t.Errorf("Fill: Got %#v Want %#v", got, want)
		}
	}

}

func TestFillAtErrors(t *testing.T) {
	testCases := []struct {
		Slice []float32
		Idx   int
	}{
		{
			Slice: []float32{e, e, 2, 3, e, 5},
			Idx:   6,
		},
		{
			Slice: []float32{},
			Idx:   0,
		},
		{
			Slice: []float32{4},
			Idx:   -1,
		},
	}
	for _, tc := range testCases {
		_, err := FillAt(tc.Slice, tc.Idx)
		if err == nil {
			t.Fatalf("Expected \"%v\" to fail FillAt.", tc)
		}
	}
}