package command

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock_service "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/mocks"
)

type serviceFixture struct {
	Ctx     context.Context
	command ICommand
	service *mock_service.MockiService
}

func addCommandSetUp(t *testing.T) serviceFixture {
	f := commonSetUp(t)
	f.command = newAddCommand(f.service)

	return f
}

func delCommandSetUp(t *testing.T) serviceFixture {
	f := commonSetUp(t)
	f.command = newDeleteCommand(f.service)

	return f
}

func getCommandSetUp(t *testing.T) serviceFixture {
	f := commonSetUp(t)
	f.command = newGetCommand(f.service)

	return f
}

func commonSetUp(t *testing.T) serviceFixture {
	t.Parallel()

	return serviceFixture{
		service: mock_service.NewMockiService(gomock.NewController(t)),
	}
}
