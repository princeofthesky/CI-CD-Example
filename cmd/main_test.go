package main

import "testing"

func TestA(t *testing.T) {
	if 1 == 2 {
		t.Fail()
	}
}
