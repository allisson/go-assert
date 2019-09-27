package assert

import "testing"

func TestEqual(t *testing.T) {
	var tests = []struct {
		expected interface{}
		current  interface{}
	}{
		{1, 1},
		{"1", "1"},
		{1.0, 1.0},
		{[]byte("byte"), []byte("byte")},
		{nil, nil},
		{map[string]string{}, map[string]string{}},
		{struct{ a string }{"a"}, struct{ a string }{"a"}},
	}
	for _, tt := range tests {
		Equal(t, tt.expected, tt.current)
	}
}

func TestNotEqual(t *testing.T) {
	var tests = []struct {
		expected interface{}
		current  interface{}
	}{
		{1, 2},
		{"1", "2"},
		{1.0, 2.0},
		{[]byte("byte"), []byte("byte2")},
		{nil, 0},
		{map[string]string{}, map[string]int{}},
		{struct{ a string }{"a"}, struct{ a string }{"b"}},
	}
	for _, tt := range tests {
		NotEqual(t, tt.expected, tt.current)
	}
}

type myStruct struct {
	a *string
}

func TestNil(t *testing.T) {
	Nil(t, nil)
	m := myStruct{}
	Nil(t, m.a)
}

func TestNotNil(t *testing.T) {
	var tests = []struct {
		current interface{}
	}{
		{1},
		{"1"},
		{1.0},
		{[]byte("byte")},
		{map[string]string{}},
		{struct{ a string }{"a"}},
	}
	for _, tt := range tests {
		NotNil(t, tt.current)
	}
}
