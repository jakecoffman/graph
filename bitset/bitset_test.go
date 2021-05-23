package bitset

import (
	"reflect"
	"testing"
)

func TestBitset(t *testing.T) {
	var bits uint64

	Set(&bits, 7)
	Set(&bits, 28)

	var indices []int
	EachSetBit(bits, func(i int) {
		indices = append(indices, i)
	})
	if !reflect.DeepEqual(indices, []int{7, 28}) {
		t.Error(indices)
	}

	if Count(bits) != 2 {
		t.Error(Count(bits))
	}

	if !IsSet(bits, 7) {
		t.Error(IsSet(bits, 7))
	}
	if IsSet(bits, 1) {
		t.Error(IsSet(bits, 1))
	}

	Clear(&bits, 7)
	if IsSet(bits, 7) {
		t.Error(IsSet(bits, 7))
	}
	indices = indices[0:0]
	EachSetBit(bits, func(i int) {
		indices = append(indices, i)
	})
	if !reflect.DeepEqual(indices, []int{28}) {
		t.Error(indices)
	}
}
