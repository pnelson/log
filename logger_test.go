package log

import (
	"reflect"
	"testing"
)

func TestDataKeys(t *testing.T) {
	m := M{"c": 3, "a": 1, "b": 2}
	have := m.Keys()
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(have, want) {
		t.Fatalf("data keys should be sorted\nhave %v\nwant %v", have, want)
	}
}
