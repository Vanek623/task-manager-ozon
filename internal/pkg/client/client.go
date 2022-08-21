package client

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	pb "gitlab.ozon.dev/Vanek623/task-manager-system/pkg/api/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const reconnectMaxCount = 5
const timeout = 2 * time.Second

// Client структура клиента
type Client struct {
	pb.ServiceClient
	id uint
}

// New создание нового клиента
func New(address string, ID uint) (*Client, error) {
	c := &Client{id: ID}

	if err := c.connect(address); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) makeConnection(address string) (*grpc.ClientConn, error) {
	ctx, cl := context.WithTimeout(context.Background(), timeout)
	defer cl()

	con, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, errors.Wrapf(err, "client #%d: cannot connect to server", c.id)
	}

	return con, err
}

func (c *Client) connect(address string) error {
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

		c.ServiceClient = pb.NewServiceClient(con)
		break
	}

	return nil
}
