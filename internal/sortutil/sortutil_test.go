package sortutil_test

import (
	"reflect"
	"testing"

	"github.com/runningmaster/sc/internal/sortutil"
)

func TestUnionInt64(t *testing.T) {
	var (
		tdt = []struct {
			in  []int64
			out []int64
		}{
			{[]int64{0, 3, 1, 2}, []int64{0, 1, 2, 3}},
		}

		out []int64
	)

	for _, tt := range tdt {
		out = sortutil.SortInt64(tt.in)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("got %v, want %v", out, tt.out)
		}
	}
}

func TestDeDupInt64(t *testing.T) {
	var (
		tdt = []struct {
			in  []int64
			out []int64
		}{
			{[]int64{1, 1, 1, 2, 2, 3, 3, 4, 4}, []int64{1, 2, 3, 4}},
		}

		out []int64
	)

	for _, tt := range tdt {
		out = sortutil.DeDupInt64(tt.in)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("got %v, want %v", out, tt.out)
		}
	}
}
