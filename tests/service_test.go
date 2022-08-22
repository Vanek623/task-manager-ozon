//go:build integration
// +build integration

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
)

func TestService_AddTask(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// arrange
		// act
		_, err := client.TaskCreate(context.Background(), &service.TaskCreateRequest{
			Title: "test",
		})
		// assert
		assert.NoError(t, err)
	})
}

func TestService_DelTask(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// arrange
		ctx := context.Background()
		// act
		addResp, err := client.TaskCreate(ctx, &service.TaskCreateRequest{
			Title: "test",
		})
		require.NoError(t, err)

		_, err = client.TaskDelete(ctx, &service.TaskDeleteRequest{
			ID: addResp.GetID(),
		})
		// assert
		assert.NoError(t, err)
	})

	t.Run("no_task", func(t *testing.T) {
		// arrange
		ctx := context.Background()
		// act
		addResp, createErr := client.TaskCreate(ctx, &service.TaskCreateRequest{
			Title: "test",
		})
		require.NoError(t, createErr)

		delRequest := &service.TaskDeleteRequest{ID: addResp.GetID()}
		var deleteErr, secondDeleteErr error
		_, deleteErr = client.TaskDelete(ctx, delRequest)
		require.NoError(t, deleteErr)

		_, secondDeleteErr = client.TaskDelete(ctx, delRequest)
		// assert
		assert.Error(t, secondDeleteErr)
	})
}

func TestService_UpdateTask(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// arrange
		ctx := context.Background()

		// act
		addResp, createErr := client.TaskCreate(ctx, &service.TaskCreateRequest{
			Title: "test_init",
		})
		require.NoError(t, createErr)

		var updateErr error
		_, updateErr = client.TaskUpdate(ctx, &service.TaskUpdateRequest{
			ID:    addResp.GetID(),
			Title: "test_edit",
		})
		// assert
		assert.NoError(t, updateErr)
	})

	t.Run("no_task", func(t *testing.T) {
		// arrange
		ctx := context.Background()

		// act
		_, _ = client.TaskDelete(ctx, &service.TaskDeleteRequest{
			ID: 1,
		})

		_, err := client.TaskUpdate(ctx, &service.TaskUpdateRequest{
			ID:    1,
			Title: "test_edit",
		})

		// assert
		assert.Error(t, err)
	})
}
