package main

import (
	"fmt"

	"github.com/agrea/ptr"
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"

	"github.com/deislabs/dapr-wasm-exp/state/inmemory"

	wapc "github.com/wapc/wapc-guest-tinygo"
)

func main() {

	wapc.RegisterFunctions(wapc.Functions{
		"saveandrestore": saveandrestore,
	})
}

func saveandrestore(payload []byte) ([]byte, error) {
	fmt.Println("saveandrestore called")
	store := inmemory.NewInMemoryStateStore(logger.NewLogger("test"))
	store.Init(state.Metadata{})
	keyA := "theFirstKey"
	valueA := string(payload)
	fmt.Printf("ValueA: %s\n", valueA)

	// set
	setReq := &state.SetRequest{
		Key:      keyA,
		Value:    valueA,
		ETag:     ptr.String("the etag"),
		Metadata: map[string]string{"ttlInSeconds": "100"},
	}
	err := store.Set(setReq)
	if err != nil {
		panic(err)
	}
	// get after set
	getReq := &state.GetRequest{
		Key: keyA,
	}

	resp, err := store.Get(getReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ValueA: %s\n", string(resp.Data))

	return []byte(fmt.Sprintf("round trip: %s\n", string(resp.Data))), nil
}
