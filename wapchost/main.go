package main

import (
	"context"
	"fmt"
	"os"

	wapc "github.com/wapc/wapc-go"
	"github.com/wapc/wapc-go/engines/wazero"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: wapchost <value>")
		return
	}
	value := os.Args[1]
	ctx := context.Background()
	guest, err := os.ReadFile("../state/state.wasm")
	if err != nil {
		panic(err)
	}

	engine := wazero.Engine()

	module, err := engine.New(ctx, nil, guest, &wapc.ModuleConfig{
		Logger: wapc.PrintlnLogger,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	})
	if err != nil {
		panic(err)
	}
	defer module.Close(ctx)

	instance, err := module.Instantiate(ctx)
	if err != nil {
		panic(err)
	}
	defer instance.Close(ctx)

	result, err := instance.Invoke(ctx, "saveandrestore", []byte(value))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))
}
