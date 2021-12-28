package main

import trace "github.com/i6u/instrument-trace"

func foo() {
	defer trace.Trace()()
	bar()
}
func bar() {
	defer trace.Trace()()
}
func main() {
	defer trace.Trace()()
	foo()
}
