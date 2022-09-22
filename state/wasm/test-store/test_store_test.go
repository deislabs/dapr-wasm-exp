package test_store

import (
	"context"
	"github.com/dapr/components-contrib/state/wasm/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadAndWrite(t *testing.T) {

	ctx := context.Background()
	p, err := state.NewStorePlugin(ctx, state.StorePluginOption{})
	assert.Nil(t, err)

	testPlugin, err := p.Load(ctx, "test_store.wasm")
	assert.Nil(t, err)

	t.Run("simple get", func(t *testing.T) {
		// set
		getReq, err := testPlugin.Get(ctx, state.GetRequest{})
		assert.Nil(t, err)
		assert.NotNil(t, getReq)
		assert.Equal(t, "test1", getReq.Etag)
	})
}
