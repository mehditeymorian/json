package json

type operation int

const (
	operationDecode operation = iota
	operationEncode
	operationAll

	decodeSign = "in"
	encodeSign = "out"
)

func (o operation) IsValid(options tagOptions) bool {
	hasIn := options.Contains(decodeSign)
	hasOut := options.Contains(encodeSign)

	if (hasIn && hasOut) || (!hasIn && !hasOut) {
		return true
	}

	return (o == operationDecode && hasIn) || (o == operationEncode && hasOut)
}
