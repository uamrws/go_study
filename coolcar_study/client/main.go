package main

import (
	"context"
	"fmt"
	trippb "go_study/coolcar_study/server/proto/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	con, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect: %+v\n", err)
	}
	tsClient := trippb.NewTripServerClient(con)
	r, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "trip456"})
	if err != nil {
		log.Fatalf("failed to get trip: %+v\n", err)
	}
	fmt.Printf("get trip successed: %T\n", r)
	fmt.Println(r)
}
