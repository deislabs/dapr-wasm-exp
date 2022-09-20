package main

import (
	"fmt"

	"github.com/agrea/ptr"
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"

	"github.com/deislabs/state/inmemory"
)

func main() {
	store := inmemory.NewInMemoryStateStore(logger.NewLogger("test"))
	store.Init(state.Metadata{})
	keyA := "theFirstKey"
	valueA := "value of key"
	// set
	setReq := &state.SetRequest{
		Key:   keyA,
		Value: valueA,
		ETag:  ptr.String("the etag"),
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

	fmt.Printf("round trip: %s\n", string(resp.Data))
}
