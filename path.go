package main

// Package main provides path constants and a function to get paths.
const (
	ResetPassword = "/reset-password"
)


// Get returns the path corresponding to the given path type.
func Get(pt string) string {
	return "this is my path: " + pt
}