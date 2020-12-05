package main

import (
	"context"
	"errors"
	"log"

	"github.com/somen440/saga-sample/proto"
	"google.golang.org/grpc"
	"k8s.io/client-go/util/workqueue"
)

const (
	sagaAddr = ":5010"
)

// Worker 処理制限を行う
type Worker struct {
	queue *workqueue.Type

	sagaConn *grpc.ClientConn
}

var defaultWorker *Worker

var (
	ErrNotImplemented = errors.New("Not Implemented")
)

func InitializeWorker() {
	sagaConn, err := grpc.Dial(sagaAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defaultWorker = &Worker{
		queue:    workqueue.New(),
		sagaConn: sagaConn,
	}
	go defaultWorker.runWorker()
}

func CloseWorker() {
	defaultWorker.sagaConn.Close()
}

func PublishCreateOrderSaga(req *proto.CreateOrderSagaRequest) {
	defaultWorker.queue.Add(req)
	log.Print("PublishCreateOrderSaga")
}

func (w *Worker) processNextItem() bool {
	log.Print("queue wait...")

	q, quit := w.queue.Get()
	if quit {
		return false
	}
	log.Print("queue get")

	switch q.(type) {
	case *proto.CreateOrderSagaRequest:
		req := q.(*proto.CreateOrderSagaRequest)
		defer w.queue.Done(req)

		client := proto.NewSagaClient(w.sagaConn)
		client.CreateOrderSaga(context.TODO(), req)

		log.Print("CreateOrderSagaRequest")
	default:
		panic(ErrNotImplemented)
	}

	log.Print("queue end")

	return true
}

func (w *Worker) runWorker() {
	for w.processNextItem() {
	}
}
