package main

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/somen440/saga-sample/proto"
)

type repository struct {
	consumerEntityMap map[int64]*proto.ConsumerEntity
}

var defaultRepo *repository

func init() {
	defaultRepo = &repository{
		consumerEntityMap: mockConsumers(),
	}
}

func (r *repository) FindOrNil(ctx context.Context, id int64) (*proto.ConsumerEntity, error) {
	consumer, ok := r.consumerEntityMap[id]
	if !ok {
		return nil, nil
	}
	return consumer, nil
}

func mockConsumers() map[int64]*proto.ConsumerEntity {
	endTime, _ := ptypes.TimestampProto(time.Date(2999, 12, 31, 23, 59, 59, 0, time.UTC))

	return map[int64]*proto.ConsumerEntity{
		// 有効な消費者
		1: {
			Id: 1,
			Card: &proto.CardEntity{
				Id:               1,
				PaymentServiceId: "psi_1111",
				ExpirationDate:   endTime,
			},
		},
		// カードを持っていなくて無効な消費者
		2: {
			Id:   2,
			Card: nil,
		},
		// カードの有効期限が切れていて無効な消費者
		3: {
			Id: 3,
			Card: &proto.CardEntity{
				Id:               2,
				PaymentServiceId: "psi_3232",
			},
		},
	}
}
