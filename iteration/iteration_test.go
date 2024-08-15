package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T){
	repeated := Repeat(5, "a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat(){
	repeatedString := Repeat(5, "a")
	fmt.Println(repeatedString)
}
