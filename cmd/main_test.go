package main

import (
	"fmt"
	"testing"
)

func TestListDeployments(t *testing.T) {
	_, err := ListDeployments("aks-mdh1we-admin")
	if err != nil {
		t.Errorf("Oh crap")
	}
	if err == nil {
		fmt.Printf("\n\nAwesomesauce!!\n")
	}
}

func TestListPods(t *testing.T) {
	_, err := ListPods("aks-mdh1we-admin")
	if err != nil {
		t.Errorf("Oh crap")
	}
	if err == nil {
		fmt.Printf("\n\nAwesomesauce!!\n")
	}
}
