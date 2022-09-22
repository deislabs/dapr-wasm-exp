package wasm

import (
	"github.com/dapr/components-contrib/state"
	"github.com/dapr/kit/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHardcodedGet(t *testing.T) {
	store, err := NewStateWasmStore(logger.NewLogger("test"), "./test-store/test_store.wasm")

	assert.Nil(t, err)

	t.Run("simple get", func(t *testing.T) {

		getReq := &state.GetRequest{
			Key: "test",
		}

		getRes, err := store.Get(getReq)
		assert.Nil(t, err)
		assert.NotNil(t, getReq)
		assert.Equal(t, "etag", *getRes.ETag)
	})
}
