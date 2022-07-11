package main

import (
	"go.uber.org/zap"
	"go_study/coolcar/server/auth"
	authpb "go_study/coolcar/server/auth/api/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 可以通过zap.NewDevelopmentConfig自定义config
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create zap logger %v", err)
	}
	lis, err := net.Listen("tcp", ":8090")

	if err != nil {
		logger.Fatal("failed to listen tcp", zap.Error(err))
	}
	ns := grpc.NewServer()
	authpb.RegisterAuthServiceServer(ns, &auth.Service{Logger: logger})
	logger.Fatal("failed to serve", zap.Error(ns.Serve(lis)))
}
