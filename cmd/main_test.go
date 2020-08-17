package main

import (
	"testing"
)

// Ensure that the passed in cluster name exists in your context before running the test
func TestListDeployments(t *testing.T) {
	_, err := ListDeployments("aks-ci1we-admin")
	if err != nil {
		t.Errorf("Error - %v", err)
	}
}

// Ensure that the passed in cluster name exists in your context before running the test
func TestListPods(t *testing.T) {
	_, err := ListPods("aks-ci1we-admin")
	if err != nil {
		t.Errorf("Error - %v", err)
	}
}
