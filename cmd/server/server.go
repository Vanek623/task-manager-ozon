package main

import (
	"context"
	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/bot"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc/credentials/insecure"

	"gitlab.ozon.dev/Vanek623/task-manager-system/cmd/config"
	apiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api"
	taskPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api"
	"google.golang.org/grpc"
)

func runGRPC(tm taskPkg.IManager) {
	listener, err := net.Listen(config.ConnectionType, config.FullAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAdminServer(s, apiPkg.New(tm))

	if err = s.Serve(listener); err != nil {
		log.Panic(err)
	}
}

func main() {
	tm := taskPkg.NewLocalManager()
	
	go bot.RunBot(tm)
	go runREST()

	runGRPC(tm)
}

func runREST() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(headerMatcherREST),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(config.FullAddress, mux); err != nil {
		panic(err)
	}
}

func headerMatcherREST(key string) (string, bool) {
	switch key {
	case "Custom":
		return key, true
	default:
		return key, false
	}
}
