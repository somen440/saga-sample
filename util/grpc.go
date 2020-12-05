package util

import (
	"context"
	"os"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/sirupsen/logrus"
)

type ErrorMap map[error]codes.Code

func NewServer(register func(*grpc.Server), isDebuggable bool, errorMap ErrorMap) *grpc.Server {
	logrusLogger := logrus.New()
	if isDebuggable {
		logrusLogger.Level = logrus.DebugLevel
	}
	logrusLogger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
		TimestampFormat: time.RFC3339,
	}
	logrusLogger.Out = os.Stdout
	codeToLevel := grpc_logrus.DefaultCodeToLevel

	logrusEntry := logrus.NewEntry(logrusLogger)
	logrusOpts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(codeToLevel),
	}
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, logrusOpts...),
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				res, err := handler(ctx, req)
				if err == nil {
					return res, nil
				}
				code, ok := errorMap[err]
				if !ok {
					return nil, status.Error(codes.Unknown, err.Error())
				}
				return nil, status.Error(code, err.Error())
			},
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry, logrusOpts...),
		),
	)

	register(grpcServer)

	return grpcServer
}
