package calc

import (
	"fmt"

	"github.com/runningmaster/sc/internal/parser"
	"github.com/runningmaster/sc/internal/sets"
	"github.com/runningmaster/sc/internal/sortutil"
)

func Execute(cmd string) ([]int64, error) {
	ast, err := parser.Parse(cmd)
	if err != nil {
		return nil, err
	}

	var stack []*parser.Node

	err = ast.Apply(func(n *parser.Node, err error) error {
		if err != nil {
			return err
		}

		stack = append(stack, n)
		return nil
	})
	if err != nil {
		return nil, err
	}

	data := sets.TestData(100, 100)

	var acc []int64
	for i := 0; i < len(stack); i++ {
		acc, err = processCommand(stack[i], data, acc...)
		if err != nil {
			return nil, err
		}
	}

	return acc, nil
}

func processCommand(n *parser.Node, data map[string][]int64, accum ...int64) ([]int64, error) {
	var vals [][]int64

	for i := range n.Vals() {
		if v, ok := data[string(rune(n.Vals()[i][0]))]; ok {
			vals = append(vals, sortutil.DeDupInt64(sortutil.SortInt64(v)))
		} else {
			return nil, fmt.Errorf("data not found for %v", n.Vals()[i][0])
		}
	}

	if len(accum) > 0 {
		vals = append(vals, accum)
	}

	var out []int64

	switch n.Type() {
	case parser.TokenSUM:
		out = sets.UnionInt64Sorted(vals...)
	case parser.TokenINT:
		out = sets.InterInt64Sorted(vals...)
	case parser.TokenDIF:
		out = sets.DiffInt64Sorted(vals...)
	default:
		return nil, fmt.Errorf("unknown command %v", n.Type())
	}

	return out, nil
}
