package wasm_state_go_plugin

import (
	"testing"

	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"
	"github.com/stretchr/testify/assert"
)

func TestReadAndWrite(t *testing.T) {
	store, err := NewStateWasmStore(logger.NewLogger("test"), "./test-store/test_store.wasm")
	assert.Nil(t, err)

	store.Init(state.Metadata{})

	keyA := "theFirstKey"
	valueA := "value of key"
	etag := "etag1"
	contentType := "text"
	t.Run("set kv with etag and then get", func(t *testing.T) {
		// set
		setReq := &state.SetRequest{
			Key:         keyA,
			Value:       valueA,
			ETag:        &etag,
			Metadata:    nil,
			Options:     state.SetStateOption{},
			ContentType: &contentType,
		}
		err := store.Set(setReq)
		assert.Nil(t, err)
		// get after set
		getReq := &state.GetRequest{
			Key: keyA,
		}
		resp, err := store.Get(getReq)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, valueA, string(resp.Data))
	})

	t.Run("set and get the second key successfully", func(t *testing.T) {
		// set
		setReq := &state.SetRequest{
			Key:   "theSecondKey",
			Value: "1234",
		}
		err := store.Set(setReq)
		assert.Nil(t, err)
		// get
		getReq := &state.GetRequest{
			Key: "theSecondKey",
		}
		resp, err := store.Get(getReq)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "1234", string(resp.Data))
	})

	t.Run("delete theFirstKey", func(t *testing.T) {
		req := &state.DeleteRequest{
			Key: "theFirstKey",
		}
		err := store.Delete(req)
		assert.Nil(t, err)
	})
}
