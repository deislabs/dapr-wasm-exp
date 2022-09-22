//go:build tinygo.wasm

package main

import (
	"context"

	"github.com/dapr/components-contrib/state/wasm/state"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	state.RegisterStore(TestStore{})
}

type TestStore struct{}

func (m TestStore) Get(context.Context, state.GetRequest) (state.GetResponse, error) {
	return state.GetResponse{
		Etag: "etag",
	}, nil
}

func (m TestStore) Init(context.Context, state.InitRequest) (state.InitResponse, error) {
	return state.InitResponse{}, nil
}
func (m TestStore) Features(context.Context, state.FeaturesRequest) (state.FeaturesResponse, error) {
	return state.FeaturesResponse{}, nil
}
func (m TestStore) Set(context.Context, state.SetRequest) (state.SetResponse, error) {
	return state.SetResponse{}, nil
}
func (m TestStore) Delete(context.Context, state.DeleteRequest) (state.DeleteResponse, error) {
	return state.DeleteResponse{}, nil
}
