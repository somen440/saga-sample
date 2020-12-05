package main

import (
	"context"
	"errors"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/somen440/saga-sample/proto"
	"github.com/somen440/saga-sample/util"
)

var (
	ErrConsumerNotFound = errors.New("service: consumer not found")
	ErrCardNotFound     = errors.New("service: card not found")
	ErrCardExpired      = errors.New("service: card expired")
)

func NewService() *proto.ConsumerService {
	srv := newService()
	return &proto.ConsumerService{
		VerifyConsumer: srv.VerifyConsumer,
	}
}

type service struct {
	r *repository
}

func newService() *service {
	return &service{
		r: defaultRepo,
	}
}

func (srv *service) VerifyConsumer(ctx context.Context, in *proto.VerifyConsumerRequest) (*proto.VerifyConsumerReply, error) {
	grpc_ctxtags.Extract(ctx).Set("consumer.id", in.GetConsumerId())
	l := ctxlogrus.Extract(ctx)
	l.Debug("VerifyConsumer started")

	consumer, err := srv.r.FindOrNil(ctx, in.GetConsumerId())
	if err != nil {
		return nil, err
	}
	if consumer == nil {
		return nil, ErrConsumerNotFound
	}
	if consumer.Card == nil {
		return nil, ErrCardNotFound
	}
	if !util.IsValidTimestamp(consumer.Card.ExpirationDate) {
		return nil, ErrCardExpired
	}

	time.Sleep(500 * time.Millisecond)

	l.Debug("VerifyConsumer success")

	return &proto.VerifyConsumerReply{Consumer: consumer}, nil
}
