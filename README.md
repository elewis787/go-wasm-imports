### Overview
This demostrates a Golang two imports when using the `//go:wasmimport` pargma. 

## Compile 
Below are the steps used to generate wasip1 modules in go and tiny-go. 

### Go 
`GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go`

#### Go Imports 

```
  (import "wasi_snapshot_preview1" "sched_yield" (func (;0;) (type 1)))
  (import "wasi_snapshot_preview1" "proc_exit" (func (;1;) (type 2)))
  (import "wasi_snapshot_preview1" "args_get" (func (;2;) (type 3)))
  (import "wasi_snapshot_preview1" "args_sizes_get" (func (;3;) (type 3)))
  (import "wasi_snapshot_preview1" "clock_time_get" (func (;4;) (type 4)))
  (import "wasi_snapshot_preview1" "environ_get" (func (;5;) (type 3)))
  (import "wasi_snapshot_preview1" "environ_sizes_get" (func (;6;) (type 3)))
  (import "wasi_snapshot_preview1" "fd_write" (func (;7;) (type 5)))
  (import "wasi_snapshot_preview1" "random_get" (func (;8;) (type 3)))
  (import "wasi_snapshot_preview1" "poll_oneoff" (func (;9;) (type 5)))
  (import "example" "one" (func (;10;) (type 6)))
  (import "example" "one" (func (;11;) (type 6)))
```


### Tiny-go 
`tinygo build -o main.wasm -target=wasip1 main.go`

#### Tiny-go Imports 

```
  (import "wasi_snapshot_preview1" "fd_write" (func $runtime.fd_write (type 2)))
  (import "example" "one" (func $main.importOne (type 0)))
  (import "wasi_snapshot_preview1" "random_get" (func $__imported_wasi_snapshot_preview1_random_get (type 4)))
```

### WAT 
You can inspect the module with tools included in the WebAssembly Binary Toolkit ( wabt )

`wasm2wat main.wasm -o main.wat`



