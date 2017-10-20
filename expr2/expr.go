package libexpr

import (
	. "asdf"
	"fmt"
)

const (
	eFsmInit          eFsm = 0
	eFsmAtomic        eFsm = 1
	eFsmSingle        eFsm = 2
	eFsmSingleOk      eFsm = 3
	eFsmExpr          eFsm = 4
	eFsmExprMulti     eFsm = 5
	eFsmExprOk        eFsm = 6
	eFsmExprMultiMore eFsm = 7
	eFsmEnd           eFsm = 8
)

var exprExprFsms = [eFsmEnd]string{
	eFsmInit:          "efsm(init)",
	eFsmAtomic:        "efsm(atomic)",
	eFsmSingle:        "efsm(single)",
	eFsmSingleOk:      "efsm(single+expr)",
	eFsmExpr:          "efsm(expr)",
	eFsmExprMulti:     "efsm(expr+multi)",
	eFsmExprOk:        "efsm(expr+multi+expr)",
	eFsmExprMultiMore: "efsm(expr+multi+expr+multi)",
}

type eFsm int

func (me eFsm) IsGood() bool {
	return me >= 0 && me < eFsmEnd
}

func (me eFsm) String() string {
	if me.IsGood() {
		return exprExprFsms[me]
	} else {
		return Unknow
	}
}

func newExpr(fsm eFsm) *Expr {
	Log.Info("create empty expr")
	return &Expr{
		fsm:      fsm,
		Children: []*Expr{},
	}
}

func newAtomicExpr(atomic *Atomic) *Expr {
	Log.Info("create atomic expr:%s", atomic)
	return &Expr{
		fsm:      eFsmAtomic,
		Atomic:   atomic,
		Children: []*Expr{},
	}
}

func newExprExpr(expr *Expr) *Expr {
	Log.Info("create expr expr")
	return &Expr{
		fsm:      eFsmExpr,
		Children: []*Expr{expr},
	}
}

func newExprSingleOk(expr *Expr) *Expr {
	Log.Info("create single expr")
	return &Expr{
		fsm:      eFsmSingleOk,
		Logic:    LogicNot,
		Children: []*Expr{expr},
	}
}

func newExprMultiOk(logic Logic, a, b *Expr) *Expr {
	return &Expr{
		fsm:      eFsmExprOk,
		Logic:    logic,
		Children: []*Expr{a, b},
	}
}

func newAndExpr(a, b *Expr) *Expr {
	Log.Info("create and expr")
	return newExprMultiOk(LogicAnd, a, b)
}

func newOrExpr(a, b *Expr) *Expr {
	Log.Info("create or expr")
	return newExprMultiOk(LogicOr, a, b)
}

type Expr struct {
	fsm      eFsm
	Logic    Logic
	Atomic   *Atomic
	Children []*Expr
}

func (me *Expr) EString() string {
	s := "{"

	count := len(me.Children)
	for idx, v := range me.Children {
		s += v.String()
		if idx < count-1 {
			s += ", "
		}
	}

	s += "}"

	return s
}

func (me *Expr) String() string {
	if me.IsAtomic() {
		return me.Atomic.String()
	} else {
		s := fmt.Sprintf("{op: %s, ", me.Logic)

		count := len(me.Children)
		for idx, v := range me.Children {
			s += fmt.Sprintf("%d: %s", idx, v.String())
			if idx < count-1 {
				s += ", "
			}
		}

		s += "}"

		return s
	}
}

func (me *Expr) TypeString() string {
	return "Expr"
}

func (me *Expr) setFsm(fsm eFsm) eFsm {
	old := me.fsm
	me.fsm = fsm

	Log.Info("%s ==> %s", old, fsm)

	return old
}

func (me *Expr) pushExpr(expr *Expr) {
	Log.Info("push expr %d==>%d", len(me.Children), len(me.Children)+1)

	me.Children = append(me.Children, expr)
}

func (me *Expr) IsAtomic() bool {
	return nil != me.Atomic
}
