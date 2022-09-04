package client

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/service/models"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const reconnectMaxCount = 5
const timeout = 2 * time.Second

type iService interface {
	AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error)
	DeleteTask(ctx context.Context, data *models.DeleteTaskData) error
	TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error)
	UpdateTask(ctx context.Context, data *models.UpdateTaskData) error
	GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error)
}

// ServiceClient структура клиента сервиса
type ServiceClient struct {
	iService
	client pb.ServiceClient
	id     uint
}

func (c *ServiceClient) AddTask(ctx context.Context, data *models.AddTaskData) (*uuid.UUID, error) {
	req := &pb.TaskCreateRequest{
		Title:       data.Title(),
		Description: nil,
	}

	resp, err := c.client.TaskCreate(ctx, req)
	if err != nil {
		return nil, err
	}

	id, err := uuid.FromBytes(resp.GetID())
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (c *ServiceClient) DeleteTask(ctx context.Context, data *models.DeleteTaskData) error {
	bytes, err := data.ID().MarshalBinary()
	if err != nil {
		return err
	}

	req := &pb.TaskDeleteRequest{ID: bytes}

	_, err = c.client.TaskDelete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *ServiceClient) TasksList(ctx context.Context, data *models.ListTaskData) ([]*models.Task, error) {
	req := &pb.TaskListRequest{
		Limit:  data.Limit(),
		Offset: data.Offset(),
	}

	resp, err := c.client.TaskList(ctx, req)
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	for _, t := range resp.GetTasks() {
		var id uuid.UUID
		id, err = uuid.FromBytes(t.GetID())
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, models.NewTask(&id, t.GetTitle()))
	}

	return tasks, nil
}

func (c *ServiceClient) UpdateTask(ctx context.Context, data *models.UpdateTaskData) error {
	id, err := data.ID().MarshalBinary()
	if err != nil {
		return err
	}

	req := &pb.TaskUpdateRequest{
		ID:    id,
		Title: data.Title(),
	}

	if str := data.Description(); str != "" {
		req.Description = &str
	}

	if _, err = c.client.TaskUpdate(ctx, req); err != nil {
		return err
	}

	return nil
}

func (c *ServiceClient) GetTask(ctx context.Context, data *models.GetTaskData) (*models.DetailedTask, error) {
	id, err := data.ID().MarshalBinary()
	if err != nil {
		return nil, err
	}

	req := &pb.TaskGetRequest{
		ID: id,
	}

	resp, err := c.client.TaskGet(ctx, req)
	if err != nil {
		return nil, err
	}

	return models.NewDetailedTask(
		resp.GetTitle(),
		resp.GetDescription(),
		time.Unix(resp.GetEdited(), 0),
	), nil
}

// New создание нового клиента
func New(address string, ID uint) (*ServiceClient, error) {
	c := &ServiceClient{id: ID}

	if err := c.connect(address); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *ServiceClient) makeConnection(address string) (*grpc.ClientConn, error) {
	ctx, cl := context.WithTimeout(context.Background(), timeout)
	defer cl()

	con, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, errors.Wrapf(err, "client #%d: cannot connect to server", c.id)
	}

	return con, err
}

func (c *ServiceClient) connect(address string) error {
	for i := 1; i <= reconnectMaxCount; i++ {
		con, err := c.makeConnection(address)

		if err != nil && i != reconnectMaxCount {
			log.Println(errors.Wrapf(err, "try to connect #%d of %d in %.2f s",
				i, reconnectMaxCount, timeout.Seconds()))
			time.Sleep(timeout)
			continue
		}

		if err != nil && i == reconnectMaxCount {
			return err
		}

		c.client = pb.NewServiceClient(con)
		break
	}

	return nil
}
