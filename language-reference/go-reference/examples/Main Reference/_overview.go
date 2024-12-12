/*
*Modules* are the top-level organizational unit of Go software. A single source code repository might contain multiple
modules, though it typically will contain just one. A module is a collection of related *packages* that get released
together.

A *package* is a collection of source files in a single directory. Any names defined in one source file are visible to 
other source files in the same package.

Begin a new Go project by creating a directory for the source files and running `$ go mod init [my-cool-go-project]`
inside the new directory. Be sure to run a `$ git init` afterwords and push up to a remote repository.
*/

package main

// A struct is defined as
type MyStruct struct {
	
}