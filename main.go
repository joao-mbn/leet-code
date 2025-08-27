// Package main is the entry point for the program'.'
package main

import "fmt"

func main() {
	for _, input := range []struct {
		path string
	}{
		{path: "/home/"},
		{path: "/home//foo/"},
		{path: "/home/user/Documents/../Pictures"},
		{path: "/../"},
		{path: "/.../a/../b/c/../d/./"},
		{path: "/a/./b/../../c/"},
	} {
		result := simplifyPath(input.path)
		fmt.Printf("Given the input: %v, the result is: %v\n", input.path, result)
	}
}
