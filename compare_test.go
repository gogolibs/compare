package compare_test

import (
	"github.com/gogolibs/compare"
	"testing"
)

func TestCompare(t *testing.T) {
	if !compare.Equal(1, 1) {
		t.Fatal("bad Equal")
	}
	if !compare.Less(1, 2) {
		t.Fatal("bad Less")
	}
	if !compare.Greater(2, 1) {
		t.Fatal("bad Greater")
	}
	if !compare.LessOrEqual(1, 1) {
		t.Fatal("bad LessOrEqual")
	}
	if !compare.LessOrEqual(1, 2) {
		t.Fatal("bad LessOrEqual")
	}
	if !compare.GreaterOrEqual(1, 1) {
		t.Fatal("bad LessOrEqual")
	}
	if !compare.GreaterOrEqual(2, 1) {
		t.Fatal("bad LessOrEqual")
	}
	if !compare.Reversed(compare.Greater[int])(1, 2) {
		t.Fatal("bad Reversed")
	}
}
