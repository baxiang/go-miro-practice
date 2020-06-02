package grpcservers

import (
	"github.com/baxiang/go-miro-practice/api/proto"
	"github.com/baxiang/go-miro-practice/internal/app/ratings/services"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
	"context"
	"github.com/pkg/errors"
)

type RatingsServer struct {
	logger  *zap.Logger
	service services.RatingsService
}


func NewRatingsServer(logger *zap.Logger, ps services.RatingsService) (*RatingsServer, error) {
	return &RatingsServer{
		logger:  logger,
		service: ps,
	}, nil
}

func (s *RatingsServer) Get(ctx context.Context, req *proto.GetRatingRequest) (*proto.Rating, error) {
	r, err := s.service.Get(req.ProductID)
	if err != nil {
		return nil, errors.Wrap(err, "product grpc service get rating error")
	}

	ut, err := ptypes.TimestampProto(r.UpdatedTime)
	if err != nil {
		return nil, errors.Wrap(err, "convert create time error")
	}

	resp := &proto.Rating{
		Id:          uint64(r.ID),
		ProductID:   r.ProductID,
		Score:       r.Score,
		UpdatedTime: ut,
	}

	return resp, nil
}