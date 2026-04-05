//go:build go1.26
// +build go1.26

// The idea behind this package is to help handle with pointers in very common tasks.
//
// It exposes just a single public function [To]
package pointer

// To returns a pointer to an object of a given type.
//
// On Go 1.26+, this is a simple wrapper around the built-in new function.
//
//	p := pointer.To(64) // assume type is int
//
//	fmt.Printf("return a pointer to %T with value %v", *p, *p)
//
// Outputs: "return a pointer to int with value 64"
func To[V any](object V) *V {
	p := new(V)
	*p = object
	return p
}