package xlsxer

import (
	"reflect"
	"testing"
)

func TestFieldInfoMatchesKey(t *testing.T) {
	tests := []struct {
		name     string
		f        fieldInfo
		key      string
		expected bool
	}{
		{
			name: "Exact Match",
			f: fieldInfo{
				keys:    []string{"key1", "key2", "key3"},
				partial: false,
			},
			key:      "key2",
			expected: true,
		},
		{
			name: "Trimmed Match",
			f: fieldInfo{
				keys:    []string{"key1", "key2", "key3"},
				partial: false,
			},
			key:      "  key3  ",
			expected: true,
		},
		{
			name: "Partial Match",
			f: fieldInfo{
				keys:    []string{"key1", "key2", "key3"},
				partial: true,
			},
			key:      "key1extra",
			expected: true,
		},
		{
			name: "Zero Width Characters Match",
			f: fieldInfo{
				keys:    []string{"key1", "key2", "key3"},
				partial: false,
			},
			key:      "k\u200Bey2",
			expected: true,
		},
		{
			name: "No Match",
			f: fieldInfo{
				keys:    []string{"key1", "key2", "key3"},
				partial: false,
			},
			key:      "nonexistent",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.f.matchesKey(tt.key)
			if got != tt.expected {
				t.Errorf("matchesKey() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsErrorType(t *testing.T) {
	tests := []struct {
		name     string
		outType  reflect.Type
		expected bool
	}{
		{
			name:     "Interface Type",
			outType:  reflect.TypeOf((*error)(nil)).Elem(),
			expected: true,
		},
		{
			name:     "Non-Interface Type",
			outType:  reflect.TypeOf(""),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isErrorType(tt.outType)
			if got != tt.expected {
				t.Errorf("isErrorType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestGetConcreteReflectValueAndType(t *testing.T) {
	tests := []struct {
		name     string
		in       interface{}
		expected reflect.Type
	}{
		{
			name:     "Test Case 1",
			in:       "test",
			expected: reflect.TypeOf(""),
		},
		{
			name:     "Test Case 2",
			in:       123,
			expected: reflect.TypeOf(0),
		},
		{
			name:     "Test Case 3",
			in:       true,
			expected: reflect.TypeOf(false),
		},
		// Add more test cases here if needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotType := getConcreteReflectValueAndType(tt.in)
			if gotType != tt.expected {
				t.Errorf("getConcreteReflectValueAndType() = %v, want %v", gotType, tt.expected)
			}
		})
	}
}
