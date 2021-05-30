package benchmarks

import (
	"math/rand"
	"testing"
	"time"
)

const size = 10_000

var ints [size]int

func init() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		ints[i] = rand.Intn(500_000)
	}
}

func BenchmarkRangeIntNoCopy(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for j := range ints {
			sum += ints[j]
		}
	}
}

func BenchmarkRangeIntCopy(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for _, theInt := range ints {
			sum += theInt
		}
	}
}

func BenchmarkForInt(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(ints); j++ {
			sum += ints[j]
		}
	}
}

type MyStruct struct {
	AnInt   int
	AString string
	Bytes   [1024]byte
	Float1  float64
	Float2  float64
}

var myStructs [size]MyStruct
var myStructPtrs [size]*MyStruct

func init() {
	for i := 0; i < len(myStructs); i++ {
		myStructs[i] = MyStruct{
			AnInt:   rand.Intn(50_000),
			AString: "Hello, world!",
			Bytes:   [1024]byte{},
			Float1:  rand.Float64(),
			Float2:  rand.Float64(),
		}
		myStructPtrs[i] = &myStructs[i]
	}
}

func BenchmarkRangeStructNoCopy(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for j := range myStructs {
			sum += myStructs[j].AnInt
		}
	}
}

func BenchmarkRangeStructCopy(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for _, myStruct := range myStructs {
			sum += myStruct.AnInt
		}
	}
}

func BenchmarkForStruct(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(myStructs); j++ {
			sum += myStructs[j].AnInt
		}
	}
}

func BenchmarkRangeStructNoCopyPtr(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for j := range myStructPtrs {
			sum += myStructPtrs[j].AnInt
		}
	}
}

func BenchmarkRangeStructCopyPtr(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for _, myStruct := range myStructPtrs {
			sum += myStruct.AnInt
		}
	}
}

func BenchmarkForStructPtr(b *testing.B) {
	sum := 0
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(myStructPtrs); j++ {
			sum += myStructPtrs[j].AnInt
		}
	}
}
