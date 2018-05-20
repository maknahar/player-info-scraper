package gutils

import (
	"testing"
	"runtime"
)

func ShouldBeEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		_, fn, line, _ := runtime.Caller(1)
		t.Fatalf("Expected %v. Got %v. Location: %s:%d", expected, actual, fn, line)
	}
}
