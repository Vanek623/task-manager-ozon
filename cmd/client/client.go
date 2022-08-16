package client

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"

	"google.golang.org/grpc/credentials/insecure"

	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
	"google.golang.org/grpc"
)

const reconnectMaxCount = 5
const reconnectTimeout = 2 * time.Second

// RunClients запуск клиентов
func RunClients(ctx context.Context, count int) {
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		func(ID uint) {
			defer wg.Done()
			run(ctx, ID)
		}(uint(i + 1))
	}

	wg.Wait()
}

const (
	address = "localhost:8081"
)

func run(ctx context.Context, ID uint) {
	// Задержка для запуска grpc сервера
	time.Sleep(reconnectTimeout)

	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	for count := 1; err != nil || con == nil; count++ {
		if count > reconnectMaxCount {
			log.Fatalf("%d: cannot connect to server", ID)
		}

		log.Printf("%d: cannot connect to server, try to connect #%d of %d in %d", ID, count, reconnectMaxCount, reconnectTimeout)
		time.Sleep(reconnectTimeout)
		con, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	server := pb.NewServiceClient(con)

	log.Printf("%d: client started", ID)

	{
		d := "Some description"
		resp, err := server.TaskCreate(ctx, &pb.TaskCreateRequest{
			Title:       fmt.Sprintf("%d: First task", ID),
			Description: &d,
		})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: first task created [%d]", ID, resp.GetID())
		}
	}
	{
		_, err := server.TaskCreate(ctx, &pb.TaskCreateRequest{
			Title: fmt.Sprintf("%d: Second task", ID),
		})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: second task created", ID)
		}
	}
	{
		response, err := server.TaskList(ctx, &pb.TaskListRequest{})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: tasks list: [%v]", ID, response)
		}
	}
	{
		r, err := server.TaskGet(ctx, &pb.TaskGetRequest{ID: 1})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: task got: [%v]", ID, r)
		}
	}
	{
		d := "edited description"
		_, err := server.TaskUpdate(ctx, &pb.TaskUpdateRequest{
			ID:          1,
			Title:       "edited task",
			Description: &d,
		})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: task updated", ID)

			r, err := server.TaskGet(ctx, &pb.TaskGetRequest{ID: 1})
			if err != nil {
				log.Println(errors.Wrapf(err, "[%d]", ID))
			} else {
				log.Printf("%d: task got: [%v]", ID, r)
			}
		}
	}
	{
		_, err := server.TaskDelete(ctx, &pb.TaskDeleteRequest{ID: 1})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: task deleted", ID)

			r, err := server.TaskList(ctx, &pb.TaskListRequest{})
			if err != nil {
				log.Println(errors.Wrapf(err, "[%d]", ID))
			} else {
				log.Printf("%d: tasks list: [%v]", ID, r)
			}
		}
	}
}
