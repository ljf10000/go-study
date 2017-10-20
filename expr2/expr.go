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
		children: []*Expr{},
	}
}

func newAtomicExpr(atomic *Atomic) *Expr {
	Log.Info("create atomic expr:%s", atomic)
	return &Expr{
		fsm:      eFsmAtomic,
		atomic:   atomic,
		children: []*Expr{},
	}
}

func newExprExpr(expr *Expr) *Expr {
	Log.Info("create expr expr")
	return &Expr{
		fsm:      eFsmExpr,
		children: []*Expr{expr},
	}
}

func newExprSingleOk(expr *Expr) *Expr {
	Log.Info("create single expr")
	return &Expr{
		fsm:      eFsmSingleOk,
		logic:    LogicNot,
		children: []*Expr{expr},
	}
}

func newExprMultiOk(logic Logic, a, b *Expr) *Expr {
	return &Expr{
		fsm:      eFsmExprOk,
		logic:    logic,
		children: []*Expr{a, b},
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
	logic    Buildin
	atomic   *Atomic
	children []*Expr
}

func (me *Expr) EString() string {
	s := "{"

	count := len(me.children)
	for idx, v := range me.children {
		s += v.String()
		if idx < count-1 {
			s += ", "
		}
	}

	s += "}"

	return s
}

func (me *Expr) String() string {
	if me.isAtomic() {
		return me.atomic.String()
	} else {
		s := fmt.Sprintf("{op: %s, ", me.logic)

		count := len(me.children)
		for idx, v := range me.children {
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

func (me *Expr) LString(level int) string {
	s := "{" + Crlf

	if me.isAtomic() {
		s += TabN(level+1) + me.atomic.LString(level+1) + "," + Crlf
	} else {
		s += TabN(level+1) + fmt.Sprintf("op: %s,", me.logic) + Crlf

		for idx, v := range me.children {
			s += TabN(level+1) + fmt.Sprintf("%d: %s,", idx, me.TypeString()+v.LString(level+1)) + Crlf
		}
	}

	s += TabN(level) + "}"

	return s
}

func (me *Expr) setFsm(fsm eFsm) eFsm {
	old := me.fsm
	me.fsm = fsm

	Log.Info("%s ==> %s", old, fsm)

	return old
}

func (me *Expr) isAtomic() bool {
	return nil != me.atomic
}

func (me *Expr) pushExpr(expr *Expr) {
	Log.Info("push expr %d==>%d", len(me.children), len(me.children)+1)

	me.children = append(me.children, expr)
}
