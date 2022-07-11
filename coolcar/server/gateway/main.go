package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	authpb "go_study/coolcar/server/auth/api/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create logger %v", err)
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{UseEnumNumbers: true}},
		),
	)
	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err = authpb.RegisterAuthServiceHandlerFromEndpoint(
		ctx, mux, ":8090", options,
	)
	if err != nil {
		logger.Fatal("failed to register", zap.Error(err))
	}
	logger.Fatal("failed to listen port", zap.Error(http.ListenAndServe(":8091", mux)))
}
