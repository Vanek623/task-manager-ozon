package main

import (
	"context"
	"log"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/config"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial(config.FullAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewAdminClient(con)

	ctx := context.Background()

	{
		d := "Some description"
		_, err := c.TaskCreate(ctx, &pb.TaskCreateRequest{
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
		_, err := c.TaskCreate(ctx, &pb.TaskCreateRequest{
			Title: "Second task",
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("second task created")
		}
	}
	{
		response, err := c.TaskList(ctx, &pb.TaskListRequest{})
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("tasks list: [%v]", response)
		}
	}
	{
		r, err := c.TaskGet(ctx, &pb.TaskGetRequest{ID: 1})
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("task got: [%v]", r)
		}
	}
	{
		d := "edited description"
		_, err := c.TaskUpdate(ctx, &pb.TaskUpdateRequest{
			ID:          1,
			Title:       "edited task",
			Description: &d,
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("task updated")

			r, err := c.TaskGet(ctx, &pb.TaskGetRequest{ID: 1})
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("task got: [%v]", r)
			}
		}
	}
	{
		_, err := c.TaskDelete(ctx, &pb.TaskDeleteRequest{ID: 1})
		if err != nil {
			log.Println(err)
		} else {
			log.Println("task deleted")

			r, err := c.TaskList(ctx, &pb.TaskListRequest{})
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("tasks list: [%v]", r)
			}
		}
	}
}
