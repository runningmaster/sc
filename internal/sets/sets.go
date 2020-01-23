package sets

import (
	"math/rand"
)

// UnionInt64 finds the union of all the given sets.
func UnionInt64(args ...[]int64) []int64 {
	if len(args) == 0 {
		return nil
	}

	res := args[0]
	tail := args[1:]

	if len(tail) == 0 {
		return unionInt64(res, nil)
	}

	for i := range tail {
		res = unionInt64(res, tail[i])
	}

	return res
}

func unionInt64(a, b []int64) []int64 {
	if len(a) > len(b) {
		a, b = b, a
	}

	tmp := make(map[int64]struct{}, len(b))
	for i := range b {
		tmp[b[i]] = struct{}{}
	}

	for i := range a {
		tmp[a[i]] = struct{}{}
	}

	res := make([]int64, 0, len(tmp))
	for k := range tmp {
		res = append(res, k)
	}

	return res
}

// UnionInt64 finds the union of all the given sets.
// The slices must be sorted in ascending order.
func UnionInt64Sorted(args ...[]int64) []int64 {
	if len(args) == 0 {
		return nil
	}

	res := args[0]
	tail := args[1:]

	if len(tail) == 0 {
		return unionInt64Sorted(res, nil)
	}

	for i := range tail {
		res = unionInt64Sorted(res, tail[i])
	}

	return res
}

func unionInt64Sorted(a, b []int64) []int64 {
	if len(a) > len(b) {
		a, b = b, a
	}

	res := make([]int64, 0, len(b)+len(a)/2)

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch {
		case a[i] < b[j]:
			res = append(res, a[i])
			i++
		case a[i] > b[j]:
			res = append(res, b[j])
			j++
		default:
			res = append(res, a[i])
			i++
			j++
		}
	}

	if i < len(a) {
		a = a[i:]
	} else if j < len(b) {
		a = b[j:]
	}

	for i := range a {
		res = append(res, a[i])
	}

	return res
}

// InterInt64 finds the intersection of all the given sets.
func InterInt64(args ...[]int64) []int64 {
	if len(args) == 0 {
		return nil
	}

	res := args[0]
	tail := args[1:]

	if len(tail) == 0 {
		return nil
	}

	for i := range tail {
		res = interInt64(res, tail[i])
	}

	return res
}

func interInt64(a, b []int64) []int64 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	if len(a) > len(b) {
		a, b = b, a
	}

	tmp := make(map[int64]struct{}, len(a))
	for i := range a {
		tmp[a[i]] = struct{}{}
	}

	res := make([]int64, 0, len(a))

	for i := range b {
		if _, ok := tmp[b[i]]; !ok {
			continue
		}

		res = append(res, b[i])
	}

	return res
}

// InterInt64 finds the intersection of all the given sets.
// The slices must be sorted in ascending order.
func InterInt64Sorted(args ...[]int64) []int64 {
	if len(args) == 0 {
		return nil
	}

	res := args[0]
	tail := args[1:]

	if len(tail) == 0 {
		return nil
	}

	for i := range tail {
		res = interInt64Sorted(res, tail[i])
	}

	return res
}

func interInt64Sorted(a, b []int64) []int64 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	if len(a) > len(b) {
		a, b = b, a
	}

	res := make([]int64, 0, len(a))

	for i, j := 0, 0; i < len(a) && j < len(b); {
		switch {
		case a[i] < b[j]:
			i++
		case a[i] > b[j]:
			j++
		default:
			res = append(res, a[i])
			i++
			j++
		}
	}

	return res
}

// DiffInt64 finds the difference between the first set and all the rest ones.
func DiffInt64(args ...[]int64) []int64 {
	if len(args) == 0 {
		return nil
	}

	res := args[0]
	tail := args[1:]

	for i := range tail {
		res = diffInt64(res, tail[i])
	}

	return res
}

func diffInt64(a, b []int64) []int64 {
	if len(a) == 0 {
		return nil
	}

	if len(b) == 0 {
		return a
	}

	tmp := make(map[int64]struct{}, len(b))
	for i := range b {
		tmp[b[i]] = struct{}{}
	}

	res := make([]int64, 0, len(a))

	for i := range a {
		if _, ok := tmp[a[i]]; ok {
			continue
		}

		res = append(res, a[i])
	}

	return res
}

// DiffInt64 finds the difference between the first set and all the rest ones.
// The slices must be sorted in ascending order.
func DiffInt64Sorted(args ...[]int64) []int64 {
	if len(args) == 0 {
		return nil
	}

	res := args[0]
	tail := args[1:]

	for i := range tail {
		res = diffInt64Sorted(res, tail[i])
	}

	return res
}

func diffInt64Sorted(a, b []int64) []int64 {
	if len(a) == 0 {
		return nil
	}

	if len(b) == 0 || a[len(a)-1] < b[0] || a[0] > b[len(b)-1] {
		return a
	}

	res := make([]int64, 0, len(a))

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch {
		case a[i] < b[j]:
			res = append(res, a[i])
			i++
		case a[i] > b[j]:
			j++
		default:
			i++
			j++
		}
	}

	if i < len(a) {
		a = a[i:]
		for i := range a {
			res = append(res, a[i])
		}
	}

	return res
}

// TestData makes unordered sets from a to z with len (l) and random max values.
func TestData(l, max int64) map[string][]int64 {
	d := make(map[string][]int64, 'z'-'a'+1)
	r := rand.New(rand.NewSource(99))

	var x, y []int64

	for i := 'a'; i <= 'z'; i++ {
		for j := 0; j <= int(r.Int63n(l)); j++ {
			x = append(x, r.Int63n(max))
		}

		if len(y) > 0 {
			x = append(x, y[len(y)/2:]...)
		}

		y = x
		d[string(i)] = x
	}

	return d
}
