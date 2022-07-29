package main

import (
	"log"
	"net"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/config"
	apiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api"
	taskPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen(config.ConnectionType, config.FullAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAdminServer(s, apiPkg.New(taskPkg.NewLocalManager()))

	if err = s.Serve(listener); err != nil {
		log.Panic(err)
	}
}
