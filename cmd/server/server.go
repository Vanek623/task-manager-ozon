package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/pkg/core/task/models"

	"gitlab.ozon.dev/Vanek623/task-manager-system/internal/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc/credentials/insecure"

	apiPkg "gitlab.ozon.dev/Vanek623/task-manager-system/internal/api"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api"
	"google.golang.org/grpc"
)

type iTaskStorage interface {
	Add(ctx context.Context, t models.Task) error
	Delete(ctx context.Context, ID uint) error
	List(ctx context.Context) ([]models.Task, error)
	Update(ctx context.Context, t models.Task) error
	Get(ctx context.Context, ID uint) (*models.Task, error)
}

// RunGRPC запускает GRPC
func RunGRPC(tm iTaskStorage) {
	listener, err := net.Listen(config.ConnectionType, config.FullAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterAdminServer(s, apiPkg.New(tm))

	log.Println("grpc started")

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

// RunREST запускает REST
func RunREST() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, config.FullAddress, opts); err != nil {
		log.Fatal(err)
	}

	log.Println("rest started")

	if err := http.ListenAndServe(config.FullHTTPAddress, mux); err != nil {
		log.Fatal(err)
	}
}
