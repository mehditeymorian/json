package json

type operation int

const (
	operationDecode operation = iota
	operationEncode
	operationAll

	decodeSign = "out"
	encodeSign = "in"
)

func (o operation) IsValid(options tagOptions) bool {
	hasIn := options.Contains(encodeSign)
	hasOut := options.Contains(decodeSign)

	if (hasIn && hasOut) || (!hasIn && !hasOut) {
		return true
	}

	if o == operationEncode {
		return hasIn
	}

	if o == operationDecode {
		return hasOut
	}

	return true
}
