package client

import (
	"github.com/pkg/errors"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const reconnectMaxCount = 5
const reconnectTimeout = 2 * time.Second

type Client struct {
	pb.ServiceClient
	id uint
}

func New(address string, ID uint) (*Client, error) {
	con, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	for count := 1; err != nil || con == nil; count++ {
		if count > reconnectMaxCount {
			return nil, errors.Errorf("client #%d: cannot connect to server", ID)
		}

		log.Printf("client #%d: cannot connect to server, try to connect #%d of %d in %d",
			ID, count, reconnectMaxCount, reconnectTimeout)
		time.Sleep(reconnectTimeout)
		con, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &Client{
		pb.NewServiceClient(con),
		ID,
	}, nil
}
