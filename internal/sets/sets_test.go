package sets_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/runningmaster/sc/internal/sets"
	"github.com/runningmaster/sc/internal/sortutil"
)

var ttUnion = []struct { //nolint: gochecknoglobals
	in  [][]int64
	out []int64
}{
	{[][]int64{
		{0, 1, 2, 3},
	}, []int64{0, 1, 2, 3}},
	{[][]int64{
		{0, 1, 2, 3},
		{4, 5},
	}, []int64{0, 1, 2, 3, 4, 5}},
	{[][]int64{
		{0, 1, 2, 3},
		{2, 3, 4, 5},
	}, []int64{0, 1, 2, 3, 4, 5}},
	{[][]int64{
		{0, 1},
		{2, 3, 4, 5},
	}, []int64{0, 1, 2, 3, 4, 5}},
	{[][]int64{
		{0, 1, 2, 3},
		{4, 5},
	}, []int64{0, 1, 2, 3, 4, 5}},
	{[][]int64{
		{0, 1, 2, 3},
		{2, 3, 4, 5},
		{7},
	}, []int64{0, 1, 2, 3, 4, 5, 7}},
	{[][]int64{
		{7},
		{0, 1, 2, 3},
		{2, 3, 4, 5},
	}, []int64{0, 1, 2, 3, 4, 5, 7}},
	{[][]int64{
		{0, 1, 2, 3},
		{7},
		{2, 3, 4, 5},
	}, []int64{0, 1, 2, 3, 4, 5, 7}},
	{[][]int64{
		{0},
	}, []int64{0}},
	{[][]int64{
		{},
	}, []int64{}},
	{[][]int64{
		nil,
	}, []int64{}},
}

func TestUnionInt64(t *testing.T) {
	var out []int64
	for i, tt := range ttUnion {
		out = sortutil.SortInt64(sets.UnionInt64(tt.in...))
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("pos %v: got %v, want %v", i, out, tt.out)
		}
	}
}

func TestUnionInt64Sorted(t *testing.T) {
	var out []int64
	for i, tt := range ttUnion {
		out = sets.UnionInt64Sorted(tt.in...)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("pos %v: got %v, want %v", i, out, tt.out)
		}
	}
}

var ttInter = []struct { //nolint: gochecknoglobals
	in  [][]int64
	out []int64
}{
	{[][]int64{
		{0, 1, 2, 3},
		{2, 3, 4, 5},
	}, []int64{2, 3}},
	{[][]int64{
		{2, 3, 4, 5},
		{0, 1, 2, 3},
	}, []int64{2, 3}},
	{[][]int64{
		{0, 1, 2, 3},
		{2, 3, 4, 5},
	}, []int64{2, 3}},
	{[][]int64{
		{0, 1, 2},
		{2, 3, 4, 5},
	}, []int64{2}},
	{[][]int64{
		{0, 1, 3},
		{2, 3},
	}, []int64{3}},
	{[][]int64{
		{2, 3, 4},
		{0, 1, 2, 3},
		{2, 3, 4, 5},
	}, []int64{2, 3}},
	{[][]int64{
		{0, 1, 2, 3},
		{2, 3, 4, 5},
		{2, 3, 5},
	}, []int64{2, 3}},
	{[][]int64{
		{0, 1, 2, 3},
		{2, 3, 5},
		{2, 3, 4, 5},
	}, []int64{2, 3}},
	{[][]int64{
		{0},
	}, nil},
	//			{[][]int64{
	//				{},
	//			}, []int64{}},
}

func TestInterInt64Sorted(t *testing.T) {
	var out []int64
	for i, tt := range ttInter {
		out = sets.InterInt64Sorted(tt.in...)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("pos %v: got %v, want %v", i, out, tt.out)
		}
	}
}

func TestInterInt64(t *testing.T) {
	var out []int64
	for i, tt := range ttInter {
		out = sets.InterInt64(tt.in...)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("pos %v: got %v, want %v", i, out, tt.out)
		}
	}
}

var ttDiff = []struct { //nolint: gochecknoglobals
	in  [][]int64
	out []int64
}{
	{[][]int64{
		{0},
		{1},
	}, []int64{0}},
	{[][]int64{
		{1},
		{0},
	}, []int64{1}},
	{[][]int64{
		{0, 1, 2},
		{2, 3, 4, 5},
	}, []int64{0, 1}},
	{[][]int64{
		{0, 1, 2, 3},
		{2, 3, 4, 5},
	}, []int64{0, 1}},
	{[][]int64{
		{2, 3, 4, 5},
		{4, 5, 6, 7},
	}, []int64{2, 3}},
	{[][]int64{
		{2, 3, 4, 5},
		{0, 1, 2, 3},
	}, []int64{4, 5}},
	{[][]int64{
		{4, 5},
		{0, 1, 2, 3},
	}, []int64{4, 5}},
	{[][]int64{
		{2, 3, 4, 5},
		{2, 3},
	}, []int64{4, 5}},
	{[][]int64{
		{2, 3, 4, 5},
		{2, 3},
		{6, 7, 9},
	}, []int64{4, 5}},
	{[][]int64{
		{2, 3, 4, 5},
	}, []int64{2, 3, 4, 5}},
	//{[][]int64{
	//	{0},
	//}, []int64{0}},
	//{[][]int64{
	//	{},
	//}, []int64{}},
}

func TestDiffInt64Sorted(t *testing.T) {
	var out []int64
	for i, tt := range ttDiff {
		out = sets.DiffInt64Sorted(tt.in...)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("pos %v: got %v, want %v", i, out, tt.out)
		}
	}
}

func TestDiffInt64(t *testing.T) {
	var out []int64
	for i, tt := range ttDiff {
		out = sets.DiffInt64(tt.in...)
		if !reflect.DeepEqual(out, tt.out) {
			t.Errorf("pos %v: got %v, want %v", i, out, tt.out)
		}
	}
}

var (
	data   = ddsort(sets.TestData(1000, 1000)) //nolint: gochecknoglobals
	result []int64                             //nolint: gochecknoglobals
)

func ddsort(data map[string][]int64) map[string][]int64 {
	fmt.Println("generate benchmark data, please wait...")
	for k := range data {
		data[k] = sortutil.DeDupInt64(sortutil.SortInt64(data[k]))
	}

	return data
}

func BenchmarkUnionInt64(b *testing.B) {
	var r []int64
	for n := 0; n < b.N; n++ {
		r = sortutil.SortInt64(sets.UnionInt64(
			data["a"], data["b"], data["c"], data["d"], data["e"], data["f"]))
	}

	result = r
}

func BenchmarkUnionInt64Sorted(b *testing.B) {
	var r []int64
	for n := 0; n < b.N; n++ {
		r = sets.UnionInt64Sorted(
			data["a"], data["b"], data["c"], data["d"], data["e"], data["f"])
	}

	result = r
}

func BenchmarkInterInt64(b *testing.B) {
	var r []int64
	for n := 0; n < b.N; n++ {
		r = sets.InterInt64(
			data["a"], data["b"], data["c"], data["d"], data["e"], data["f"])
	}

	result = r
}

func BenchmarkInterInt64Sorted(b *testing.B) {
	var r []int64
	for n := 0; n < b.N; n++ {
		r = sets.InterInt64Sorted(
			data["a"], data["b"], data["c"], data["d"], data["e"], data["f"])
	}

	result = r
}

func BenchmarkDiffInt64(b *testing.B) {
	var r []int64
	for n := 0; n < b.N; n++ {
		r = sets.DiffInt64(
			data["a"], data["b"], data["c"], data["d"], data["e"], data["f"])
	}

	result = r
}

func BenchmarkDiffInt64Sorted(b *testing.B) {
	var r []int64
	for n := 0; n < b.N; n++ {
		r = sets.DiffInt64Sorted(
			data["a"], data["b"], data["c"], data["d"], data["e"], data["f"])
	}

	result = r
}
