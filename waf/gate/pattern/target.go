package pattern

type Target int

const (
	Method Target = iota
	Any
)

func (t Target) String() string {
	switch t {
	case Method:
		return "Method"
	case Any:
	default:
		return "Any"
	}
	return "Any"
}
