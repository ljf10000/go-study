package libexpr

import (
	. "asdf"
	"fmt"
)

const (
	aFsmInit  aFsm = 0
	aFsmKey   aFsm = 1
	aFsmValue aFsm = 2
	aFsmKeyOp aFsm = 3
	aFsmOk    aFsm = 4
	aFsmEnd   aFsm = 5
)

var exprAtomicFsms = [aFsmEnd]string{
	aFsmInit:  "afsm(init)",
	aFsmKey:   "afsm(key)",
	aFsmValue: "afsm(value)",
	aFsmKeyOp: "afsm(key+op)",
	aFsmOk:    "afsm(key+op+value)",
}

type aFsm int

func (me aFsm) IsGood() bool {
	return me >= 0 && me < aFsmEnd
}

func (me aFsm) String() string {
	if me.IsGood() {
		return exprAtomicFsms[me]
	} else {
		return Unknow
	}
}

func newAtomic() *Atomic {
	return &Atomic{}
}

type Atomic struct {
	Op  Op
	K   *Keyword
	V   string
	fsm aFsm
}

func (me *Atomic) TypeString() string {
	return "Atomic"
}

func (me *Atomic) String() string {
	return fmt.Sprintf("{k:%s,op:%s,v:%s}", me.K, me.Op, me.V)
}

func (me *Atomic) setFsm(fsm aFsm) aFsm {
	old := me.fsm
	me.fsm = fsm

	Log.Info("%s ==> %s", old, fsm)

	return old
}

func (me *Atomic) toDeft() *Atomic {
	if aFsmValue != me.fsm {
		Panic("cannot convert fsm[%s] to deft", me.fsm)
	}

	me.K = deftKeyword()
	me.Op = OpInclude
	me.setFsm(aFsmOk)

	return me
}
