package sortutil

import (
	"sort"
)

// SortInt64 sorts slice of int64 values.
func SortInt64(v []int64) []int64 {
	sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
	return v
}

// DeDupInt64 deduplicates slice of int64 values.
// The slices must be sorted in ascending order.
func DeDupInt64(v []int64) []int64 {
	var j int

	for i := 1; i < len(v); i++ {
		if v[j] == v[i] {
			continue
		}
		j++

		v[j] = v[i]
	}

	return v[:j+1]
}
