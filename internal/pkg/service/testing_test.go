package service

import (
	"github.com/golang/mock/gomock"
	mock_service "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/mocks"
	"testing"
)

type serviceFixture struct {
	service *Service
	storage *mock_service.MockiStorage
}

func setUp(t *testing.T) serviceFixture {
	t.Parallel()
	var f serviceFixture
	var err error
	f.storage = mock_service.NewMockiStorage(gomock.NewController(t))
	f.service, err = New(f.storage)

	if err != nil {
		t.Fatal(err)
	}

	return f
}