package pointer_test

import (
	"fmt"

	"github.com/peczenyj/pointer"
)

func ExampleTo() {
	p := pointer.To(64) // assume type is int

	if p != nil {
		fmt.Printf("return a pointer to %T with value %v", *p, *p)
	}

	// Output:
	// return a pointer to int with value 64
}

func ExampleTo_int32() {
	p := pointer.To[int32](127)

	if p != nil {
		fmt.Printf("return a pointer to %T with value %v", *p, *p)
	}

	// Output:
	// return a pointer to int32 with value 127
}

func ExampleTo_string() {
	p := pointer.To("pointer")

	if p != nil {
		fmt.Printf("return a pointer to %T with value %q", *p, *p)
	}

	// Output:
	// return a pointer to string with value "pointer"
}

func ExampleTo_nil() {
	p := pointer.To[*string](nil) // pointer to a pointer

	if p != nil {
		fmt.Printf("return a pointer to %T with value %v", *p, *p)
	}

	// Output:
	// return a pointer to *string with value <nil>
}
