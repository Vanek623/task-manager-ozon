package command

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
)

func TestAddCommand_Execute(t *testing.T) {
	t.Run("normal_full", func(t *testing.T) {
		// arrange
		f := addCommandSetUp(t)
		f.service.EXPECT().AddTask(gomock.Any(), models.NewAddTaskData("test", "test")).
			Return(uint64(1), nil)
		// act
		res := f.command.Execute(f.Ctx, "test test")
		// assert
		assert.Equal(t, res, "Task #1 added")
	})

	t.Run("no_args", func(t *testing.T) {
		// arrange
		f := addCommandSetUp(t)
		// act
		res := f.command.Execute(f.Ctx, "")
		// assert
		assert.NotEqual(t, res, "Task #1 added")
	})

	t.Run("internal_error", func(t *testing.T) {
		// arrange
		f := addCommandSetUp(t)
		f.service.EXPECT().AddTask(gomock.Any(), models.NewAddTaskData("test", "test")).
			Return(uint64(0), errors.New("internal error"))
		// act
		res := f.command.Execute(f.Ctx, "test test")
		// assert
		assert.NotEqual(t, res, "Task #1 added")
	})
}

func TestDeleteCommand_Execute(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// arrange
		f := delCommandSetUp(t)
		f.service.EXPECT().DeleteTask(gomock.Any(), models.NewDeleteTaskData(1)).Return(nil)
		// act
		res := f.command.Execute(f.Ctx, "1")
		// assert
		assert.Equal(t, res, "Task deleted")
	})

	t.Run("no args", func(t *testing.T) {
		// arrange
		f := delCommandSetUp(t)
		// act
		res := f.command.Execute(f.Ctx, "")
		// assert
		assert.NotEqual(t, res, "Task deleted")
	})

	t.Run("incorrect id", func(t *testing.T) {
		// arrange
		f := delCommandSetUp(t)
		// act
		res := f.command.Execute(f.Ctx, "a")
		// assert
		assert.NotEqual(t, res, "Task deleted")
	})

	t.Run("no task", func(t *testing.T) {
		// arrange
		f := delCommandSetUp(t)
		f.service.EXPECT().DeleteTask(gomock.Any(), models.NewDeleteTaskData(1)).Return(errors.New(""))
		// act
		res := f.command.Execute(f.Ctx, "1")
		// assert
		assert.NotEqual(t, res, "Task deleted")
	})
}

func TestGetCommand_Execute(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// arrange
		f := getCommandSetUp(t)
		actual := models.NewDetailedTask("test", "test", time.Now())
		f.service.EXPECT().GetTask(gomock.Any(), models.NewGetTaskData(1)).Return(actual, nil)
		// act
		res := f.command.Execute(f.Ctx, "1")
		// assert
		assert.Equal(t, res, actual.String())
	})

	t.Run("no_args", func(t *testing.T) {
		// arrange
		f := getCommandSetUp(t)
		actual := models.NewDetailedTask("test", "test", time.Now())
		f.service.EXPECT().GetTask(gomock.Any(), models.NewGetTaskData(1)).Return(nil, errors.New(""))
		// act
		res := f.command.Execute(f.Ctx, "1")
		// assert
		assert.NotEqual(t, res, actual.String())
	})

	t.Run("internal error", func(t *testing.T) {
		// arrange
		f := getCommandSetUp(t)
		actual := models.NewDetailedTask("test", "test", time.Now())
		f.service.EXPECT().GetTask(gomock.Any(), models.NewGetTaskData(1)).Return(nil, errors.New(""))
		// act
		res := f.command.Execute(f.Ctx, "1")
		// assert
		assert.NotEqual(t, res, actual.String())
	})
}
