package common

const (
	// OK ..
	OK = 0
	// [1, 1000] preserved range

	// [1001, 2000] for cli handler and cli module

	// ErrNoOpFound no operator match
	ErrNoOpFound = 1001
	ErrNewOp     = 1002
	ErrOpExec    = 1003
	ErrTimeout   = 1005
)
