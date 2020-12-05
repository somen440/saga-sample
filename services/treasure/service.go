package main

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/somen440/saga-sample/proto"
)

func NewService() *proto.TreasurerService {
	return &proto.TreasurerService{
		AuthorizeCard: AuthorizeCard,
	}
}

func AuthorizeCard(ctx context.Context, in *proto.AuthorizeCardRequest) (*proto.AuthorizeCardReply, error) {
	grpc_ctxtags.Extract(ctx).Set("payment_service.id", in.GetPaymentServiceId())
	l := ctxlogrus.Extract(ctx)
	l.Debug("AuthorizeCard started")

	// 外部サービスにカード認証確認投げてる想定

	time.Sleep(500 * time.Millisecond)

	l.Debug("AuthorizeCard success")

	return &proto.AuthorizeCardReply{Ok: true}, nil
}
