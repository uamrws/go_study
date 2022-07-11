package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	trippb "go_study/coolcar_study/server/proto/gen/go"
	trip "go_study/coolcar_study/server/tripservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("failed to listen: %+v", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServerServer(s, &trip.Service{})
	log.Fatal(s.Serve(lis))
}
func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				protojson.MarshalOptions{UseEnumNumbers: true},
				protojson.UnmarshalOptions{},
			},
		),
	)
	err := trippb.RegisterTripServerHandlerFromEndpoint(
		c,
		mux,
		":8090",
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatalf("failed to start grpc gateway: %+v", err)
	}
	err = http.ListenAndServe(":8091", mux)
	if err != nil {
		log.Fatalf("failed to start grpc gateway: %+v", err)
	}
}
