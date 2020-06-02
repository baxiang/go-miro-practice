package grpcclients

import (
	"github.com/pkg/errors"
	"github.com/baxiang/go-miro-practice/api/proto"
	"github.com/baxiang/go-miro-practice/internal/pkg/transports/grpc"
)

func NewDetailsClient(client *grpc.Client) (proto.DetailsClient, error) {
	conn,err := client.Dial("Details")
	if err != nil {
		return nil,errors.Wrap(err,"detail client dial error")
	}
	c := proto.NewDetailsClient(conn)

	return c, nil
}