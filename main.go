package main

//go:wasmimport example one
func importOne()

//go:wasmimport example one
func importTwo()

func main() {
	importOne()
	importTwo()
}
