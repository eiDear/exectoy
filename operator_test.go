package distsqlrun

import (
	"math/rand"
	"testing"
)

func BenchmarkFilterOperator(b *testing.B) {
	var source repeatableBatchSource
	source.numOutputCols = 4
	source.Init()

	seed := int64(12345)
	rngesus := rand.New(rand.NewSource(seed))

	for i := 0; i < source.numOutputCols*batchRowLen; i++ {
		source.internalBatch[i] = rngesus.Int() % 128
	}

	var fop filterOperator
	fop.input = &source
	fop.Init()

	b.SetBytes(int64(8 * batchRowLen * source.numOutputCols))

	for i := 0; i < b.N; i++ {
		fop.Next()
	}
}
