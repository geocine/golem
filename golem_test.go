package main

import "testing"

func TestSplitTrimString(t *testing.T) {
	splitString := SplitTrimString("Hello | world")
	if splitString != "Hello" {
		t.Errorf("string was incorrect, got: %s, want: %s.", splitString, "Hello")
	}
}
