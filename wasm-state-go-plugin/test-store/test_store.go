//go:build tinygo.wasm

package main

import (
	"context"
	"errors"
	"github.com/deislabs/dapr-wasm-exp/wasm-state-go-plugin/state"
)

// main is required for TinyGo to compile to Wasm.
func main() {
	state.RegisterStore(TestStore{items: map[string]*inMemStateStoreItem{}})
}

type inMemStateStoreItem struct {
	data []byte
}

type TestStore struct {
	items map[string]*inMemStateStoreItem
}

func (store TestStore) Get(_ context.Context, req state.GetRequest) (state.GetResponse, error) {
	item, ok := store.items[req.Key]

	if !ok {
		return state.GetResponse{}, errors.New("unknown key")
	}

	return state.GetResponse{
		Data:        item.data,
		Etag:        "test",
		Metadata:    nil,
		ContentType: "",
	}, nil
}

func (store TestStore) Init(context.Context, state.InitRequest) (state.InitResponse, error) {
	return state.InitResponse{}, nil
}

func (store TestStore) Features(context.Context, state.FeaturesRequest) (state.FeaturesResponse, error) {
	return state.FeaturesResponse{Features: []string{"TRANSACTIONAL"}}, nil
}

func (store TestStore) Set(_ context.Context, req state.SetRequest) (state.SetResponse, error) {
	store.items[req.Key] = &inMemStateStoreItem{
		data: req.Data,
	}

	return state.SetResponse{}, nil
}

func (store TestStore) Delete(_ context.Context, req state.DeleteRequest) (state.DeleteResponse, error) {
	delete(store.items, req.Key)

	return state.DeleteResponse{}, nil
}
