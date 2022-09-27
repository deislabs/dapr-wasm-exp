# Dapr'ish Wasm State Go Plugin
This experiment illustrates how Wasm and the [Go Plugin System over WebAssembly](https://github.com/knqyf263/go-plugin) could be used to create a dynamically loaded plugin model to embed the capabilities of Dapr within an application. If embedded and dynamically loaded in Dapr, it could reduce the size of Dapr to only the capabilities needed, rather than all capabilities in the Dapr universe. If embedded in a non-Dapr application, it could enable the application to eliminate the need for running a Dapr sidecar.

## What's in this directory
```
.
├── README.md
├── go.mod
├── go.sum
├── state
│   ├── state.pb.go
│   ├── state.proto
│   ├── state_host.pb.go
│   ├── state_plugin.pb.go
│   └── state_vtproto.pb.go
├── test-store
│   ├── test_store.go
│   └── test_store.wasm
├── wasm.go
└── wasm_test.go
```

- [state directory](./state): contains the protobuf interface for the state store which defines the contract between the state store Wasm module and the consuming host.
- [test-store directory](./test-store): contains the Go implementation of an in-memory state store which will be compiled to Wasm and the used from [wasm.go](./wasm.go).
- [wasm_test.go](./wasm_test.go): is a set of tests that exercise wasm.go, which loads the Wasm in-memory state store and verifies the calls to the state store API.

## Building and Executing the Project
We have committed the state store Wasm module which we compiled ahead of time, so if you would like to simply run the tests, you can do that by executing `go test -v .` in this directory. You should see the following output.

```
$ go test . -v
=== RUN   TestReadAndWrite
=== RUN   TestReadAndWrite/set_kv_with_etag_and_then_get
=== RUN   TestReadAndWrite/set_and_get_the_second_key_successfully
=== RUN   TestReadAndWrite/delete_theFirstKey
--- PASS: TestReadAndWrite (0.07s)
    --- PASS: TestReadAndWrite/set_kv_with_etag_and_then_get (0.00s)
    --- PASS: TestReadAndWrite/set_and_get_the_second_key_successfully (0.00s)
    --- PASS: TestReadAndWrite/delete_theFirstKey (0.00s)
PASS
ok  	github.com/deislabs/dapr-wasm-exp/wasm-state-go-plugin	0.248s
```

### Building the test-store Wasm
To build the test-store Wasm bin, you will need [tinygo](https://tinygo.org/getting-started/install/) installed.

Once tinygo is installed, you should be able to build the Wasm bin using the following command.
```shell
tinygo build -x -wasm-abi=generic -target=wasi -o test-store/test_store.wasm test-store/test_store.go
```

## Challenges
- os.Exec is not implemented in tinygo and initially provided compile errors
- `reflect.ArrayOf` and `reflect.MapOf` is not implemented and used in the package mapstructure, which is used by other Dapr components.
- `json-iterator` would not compile with tinygo and needed to be replaced when building the original [Dapr in-memory state store](https://github.com/dapr/components-contrib/blob/master/state/in-memory/in_memory.go).
- Timers were not implemented and needed a dev build of tinygo, but are still unstable. These were used in the original [Dapr in-memory state store here](https://github.com/dapr/components-contrib/blob/ad5397ea25489943acf5156c8af764e6f79ea788/state/in-memory/in_memory.go#L395). 
