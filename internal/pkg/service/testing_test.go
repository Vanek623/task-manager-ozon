package service

import (
	"github.com/golang/mock/gomock"
	mock_service "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/mocks"
	"testing"
)

type serviceFixture struct {
	service iService
	storage *mock_service.MockiStorage
}

func setUp(t *testing.T) serviceFixture {
	t.Parallel()

	f := serviceFixture{}
	f.storage = mock_service.NewMockiStorage(gomock.NewController(t))
	f.service = New(f.storage)

	return f
}
