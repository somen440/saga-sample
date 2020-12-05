package main

import (
	"context"
	"sync"

	"github.com/somen440/saga-sample/proto"
)

type repository struct {
	maxID int64

	orderEntityMap map[int64]*proto.OrderEntity

	itemIDMap map[int64]bool

	cond *sync.Cond
}

var defaultRepo *repository

func init() {
	defaultRepo = &repository{
		maxID:          0,
		orderEntityMap: map[int64]*proto.OrderEntity{},
		itemIDMap:      map[int64]bool{},
		cond:           sync.NewCond(&sync.Mutex{}),
	}
}

func (r *repository) Exists(ctx context.Context, id int64) bool {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	_, ok := r.orderEntityMap[id]
	return ok
}

func (r *repository) ExistsItemID(ctx context.Context, itemID int64) bool {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	_, ok := r.itemIDMap[itemID]
	return ok
}

func (r *repository) FindOrNil(ctx context.Context, id int64) (*proto.OrderEntity, error) {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	order, ok := r.orderEntityMap[id]
	if !ok {
		return nil, nil
	}

	return order, nil
}

func (r *repository) Create(ctx context.Context, itemID, consumerID int64) (*proto.OrderEntity, error) {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	order := &proto.OrderEntity{
		Id:         r.maxID + 1,
		ItemId:     itemID,
		ConsumerId: consumerID,
		Status:     proto.OrderStatus_ORDER_CREATED,
	}

	r.orderEntityMap[order.Id] = order
	r.itemIDMap[order.ItemId] = true
	r.maxID++

	return order, nil
}

func (r *repository) UpdateStatus(ctx context.Context, order *proto.OrderEntity, status proto.OrderStatus) error {
	r.cond.L.Lock()
	defer r.cond.L.Unlock()

	order.Status = status
	r.orderEntityMap[order.Id] = order

	return nil
}
