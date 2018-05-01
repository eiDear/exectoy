package exectoy

// These will be autogenerated.
// col1 is the column being filtered.
// c is the constant being compared against.
// res is the output selection vector.
// returned is the number of rows in the selection vector
type selectLTIntIntConstOp struct {
	input ExecOp

	col1Idx  int
	constArg int
}

func (p *selectLTIntIntConstOp) Init() {}

func (p *selectLTIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	col1 := flow.b[p.col1Idx].(intColumn)
	idx := 0
	// hard to see how to eliminate this if.
	if flow.useSel {
		for s := 0; s < flow.n; s++ {
			i := flow.sel[s]
			if col1[i] < p.constArg {
				flow.sel[idx] = i
				idx++
			}
		}
	} else {
		for i := 0; i < flow.n; i++ {
			if col1[i] < p.constArg {
				flow.sel[idx] = i
				idx++
			}
		}
	}
	flow.n = idx
	flow.useSel = true
	return flow
}

type selectLTIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int
}

func (p *selectLTIntIntOp) Init() {}

func (p *selectLTIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)

	idx := 0
	if flow.useSel {
		for s := 0; s < flow.n; s++ {
			i := flow.sel[s]
			if col1[i] < col2[i] {
				flow.sel[idx] = i
				idx++
			}
		}
	} else {
		for i := 0; i < flow.n; i++ {
			if col1[i] < col2[i] {
				flow.sel[idx] = i
				idx++
			}
		}
		flow.useSel = true
	}
	flow.n = idx
	return flow
}
