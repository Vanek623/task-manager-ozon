package client

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	clientPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/client"
	"log"
	"sync"

	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
)

// RunClients запуск клиентов
func RunClients(ctx context.Context, count int) {
	var wg sync.WaitGroup
	wg.Add(count)
	for i := 0; i < count; i++ {
		func(ID uint) {
			defer wg.Done()
			Run(ctx, ID)
		}(uint(i + 1))
	}

	wg.Wait()
}

const (
	address = "localhost:8081"
)

// Run запускает клиента
func Run(ctx context.Context, ID uint) {
	client, err := clientPkg.New(address, ID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d: client started", ID)

	{
		d := "Some description"
		resp, err := client.TaskCreate(ctx, &pb.TaskCreateRequest{
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
		_, err := client.TaskCreate(ctx, &pb.TaskCreateRequest{
			Title: fmt.Sprintf("%d: Second task", ID),
		})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: second task created", ID)
		}
	}
	{
		response, err := client.TaskList(ctx, &pb.TaskListRequest{})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: tasks list: [%v]", ID, response)
		}
	}
	{
		r, err := client.TaskGet(ctx, &pb.TaskGetRequest{ID: 1})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: task got: [%v]", ID, r)
		}
	}
	{
		d := "edited description"
		_, err := client.TaskUpdate(ctx, &pb.TaskUpdateRequest{
			ID:          1,
			Title:       "edited task",
			Description: &d,
		})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: task updated", ID)

			r, err := client.TaskGet(ctx, &pb.TaskGetRequest{ID: 1})
			if err != nil {
				log.Println(errors.Wrapf(err, "[%d]", ID))
			} else {
				log.Printf("%d: task got: [%v]", ID, r)
			}
		}
	}
	{
		_, err := client.TaskDelete(ctx, &pb.TaskDeleteRequest{ID: 1})
		if err != nil {
			log.Println(errors.Wrapf(err, "[%d]", ID))
		} else {
			log.Printf("%d: task deleted", ID)

			r, err := client.TaskList(ctx, &pb.TaskListRequest{})
			if err != nil {
				log.Println(errors.Wrapf(err, "[%d]", ID))
			} else {
				log.Printf("%d: tasks list: [%v]", ID, r)
			}
		}
	}
}
