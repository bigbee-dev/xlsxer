package xlsxer

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestToString(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want string
	}{
		{
			name: "String",
			in:   "test",
			want: "test",
		},
		{
			name: "Bool",
			in:   true,
			want: "true",
		},
		{
			name: "Int",
			in:   123,
			want: "123",
		},
		{
			name: "Uint",
			in:   uint(123),
			want: "123",
		},
		{
			name: "Float32",
			in:   float32(123.45),
			want: "123.45",
		},
		{
			name: "Float64",
			in:   float64(123.45),
			want: "123.45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toString(tt.in)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString_Error(t *testing.T) {
	_, err := toString([]int{1, 2, 3})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestToBool(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want bool
		err  error
	}{
		{
			name: "String - Yes",
			in:   "yes",
			want: true,
			err:  nil,
		},
		{
			name: "String - No",
			in:   "no",
			want: false,
			err:  nil,
		},
		{
			name: "String - Empty",
			in:   "",
			want: false,
			err:  nil,
		},
		{
			name: "String - True",
			in:   "true",
			want: true,
			err:  nil,
		},
		{
			name: "String - False",
			in:   "false",
			want: false,
			err:  nil,
		},
		{
			name: "String - Invalid",
			in:   "invalid",
			want: false,
			err:  fmt.Errorf("strconv.ParseBool: parsing \"invalid\": invalid syntax"),
		},
		{
			name: "Bool - True",
			in:   true,
			want: true,
			err:  nil,
		},
		{
			name: "Bool - False",
			in:   false,
			want: false,
			err:  nil,
		},
		{
			name: "Int - Non-zero",
			in:   123,
			want: true,
			err:  nil,
		},
		{
			name: "Int - Zero",
			in:   0,
			want: false,
			err:  nil,
		},
		{
			name: "Uint - Non-zero",
			in:   uint(123),
			want: true,
			err:  nil,
		},
		{
			name: "Uint - Zero",
			in:   uint(0),
			want: false,
			err:  nil,
		},
		{
			name: "Float32 - Non-zero",
			in:   float32(123.45),
			want: true,
			err:  nil,
		},
		{
			name: "Float32 - Zero",
			in:   float32(0),
			want: false,
			err:  nil,
		},
		{
			name: "Float64 - Non-zero",
			in:   float64(123.45),
			want: true,
			err:  nil,
		},
		{
			name: "Float64 - Zero",
			in:   float64(0),
			want: false,
			err:  nil,
		},
		{
			name: "Unknown Type",
			in:   []int{1, 2, 3},
			want: false,
			err:  fmt.Errorf("No known conversion from []int to bool"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBool(tt.in)
			if err != nil {
				if tt.err == nil {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != tt.err.Error() {
					t.Errorf("Unexpected error: got %v, want %v", err, tt.err)
				}
			} else if got != tt.want {
				t.Errorf("toBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want int64
		err  error
	}{
		{
			name: "String - Empty",
			in:   "",
			want: 0,
			err:  nil,
		},
		{
			name: "String - Integer",
			in:   "123",
			want: 123,
			err:  nil,
		},
		{
			name: "String - Float",
			in:   "123.45",
			want: 123,
			err:  nil,
		},
		{
			name: "Bool - True",
			in:   true,
			want: 1,
			err:  nil,
		},
		{
			name: "Bool - False",
			in:   false,
			want: 0,
			err:  nil,
		},
		{
			name: "Int - Int",
			in:   123,
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int8",
			in:   int8(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int16",
			in:   int16(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int32",
			in:   int32(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int64",
			in:   int64(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint",
			in:   uint(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint8",
			in:   uint8(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint16",
			in:   uint16(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint32",
			in:   uint32(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint64",
			in:   uint64(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Float32 - Float32",
			in:   float32(123.45),
			want: 123,
			err:  nil,
		},
		{
			name: "Float64 - Float64",
			in:   float64(123.45),
			want: 123,
			err:  nil,
		},
		{
			name: "Unknown Type",
			in:   []int{1, 2, 3},
			want: 0,
			err:  fmt.Errorf("No known conversion from []int to int"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toInt(tt.in)
			if err != nil {
				if tt.err == nil {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != tt.err.Error() {
					t.Errorf("Unexpected error: got %v, want %v", err, tt.err)
				}
			} else if got != tt.want {
				t.Errorf("toInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUint(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want uint64
		err  error
	}{
		{
			name: "String - Empty",
			in:   "",
			want: 0,
			err:  nil,
		},
		{
			name: "String - Integer",
			in:   "123",
			want: 123,
			err:  nil,
		},
		{
			name: "String - Float",
			in:   "123.45",
			want: 123,
			err:  nil,
		},
		{
			name: "Bool - True",
			in:   true,
			want: 1,
			err:  nil,
		},
		{
			name: "Bool - False",
			in:   false,
			want: 0,
			err:  nil,
		},
		{
			name: "Int - Int",
			in:   123,
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int8",
			in:   int8(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int16",
			in:   int16(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int32",
			in:   int32(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int64",
			in:   int64(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint",
			in:   uint(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint8",
			in:   uint8(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint16",
			in:   uint16(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint32",
			in:   uint32(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint64",
			in:   uint64(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Float32 - Float32",
			in:   float32(123.45),
			want: 123,
			err:  nil,
		},
		{
			name: "Float64 - Float64",
			in:   float64(123.45),
			want: 123,
			err:  nil,
		},
		{
			name: "Unknown Type",
			in:   []int{1, 2, 3},
			want: 0,
			err:  fmt.Errorf("No known conversion from []int to uint"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toUint(tt.in)
			if err != nil {
				if tt.err == nil {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != tt.err.Error() {
					t.Errorf("Unexpected error: got %v, want %v", err, tt.err)
				}
			} else if got != tt.want {
				t.Errorf("toUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat(t *testing.T) {
	float64EqualityThreshold := 0.00001

	tests := []struct {
		name string
		in   interface{}
		want float64
		err  error
	}{
		{
			name: "String - Empty",
			in:   "",
			want: 0,
			err:  nil,
		},
		{
			name: "String - Integer",
			in:   "123",
			want: 123,
			err:  nil,
		},
		{
			name: "String - Float",
			in:   "123.45",
			want: 123.45,
			err:  nil,
		},
		{
			name: "Bool - True",
			in:   true,
			want: 1,
			err:  nil,
		},
		{
			name: "Bool - False",
			in:   false,
			want: 0,
			err:  nil,
		},
		{
			name: "Int - Int",
			in:   123,
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int8",
			in:   int8(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int16",
			in:   int16(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int32",
			in:   int32(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Int - Int64",
			in:   int64(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint",
			in:   uint(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint8",
			in:   uint8(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint16",
			in:   uint16(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint32",
			in:   uint32(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Uint - Uint64",
			in:   uint64(123),
			want: 123,
			err:  nil,
		},
		{
			name: "Float32 - Float32",
			in:   float32(123.46),
			want: 123.46,
			err:  nil,
		},
		{
			name: "Float64 - Float64",
			in:   float64(123.46),
			want: 123.46,
			err:  nil,
		},
		{
			name: "Unknown Type",
			in:   []int{1, 2, 3},
			want: 0,
			err:  fmt.Errorf("No known conversion from []int to float"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toFloat(tt.in)
			if err != nil {
				if tt.err == nil {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != tt.err.Error() {
					t.Errorf("Unexpected error: got %v, want %v", err, tt.err)
				}
			} else if math.Abs(got-tt.want) > float64EqualityThreshold {
				t.Errorf("toFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshall(t *testing.T) {
	tests := []struct {
		name  string
		field reflect.Value
		want  string
		err   error
	}{
		{
			name:  "TypeMarshaller",
			field: reflect.ValueOf(&mockTypeMarshaller{}),
			want:  "mockTypeMarshaller.MarshalCSV",
			err:   nil,
		},
		{
			name:  "TextMarshaler",
			field: reflect.ValueOf(&mockTextMarshaler{}),
			want:  "mockTextMarshaler.MarshalText",
			err:   nil,
		},
		{
			name:  "Stringer",
			field: reflect.ValueOf(&mockStringer{}),
			want:  "mockStringer.String",
			err:   nil,
		},
		{
			name:  "NilField",
			field: reflect.ValueOf((*mockTypeMarshaller)(nil)),
			want:  "",
			err:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := marshall(tt.field)
			if err != nil {
				if tt.err == nil {
					t.Errorf("Unexpected error: %v", err)
				} else if err.Error() != tt.err.Error() {
					t.Errorf("Unexpected error: got %v, want %v", err, tt.err)
				}
			} else if got != tt.want {
				t.Errorf("marshall() = %v, want %v", got, tt.want)
			}
		})
	}
}

type mockTypeMarshaller struct{}

func (m *mockTypeMarshaller) MarshalCSV() (string, error) {
	return "mockTypeMarshaller.MarshalCSV", nil
}

type mockTextMarshaler struct{}

func (m *mockTextMarshaler) MarshalText() ([]byte, error) {
	return []byte("mockTextMarshaler.MarshalText"), nil
}

type mockStringer struct{}

func (m *mockStringer) String() string {
	return "mockStringer.String"
}
