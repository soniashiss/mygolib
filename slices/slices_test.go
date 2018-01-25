package slices

import (
	"testing"
	"reflect"
)

func TestReverseStringSlice(t *testing.T) {
	oldSlice := []string{"aa", "bb", "cc"}
	wantSlice := []string{"cc", "bb", "aa"}
	gotSlice := ReverseStringSlice(oldSlice)

	if !reflect.DeepEqual(wantSlice, gotSlice) {
		t.Errorf("want: %v, got: %v", wantSlice, gotSlice)
	}
}
