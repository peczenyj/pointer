// The idea behing this package is to help handle with pointers in very common tasks.
//
// It exposes just a single public function [To]
package pointer

// To returns a pointer to an object of a given type.
//
//	p := pointer.To(64) // assume type is int
//
//	fmt.Printf("return a pointer to %T with value %v", *p, *p)
//
// Outputs: "return a pointer to int with value 64"
func To[V any](object V) *V {
	return &object
}
