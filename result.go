package floc

//go:generate stringer -type=Result

// Result is the result of control flow execution.
type Result int

// Possible results.
const (
	None Result = iota
	Completed
	Canceled
)

// IsNone tests if the resilt is None.
func (result Result) IsNone() bool {
	return result == None
}

// IsCompleted tests if the resilt is Completed.
func (result Result) IsCompleted() bool {
	return result == Completed
}

// IsCanceled tests if the resilt is Canceled.
func (result Result) IsCanceled() bool {
	return result == Canceled
}
