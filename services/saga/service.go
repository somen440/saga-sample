package main

import (
	"context"
	"errors"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/somen440/saga-sample/proto"
	"google.golang.org/grpc"
)

const (
	consumerAddr = ":5000"
	kichenAddr   = ":5001"
	orderAddr    = ":5002"
	treasureAddr = ":5003"
)

var (
	ErrNotFound = errors.New("service: order not found")
)

func NewService() *proto.SagaService {
	srv := newService()
	return &proto.SagaService{
		CreateOrderSaga: srv.CreateOrderSaga,
	}
}

type service struct {
}

func newService() *service {
	return &service{}
}

type CreateOrderSagaState struct {
	OrderID          int64
	ConsumerID       int64
	PaymentServiceID string
	TicketID         int64
}

type Step struct {
	invokeParticipant func() (interface{}, error)
	onReply           func(reply interface{})
	withCompensation  func() error
}

func (srv *service) CreateOrderSaga(ctx context.Context, in *proto.CreateOrderSagaRequest) (*proto.CreateOrderSagaReply, error) {
	orderID := in.OrderId
	consumerID := in.ConsumerId

	grpc_ctxtags.Extract(ctx).
		Set("order.id", orderID).
		Set("consumer.id", consumerID)

	l := ctxlogrus.Extract(ctx)
	l.Debug("CreateOrderSaga started")

	consumerConn, err := grpc.DialContext(ctx, consumerAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		return nil, err
	}
	defer consumerConn.Close()
	consumerClient := proto.NewConsumerClient(consumerConn)

	orderConn, err := grpc.DialContext(ctx, orderAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		return nil, err
	}
	defer orderConn.Close()
	orderClient := proto.NewOrderClient(orderConn)

	kitchenConn, err := grpc.DialContext(ctx, kichenAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		return nil, err
	}
	defer kitchenConn.Close()
	kitchenClient := proto.NewKitchenClient(kitchenConn)

	treasureConn, err := grpc.DialContext(ctx, treasureAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		return nil, err
	}
	defer treasureConn.Close()
	treasureClient := proto.NewTreasurerClient(treasureConn)

	state := &CreateOrderSagaState{
		OrderID:    orderID,
		ConsumerID: consumerID,
	}

	steps := []*Step{}
	steps = append(steps,
		&Step{
			withCompensation: func() error {
				_, err = orderClient.RejectOrder(ctx, &proto.RejectOrderRequest{OrderId: state.OrderID})
				return err
			},
		},
		&Step{
			invokeParticipant: func() (interface{}, error) {
				return consumerClient.VerifyConsumer(ctx, &proto.VerifyConsumerRequest{ConsumerId: consumerID})
			},
			onReply: func(reply interface{}) {
				state.PaymentServiceID = reply.(*proto.VerifyConsumerReply).Consumer.Card.PaymentServiceId
			},
		},
		&Step{
			invokeParticipant: func() (interface{}, error) {
				return kitchenClient.CreateTicket(ctx, &proto.CreateTicketRequest{OrderId: state.OrderID})
			},
			onReply: func(reply interface{}) {
				state.TicketID = reply.(*proto.CreateTicketReply).Ticket.Id
			},
			withCompensation: func() error {
				_, err = kitchenClient.RejectTicket(ctx, &proto.RejectTicketRequest{TicketId: state.TicketID})
				return err
			},
		},
		&Step{
			invokeParticipant: func() (interface{}, error) {
				return treasureClient.AuthorizeCard(ctx, &proto.AuthorizeCardRequest{PaymentServiceId: state.PaymentServiceID})
			},
		},
		&Step{
			invokeParticipant: func() (interface{}, error) {
				return orderClient.ApproveOrder(ctx, &proto.ApproveOrderRequest{OrderId: state.OrderID})
			},
		},
	)

	failedIndex := -1
	for i, step := range steps {
		if step.invokeParticipant == nil {
			continue
		}
		reply, err := step.invokeParticipant()
		if err != nil {
			failedIndex = i
			break
		}
		if step.onReply == nil {
			continue
		}
		step.onReply(reply)
	}

	for i := failedIndex; i >= 0; i-- {
		if steps[i].withCompensation == nil {
			continue
		}
		if err := steps[i].withCompensation(); err != nil {
			return nil, err
		}
	}

	l.Debug("CreateOrderSaga ended")
	return &proto.CreateOrderSagaReply{}, nil
}
