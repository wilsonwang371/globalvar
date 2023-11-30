package pkg

import (
	"fmt"
)

var globalvars = make(map[string]interface{})

// Keys returns a slice of all global variable keys
func Keys() []string {
	keys := make([]string, 0, len(globalvars))
	for k := range globalvars {
		keys = append(keys, k)
	}
	return keys
}

// Get a global variable and return an error if not found
func Get(key string) (interface{}, error) {
	if _, ok := globalvars[key]; ok {
		return globalvars[key], nil
	}
	return nil, fmt.Errorf("key %s not found", key)
}

// Set a global variable and return the previous value if any
func Set(key string, value interface{}) interface{} {
	if _, ok := globalvars[key]; ok {
		oldvalue := globalvars[key]
		globalvars[key] = value
		return oldvalue
	}
	globalvars[key] = value
	return nil
}

// Delete a global variable and return an error if not found
func Delete(key string) error {
	if _, ok := globalvars[key]; ok {
		delete(globalvars, key)
		return nil
	}
	return fmt.Errorf("key %s not found", key)
}

// Clear all global variables
func Clear() {
	globalvars = make(map[string]interface{})
}
