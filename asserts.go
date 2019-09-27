package assert

import (
	"reflect"
	"testing"
)

// Equal compares two values using reflect.DeepEqual and fails if not equals
func Equal(t *testing.T, expected, current interface{}) {
	t.Helper()

	if !reflect.DeepEqual(expected, current) {
		t.Fatalf("assertion_type=Equal, expected_value=%#v, expected_type=%T, current_value=%#v, current_type=%T", expected, expected, current, current)
	}
}

// NotEqual compares two values using reflect.DeepEqual and fails if equals
func NotEqual(t *testing.T, expected, current interface{}) {
	t.Helper()

	if reflect.DeepEqual(expected, current) {
		t.Fatalf("assertion_type=NotEqual, expected_value=%#v, expected_type=%T, current_value=%#v, current_type=%T", expected, expected, current, current)
	}
}

// Nil fails if value is not nil
func Nil(t *testing.T, current interface{}) {
	t.Helper()

	if !reflect.DeepEqual(nil, current) {
		t.Fatalf("assertion_type=Nil, current_value=%#v, current_type=%T", current, current)
	}
}

// NotNil fails is value is nil
func NotNil(t *testing.T, current interface{}) {
	t.Helper()

	if reflect.DeepEqual(nil, current) {
		t.Fatalf("assertion_type=NotNil, current_value=%#v, current_type=%T", current, current)
	}
}
