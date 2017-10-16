package libexpr

type Atomic struct {
	op    Buildin
	obj   *Object
	value string
	level int
}

func newAtomic(op Buildin, k, v string, level int) *Atomic {
	return &Atomic{
		op:    op,
		obj:   newObject(k),
		value: v,
		level: level,
	}
}

func emptyAtomic(level int) *Atomic {
	return &Atomic{
		op:    BuildinUnknow,
		level: level,
	}
}
