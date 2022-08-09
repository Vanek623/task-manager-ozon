package service

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	serverPkg "gitlab.ozon.dev/Vanek623/task-manager-system/cmd/server"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address           = "localhost:8082"
	reconnectMaxCount = 5
	reconnectTimeout  = time.Second
)

type storage struct {
	iTaskStorage
	client pb.StorageClient
}

func newStorage() (*storage, error) {
	time.Sleep(reconnectTimeout)

	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	for count := 1; err != nil || con == nil; count++ {
		if count > reconnectMaxCount {
			return nil, errors.Wrap(err, "service_storage: cannot connect to storage server")
		}

		log.Printf("cannot connect to server, try to connect #%d of %d in %d\n", count, reconnectMaxCount, reconnectTimeout)
		time.Sleep(reconnectTimeout)
		con, err = grpc.Dial(serverPkg.FullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &storage{
		client: pb.NewStorageClient(con),
	}, nil
}

func (s storage) Add(ctx context.Context, t models.Task) (uint, error) {
	resp, err := s.client.TaskAdd(ctx, &pb.TaskAddRequest{Task: codeTask(&t)})
	if err != nil {
		return 0, err
	}

	return uint(resp.GetID()), nil
}

func (s storage) Delete(ctx context.Context, ID uint) error {
	_, err := s.client.TaskDelete(ctx, &pb.TaskDeleteRequest{ID: uint64(ID)})
	if err != nil {
		return err
	}

	return nil
}

func (s storage) List(ctx context.Context, limit, offset uint) ([]models.Task, error) {
	resp, err := s.client.TaskList(ctx, &pb.TaskListRequest{
		Limit:  uint32(limit),
		Offset: uint32(offset),
	})

	if err != nil {
		return nil, err
	}

	res := make([]models.Task, 0, len(resp.GetTasks()))
	for _, task := range resp.GetTasks() {
		decoded, err := decodeTask(task)
		if err != nil {
			return nil, err
		}
		res = append(res, *decoded)
	}

	return res, nil
}

func (s storage) Update(ctx context.Context, t models.Task) error {
	_, err := s.client.TaskUpdate(ctx, &pb.TaskUpdateRequest{Task: codeTask(&t)})
	if err != nil {
		return err
	}

	return nil
}

func (s storage) Get(ctx context.Context, ID uint) (*models.Task, error) {
	resp, err := s.client.TaskGet(ctx, &pb.TaskGetRequest{ID: uint64(ID)})
	if err != nil {
		return nil, err
	}

	decoded, err := decodeTask(resp.GetTask())
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func codeTask(in *models.Task) *pb.Task {
	task := pb.Task{
		ID:      uint64(in.ID),
		Title:   in.Title,
		Created: in.Created.Unix(),
		Updated: in.Edited.Unix(),
	}

	if in.Description != "" {
		task.Description = &in.Description
	}

	return &task
}

func decodeTask(in *pb.Task) (*models.Task, error) {
	if in == nil {
		return nil, errors.New("task_decoding: empty data")
	}

	return &models.Task{
		ID:          uint(in.GetID()),
		Title:       in.GetTitle(),
		Description: in.GetDescription(),
		Created:     time.Unix(in.GetCreated(), 0),
		Edited:      time.Unix(in.GetUpdated(), 0),
	}, nil
}
