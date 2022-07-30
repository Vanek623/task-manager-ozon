package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc/credentials/insecure"

	apiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api"
	taskPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api"
	"google.golang.org/grpc"
)

// RunGRPC запускает GRPC
func RunGRPC(tm taskPkg.IManager) {
	listener, err := net.Listen(config.ConnectionType, config.FullAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAdminServer(s, apiPkg.New(tm))

	log.Printf("grpc started")

	if err = s.Serve(listener); err != nil {
		log.Panic(err)
	}
}

// RunREST запускает REST
func RunREST() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(headerMatcherREST),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, config.FullAddress, opts); err != nil {
		panic(err)
	}
	log.Printf("rest started")

	if err := http.ListenAndServe(config.FullHTTPAddress, mux); err != nil {
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
