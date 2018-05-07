// Code generated by execgen; DO NOT EDIT.

package exectoy

type projPlusIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projPlusIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(intColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] + p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] + p.constArg
		}
	}
	return flow
}

func (p projPlusIntIntConstOp) Init() {}

type projPlusIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projPlusIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(intColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] + col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] + col2[i]
		}
	}
	return flow
}

func (p projPlusIntIntOp) Init() {}

type projPlusDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projPlusDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(float64Column)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] + p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] + p.constArg
		}
	}
	return flow
}

func (p projPlusDoubleDoubleConstOp) Init() {}

type projPlusDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projPlusDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(float64Column)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] + col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] + col2[i]
		}
	}
	return flow
}

func (p projPlusDoubleDoubleOp) Init() {}

type projMinusIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projMinusIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(intColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] - p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] - p.constArg
		}
	}
	return flow
}

func (p projMinusIntIntConstOp) Init() {}

type projMinusIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projMinusIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(intColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] - col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] - col2[i]
		}
	}
	return flow
}

func (p projMinusIntIntOp) Init() {}

type projMinusDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projMinusDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(float64Column)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] - p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] - p.constArg
		}
	}
	return flow
}

func (p projMinusDoubleDoubleConstOp) Init() {}

type projMinusDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projMinusDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(float64Column)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] - col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] - col2[i]
		}
	}
	return flow
}

func (p projMinusDoubleDoubleOp) Init() {}

type projMulIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projMulIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(intColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] - p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] - p.constArg
		}
	}
	return flow
}

func (p projMulIntIntConstOp) Init() {}

type projMulIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projMulIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(intColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] - col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] - col2[i]
		}
	}
	return flow
}

func (p projMulIntIntOp) Init() {}

type projMulDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projMulDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(float64Column)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] - p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] - p.constArg
		}
	}
	return flow
}

func (p projMulDoubleDoubleConstOp) Init() {}

type projMulDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projMulDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(float64Column)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] - col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] - col2[i]
		}
	}
	return flow
}

func (p projMulDoubleDoubleOp) Init() {}

type projIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(intColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] / p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] / p.constArg
		}
	}
	return flow
}

func (p projIntIntConstOp) Init() {}

type projIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(intColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] / col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] / col2[i]
		}
	}
	return flow
}

func (p projIntIntOp) Init() {}

type projDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(float64Column)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] / p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] / p.constArg
		}
	}
	return flow
}

func (p projDoubleDoubleConstOp) Init() {}

type projDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(float64Column)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] / col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] / col2[i]
		}
	}
	return flow
}

func (p projDoubleDoubleOp) Init() {}

type projEQIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projEQIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] == p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] == p.constArg
		}
	}
	return flow
}

func (p projEQIntIntConstOp) Init() {}

type projEQIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projEQIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] == col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] == col2[i]
		}
	}
	return flow
}

func (p projEQIntIntOp) Init() {}

type projEQDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projEQDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] == p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] == p.constArg
		}
	}
	return flow
}

func (p projEQDoubleDoubleConstOp) Init() {}

type projEQDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projEQDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] == col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] == col2[i]
		}
	}
	return flow
}

func (p projEQDoubleDoubleOp) Init() {}

type projNEIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projNEIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] != p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] != p.constArg
		}
	}
	return flow
}

func (p projNEIntIntConstOp) Init() {}

type projNEIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projNEIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] != col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] != col2[i]
		}
	}
	return flow
}

func (p projNEIntIntOp) Init() {}

type projNEDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projNEDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] != p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] != p.constArg
		}
	}
	return flow
}

func (p projNEDoubleDoubleConstOp) Init() {}

type projNEDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projNEDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] != col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] != col2[i]
		}
	}
	return flow
}

func (p projNEDoubleDoubleOp) Init() {}

type projLTIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projLTIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] < p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] < p.constArg
		}
	}
	return flow
}

func (p projLTIntIntConstOp) Init() {}

type projLTIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projLTIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] < col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] < col2[i]
		}
	}
	return flow
}

func (p projLTIntIntOp) Init() {}

type projLTDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projLTDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] < p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] < p.constArg
		}
	}
	return flow
}

func (p projLTDoubleDoubleConstOp) Init() {}

type projLTDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projLTDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] < col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] < col2[i]
		}
	}
	return flow
}

func (p projLTDoubleDoubleOp) Init() {}

type projLTEIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projLTEIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] <= p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] <= p.constArg
		}
	}
	return flow
}

func (p projLTEIntIntConstOp) Init() {}

type projLTEIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projLTEIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] <= col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] <= col2[i]
		}
	}
	return flow
}

func (p projLTEIntIntOp) Init() {}

type projLTEDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projLTEDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] <= p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] <= p.constArg
		}
	}
	return flow
}

func (p projLTEDoubleDoubleConstOp) Init() {}

type projLTEDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projLTEDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] <= col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] <= col2[i]
		}
	}
	return flow
}

func (p projLTEDoubleDoubleOp) Init() {}

type projGTIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projGTIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] > p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] > p.constArg
		}
	}
	return flow
}

func (p projGTIntIntConstOp) Init() {}

type projGTIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projGTIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] > col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] > col2[i]
		}
	}
	return flow
}

func (p projGTIntIntOp) Init() {}

type projGTDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projGTDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] > p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] > p.constArg
		}
	}
	return flow
}

func (p projGTDoubleDoubleConstOp) Init() {}

type projGTDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projGTDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] > col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] > col2[i]
		}
	}
	return flow
}

func (p projGTDoubleDoubleOp) Init() {}

type projGTEIntIntConstOp struct {
	input ExecOp

	colIdx   int
	constArg int

	outputIdx int
}

func (p *projGTEIntIntConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] >= p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] >= p.constArg
		}
	}
	return flow
}

func (p projGTEIntIntConstOp) Init() {}

type projGTEIntIntOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projGTEIntIntOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(intColumn)
	col2 := flow.b[p.col2Idx].(intColumn)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] >= col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] >= col2[i]
		}
	}
	return flow
}

func (p projGTEIntIntOp) Init() {}

type projGTEDoubleDoubleConstOp struct {
	input ExecOp

	colIdx   int
	constArg float64

	outputIdx int
}

func (p *projGTEDoubleDoubleConstOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col := flow.b[p.colIdx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col[i] >= p.constArg
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col[i] >= p.constArg
		}
	}
	return flow
}

func (p projGTEDoubleDoubleConstOp) Init() {}

type projGTEDoubleDoubleOp struct {
	input ExecOp

	col1Idx int
	col2Idx int

	outputIdx int
}

func (p *projGTEDoubleDoubleOp) Next() dataFlow {
	flow := p.input.Next()

	projCol := flow.b[p.outputIdx].(boolColumn)
	col1 := flow.b[p.col1Idx].(float64Column)
	col2 := flow.b[p.col2Idx].(float64Column)
	n := flow.n
	if flow.useSel {
		for s := 0; s < n; s++ {
			i := flow.sel[s]
			projCol[i] = col1[i] >= col2[i]
		}
	} else {
		for i := 0; i < n; i++ {
			projCol[i] = col1[i] >= col2[i]
		}
	}
	return flow
}

func (p projGTEDoubleDoubleOp) Init() {}
