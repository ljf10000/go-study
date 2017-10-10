package libexpr

const (
	StateNormal State = 0
	StateString State = 1
	StateEnd    State = 2
)

type State int
