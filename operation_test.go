package json

import "testing"

func Test_operation_IsValid(t *testing.T) {
	tests := []struct {
		name string
		o    operation
		args tagOptions
		want bool
	}{
		{
			name: "valid encode operation, having only encode option",
			o:    operationEncode,
			args: encodeSign,
			want: true,
		},
		{
			name: "invalid encode operation, having out option",
			o:    operationEncode,
			args: decodeSign,
			want: false,
		},
		{
			name: "valid decode operation, having only decode option",
			o:    operationDecode,
			args: decodeSign,
			want: true,
		},
		{
			name: "invalid decode operation, having only encode option",
			o:    operationDecode,
			args: encodeSign,
			want: false,
		},
		{
			name: "valid encode operation, having both encode and decode options",
			o:    operationEncode,
			args: encodeSign + "," + decodeSign,
			want: true,
		},
		{
			name: "valid decode operation, having both encode and decode options",
			o:    operationDecode,
			args: encodeSign + "," + decodeSign,
			want: true,
		},
		{
			name: "valid encode operation, having no particular options specified",
			o:    operationEncode,
			args: "",
			want: true,
		},
		{
			name: "valid decode operation, having no particular options specified",
			o:    operationDecode,
			args: "",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.IsValid(tt.args); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
