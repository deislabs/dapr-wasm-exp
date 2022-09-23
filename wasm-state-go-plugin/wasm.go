/*
Copyright 2021 The Dapr Authors
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package wasm_state_go_plugin

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"
	stateWasm "github.com/deislabs/dapr-wasm-exp/wasm-state-go-plugin/state"
)

type wasmStore struct {
	store stateWasm.Store
}

func NewStateWasmStore(logger logger.Logger, path string) (state.Store, error) {
	ctx := context.Background()
	storePlugin, err := stateWasm.NewStorePlugin(ctx, stateWasm.StorePluginOption{})
	if err != nil {
		return nil, err
	}

	store, err := storePlugin.Load(ctx, path)

	return &wasmStore{
		store: store,
	}, nil
}

func (wasmStore *wasmStore) Init(metadata state.Metadata) error {
	_, err := wasmStore.store.Init(context.Background(), stateWasm.InitRequest{Metadata: &stateWasm.Metadata{}})

	return err
}

func (wasmStore *wasmStore) Close() error {
	return nil
}

func (wasmStore *wasmStore) Features() []state.Feature {
	return []state.Feature{state.FeatureETag, state.FeatureTransactional}
}

func (wasmStore *wasmStore) Delete(req *state.DeleteRequest) error {
	var etag string = ""

	if req.ETag != nil {
		etag = *req.ETag
	}

	_, err := wasmStore.store.Delete(context.Background(), stateWasm.DeleteRequest{
		Key:      req.Key,
		Etag:     etag,
		Metadata: req.Metadata,
		Options:  &stateWasm.DeleteStateOption{Concurrency: req.Options.Concurrency, Consistency: req.Options.Consistency},
	})

	return err
}

func (wasmStore *wasmStore) Get(req *state.GetRequest) (*state.GetResponse, error) {
	res, err := wasmStore.store.Get(context.Background(), stateWasm.GetRequest{
		Key:      req.Key,
		Metadata: req.Metadata,
		Options:  &stateWasm.GetStateOption{Consistency: req.Options.Consistency},
	})

	if err != nil {
		return nil, err
	}

	return &state.GetResponse{
		Data:        res.Data,
		ETag:        &res.Etag,
		Metadata:    res.Metadata,
		ContentType: &res.ContentType,
	}, nil
}

func (wasmStore *wasmStore) Set(req *state.SetRequest) error {
	var etag string = ""
	var contentType string = ""

	if req.ETag != nil {
		etag = *req.ETag
	}

	if req.ContentType != nil {
		contentType = *req.ContentType
	}

	data := getBytes(req.Value)
	_, err := wasmStore.store.Set(context.Background(), stateWasm.SetRequest{
		Key:         req.Key,
		Data:        data,
		Etag:        etag,
		Metadata:    req.Metadata,
		Options:     &stateWasm.SetStateOption{Consistency: req.Options.Consistency},
		ContentType: contentType,
	})

	return err
}

func getBytes(key interface{}) []byte {
	switch key.(type) {
	case string:
		return []byte(key.(string))
	default:
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(key)
		if err != nil {
			return []byte{}
		}
		return buf.Bytes()
	}
}

func (wasmStore *wasmStore) BulkSet(req []state.SetRequest) error {
	return nil
}

func (wasmStore *wasmStore) Multi(request *state.TransactionalStateRequest) error {
	return nil
}

func (wasmStore *wasmStore) BulkDelete(req []state.DeleteRequest) error {
	return nil
}

func (wasmStore *wasmStore) BulkGet(req []state.GetRequest) (bool, []state.BulkGetResponse, error) {
	return false, nil, nil
}
