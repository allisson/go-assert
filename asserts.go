package assert

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Equal compares two values using reflect.DeepEqual and fails if not equals
func Equal(t *testing.T, expected, current interface{}) {
	t.Helper()

	if !reflect.DeepEqual(expected, current) {
		t.Fatalf("assertion_type=Equal, expected_value=%#v, expected_type=%T, current_value=%#v, current_type=%T, diff=(-expected +current):\n%s", expected, expected, current, current, cmp.Diff(expected, current))
	}
}

// NotEqual compares two values using reflect.DeepEqual and fails if equals
func NotEqual(t *testing.T, expected, current interface{}) {
	t.Helper()

	if reflect.DeepEqual(expected, current) {
		t.Fatalf("assertion_type=NotEqual, expected_value=%#v, expected_type=%T, current_value=%#v, current_type=%T", expected, expected, current, current)
	}
}

// containsKind checks if a specified kind in the slice of kinds.
func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}

// isNil checks if a specified object is nil or not, without Failing.
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && value.IsNil() {
		return true
	}

	return false
}

// Nil fails if value is not nil
func Nil(t *testing.T, current interface{}) {
	t.Helper()

	if !isNil(current) {
		t.Fatalf("assertion_type=Nil, current_value=%#v, current_type=%T", current, current)
	}
}

// NotNil fails is value is nil
func NotNil(t *testing.T, current interface{}) {
	t.Helper()

	if isNil(current) {
		t.Fatalf("assertion_type=NotNil, current_value=%#v, current_type=%T", current, current)
	}
}
