package exectoy

type sortedDistinctOperator struct {
	input ExecOp

	sortedDistinctCols []int
	cols               []column

	lastVal tuple

	outputVec []bool

	numOutputCols int
	internalBatch batch
}

func (p *sortedDistinctOperator) Init() {
	p.internalBatch = make(batch, p.numOutputCols)
	p.cols = make([]column, len(p.sortedDistinctCols))
	p.lastVal = make(tuple, len(p.sortedDistinctCols))
	p.outputVec = make([]bool, batchRowLen)
}

func (p *sortedDistinctOperator) Next() dataFlow {
	// outputBitmap contains row indexes that we will output
	flow := p.input.Next()
	for i, c := range p.sortedDistinctCols {
		p.cols[i] = flow.b[c]
	}
	// TODO(jordan) p.lastVal is wrong in the very first invocation.

	if flow.useSel {
		for cIdx, col := range p.cols {
			lastVal := p.lastVal[cIdx]
			for s := 0; s < flow.n; s++ {
				i := flow.sel[s]
				if col[i] != lastVal {
					p.outputVec[i] = true
					lastVal = col[i]
				}
			}
			p.lastVal[cIdx] = lastVal
		}
	} else {
		for cIdx, col := range p.cols {
			lastVal := p.lastVal[cIdx]
			for i := 0; i < flow.n; i++ {
				if col[i] != lastVal {
					p.outputVec[i] = true
					lastVal = col[i]
				}
			}
			p.lastVal[cIdx] = lastVal
		}
	}
	// convert outputVec to sel

	idx := 0
	if flow.useSel {
		max := flow.sel[flow.n-1]
		for i := 0; i < max; i++ {
			if p.outputVec[i] {
				flow.sel[idx] = i
				idx++
			}
		}
	} else {
		for i := 0; i < flow.n; i++ {
			if p.outputVec[i] {
				flow.sel[idx] = i
				idx++
			}
		}
	}

	flow.useSel = true
	flow.n = idx
	return flow
}
