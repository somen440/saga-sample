package main

import (
	"context"
	"errors"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/somen440/saga-sample/proto"
)

var (
	ErrNotFound = errors.New("service: order not found")
)

func NewService() *proto.OrderService {
	srv := newService()
	return &proto.OrderService{
		GetOrder:     srv.GetOrder,
		CreateOrder:  srv.CreateOrder,
		ApproveOrder: srv.ApproveOrder,
		RejectOrder:  srv.RejectOrder,
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

func (srv *service) GetOrder(ctx context.Context, in *proto.GetOrderRequest) (*proto.GetOrderReply, error) {
	grpc_ctxtags.Extract(ctx).
		Set("order.id", in.GetOrderId())

	l := ctxlogrus.Extract(ctx)
	l.Debug("GetOrderfunc started")

	order, err := srv.r.FindOrNil(ctx, in.GetOrderId())
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, ErrNotFound
	}

	l.Debug("GetOrderfunc success")

	return &proto.GetOrderReply{Order: order}, nil
}

func (srv *service) CreateOrder(ctx context.Context, in *proto.CreateOrderRequest) (*proto.CreateOrderReply, error) {
	itemID := in.GetItemId()
	consumerID := in.GetConsumerId()

	grpc_ctxtags.Extract(ctx).
		Set("consumer.id", itemID).
		Set("item.id", consumerID)

	l := ctxlogrus.Extract(ctx)
	l.Debug("CreateOrder started")

	order := &proto.OrderEntity{
		ItemId:     itemID,
		ConsumerId: consumerID,
	}
	order, err := srv.r.Create(ctx, itemID, consumerID)
	if err != nil {
		return nil, err
	}

	PublishCreateOrderSaga(&proto.CreateOrderSagaRequest{OrderId: order.Id, ConsumerId: consumerID})

	l.WithField("order.id", order.Id).Debug("CreateOrder success")

	return &proto.CreateOrderReply{Order: order}, nil
}

func (srv *service) ApproveOrder(ctx context.Context, in *proto.ApproveOrderRequest) (*proto.ApproveOrderReply, error) {
	orderID := in.GetOrderId()

	grpc_ctxtags.Extract(ctx).Set("order.id", orderID)
	l := ctxlogrus.Extract(ctx)
	l.Debug("ApproveOrder started")

	order, err := srv.r.FindOrNil(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, ErrNotFound
	}
	if err := srv.r.UpdateStatus(ctx, order, proto.OrderStatus_ORDER_APPROVED); err != nil {
		return nil, err
	}

	time.Sleep(500 * time.Millisecond)

	l.Debug("ApproveOrder success")

	return &proto.ApproveOrderReply{}, nil
}

func (srv *service) RejectOrder(ctx context.Context, in *proto.RejectOrderRequest) (*proto.RejectOrderReply, error) {
	orderID := in.GetOrderId()

	grpc_ctxtags.Extract(ctx).Set("order.id", orderID)
	l := ctxlogrus.Extract(ctx)
	l.Debug("ApproveOrder started")

	order, err := srv.r.FindOrNil(ctx, orderID)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, ErrNotFound
	}
	if err := srv.r.UpdateStatus(ctx, order, proto.OrderStatus_ORDER_CREATE_REJECTED); err != nil {
		return nil, err
	}

	return &proto.RejectOrderReply{}, nil
}
