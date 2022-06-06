package test

import (
	"fmt"
	"testing"
)

func TestEmptyKey(t *testing.T) {
	fmt.Println("Hello GO")
}

func TestCGO(t *testing.T) {
	fmt.Printf("result : %d\n", helloC())
}
