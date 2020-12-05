package main

import (
	"context"
	"sync"

	"github.com/somen440/saga-sample/proto"
)

type repository struct {
	maxID int64

	ticketEntityMap map[int64]*proto.TicketEntity

	cond *sync.Cond
}

var defaultRepo *repository

func init() {
	defaultRepo = &repository{
		maxID:           0,
		ticketEntityMap: map[int64]*proto.TicketEntity{},
		cond:            sync.NewCond(&sync.Mutex{}),
	}
}

func (r *repository) FindOrNil(ctx context.Context, id int64) (*proto.TicketEntity, error) {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	ticket, ok := r.ticketEntityMap[id]
	if !ok {
		return nil, nil
	}

	return ticket, nil
}

func (r *repository) Create(ctx context.Context, orderID int64) (*proto.TicketEntity, error) {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	ticket := &proto.TicketEntity{
		Id:      r.maxID + 1,
		OrderId: orderID,
		Status:  proto.TicketStatus_TICKET_CREATED,
	}

	r.ticketEntityMap[ticket.Id] = ticket
	r.maxID++

	return ticket, nil
}

func (r *repository) UpdateStatus(ctx context.Context, ticket *proto.TicketEntity, status proto.TicketStatus) error {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	ticket.Status = status
	r.ticketEntityMap[ticket.Id] = ticket

	return nil
}
