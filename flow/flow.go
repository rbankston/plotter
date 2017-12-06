package flow

// flow is an interface that describes processes for getting work done. An
// object that implements this interface must provide a private perform()
// method.
type Flow interface {
	perform() error
}

// Perform, given a workflow, will perform the work defined by said workflow.
func Perform(f Flow) error {
	return f.perform()
}