// Package hello provides a simple greeting.
package hello

import "fmt"

// SayHello returns a string containing a greeting.
func SayHello(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

// PrintSayHello prints a greeting.
func PrintSayHello(name string) {
	fmt.Println(SayHello(name))
}
