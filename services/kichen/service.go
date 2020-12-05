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
	ErrNotFound = errors.New("service: ticket not found")
)

func NewService() *proto.KitchenService {
	srv := newService()
	return &proto.KitchenService{
		CreateTicket:  srv.CreateTicket,
		ApproveTicket: srv.ApproveTicket,
		RejectTicket:  srv.RejectTicket,
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

func (srv *service) CreateTicket(ctx context.Context, in *proto.CreateTicketRequest) (*proto.CreateTicketReply, error) {
	orderID := in.GetOrderId()

	grpc_ctxtags.Extract(ctx).Set("order.id", orderID)
	l := ctxlogrus.Extract(ctx)
	l.Debug("CreateTicket started")

	ticket, err := srv.r.Create(ctx, orderID)
	if err != nil {
		return nil, err
	}

	return &proto.CreateTicketReply{Ticket: ticket}, nil
}

func (srv *service) ApproveTicket(ctx context.Context, in *proto.ApproveTicketRequest) (*proto.ApproveTicketReply, error) {
	ticketID := in.GetTicketId()

	grpc_ctxtags.Extract(ctx).Set("ticket.id", ticketID)
	l := ctxlogrus.Extract(ctx)
	l.Debug("ApproveTicket started")

	ticket, err := srv.r.FindOrNil(ctx, ticketID)
	if err != nil {
		return nil, err
	}
	if ticket == nil {
		return nil, ErrNotFound
	}
	if err := srv.r.UpdateStatus(ctx, ticket, proto.TicketStatus_TICKET_APPROVED); err != nil {
		return nil, err
	}

	time.Sleep(500 * time.Millisecond)

	l.Debug("ApproveTicket success")

	return &proto.ApproveTicketReply{}, nil
}

func (srv *service) RejectTicket(ctx context.Context, in *proto.RejectTicketRequest) (*proto.RejectTicketReply, error) {
	ticketID := in.GetTicketId()

	grpc_ctxtags.Extract(ctx).Set("ticket.id", ticketID)
	l := ctxlogrus.Extract(ctx)
	l.Debug("RejectTicket started")

	ticket, err := srv.r.FindOrNil(ctx, ticketID)
	if err != nil {
		return nil, err
	}
	if ticket == nil {
		return nil, ErrNotFound
	}
	if err := srv.r.UpdateStatus(ctx, ticket, proto.TicketStatus_TICKET_CREATE_REJECTED); err != nil {
		return nil, err
	}

	return &proto.RejectTicketReply{}, nil
}
