package main

import (
	"github.com/wilsonwang371/globalvar/pkg"
)

func main() {
	// Set a global variable
	pkg.Set("foo", "bar")
	// Get a global variable
	if value, err := pkg.Get("foo"); err == nil {
		if value != "bar" {
			panic(value)
		}
	} else {
		panic(value)
	}
	// Delete a global variable
	pkg.Delete("foo")
	if _, err := pkg.Get("foo"); err == nil {
		panic("foo should be deleted")
	}
	// Clear all global variables
	pkg.Clear()
}
