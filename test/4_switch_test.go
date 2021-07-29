package test

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	useManyValSwitch("v")
	useManyValSwitch("a")

	useExpressionSwitch(11)
}

func useManyValSwitch(a string) {
	switch a {
	case "a", "b", "c":
		fmt.Println("one of a,b,c")
	case "d", "e", "f":
		fmt.Println("one of d,e,f")
	default:
		fmt.Println("not at a-f")
	}
}

func useExpressionSwitch(i int) {
	switch {
	case i >= 1 && i < 10:
		fmt.Println(i," is between 1 and 9")
	case i >= 10 && i< 20:
		fmt.Println(i," is between 10 and 19")
	default:
		fmt.Println(i," is more than 19")
	}
}
