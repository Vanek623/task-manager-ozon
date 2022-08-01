package client

import (
	"context"
	"log"
	"time"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/config"

	"google.golang.org/grpc/credentials/insecure"

	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api"
	"google.golang.org/grpc"
)

const reconnectMaxCount = 5
const reconnectTimeout = 2 * time.Second

// Run запустить клиента
func Run() {
	// Задержка для запуска grpc сервера
	time.Sleep(reconnectTimeout)

	con, err := grpc.Dial(config.FullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	for count := 1; err != nil || con == nil; count++ {
		if count > reconnectMaxCount {
			log.Fatal("cannot connect to server")
		}

		log.Printf("cannot connect to server, try to connect #%d of %d in %d", count, reconnectMaxCount, reconnectTimeout)
		time.Sleep(reconnectTimeout)
		con, err = grpc.Dial(config.FullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	server := pb.NewAdminClient(con)

	log.Println("client started")

	ctx := context.Background()
	{
		d := "Some description"
		_, err := server.TaskCreate(ctx, &pb.TaskCreateRequest{
			Title:       "First task",
			Description: &d,
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("first task created")
		}
	}
	{
		_, err := server.TaskCreate(ctx, &pb.TaskCreateRequest{
			Title: "Second task",
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("second task created")
		}
	}
	{
		response, err := server.TaskList(ctx, &pb.TaskListRequest{})
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("tasks list: [%v]", response)
		}
	}
	{
		r, err := server.TaskGet(ctx, &pb.TaskGetRequest{ID: 1})
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("task got: [%v]", r)
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
			log.Println(err)
		} else {
			log.Println("task updated")

			r, err := server.TaskGet(ctx, &pb.TaskGetRequest{ID: 1})
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("task got: [%v]", r)
			}
		}
	}
	{
		_, err := server.TaskDelete(ctx, &pb.TaskDeleteRequest{ID: 1})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("task deleted")

			r, err := server.TaskList(ctx, &pb.TaskListRequest{})
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("tasks list: [%v]", r)
			}
		}
	}
}
