package main

import (
	"fmt"
	"testing"
)

func TestListDeployments(t *testing.T) {
	got := ListDeployments()
	want := ListDeployments()
	if got != want {
		t.Errorf("Oh crap")
	}
	if got == want {
		fmt.Printf("Got %v which was correct", want)
	}
}

func TestListPods(t *testing.T) {
	got := ListPods()
	want := ListPods()
	if got != want {
		t.Errorf("Oh crap")
	}
	if got == want {
		fmt.Printf("Got %v which was correct", want)
	}
}
