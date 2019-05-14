package main

import (
	"golang.org/x/net/context"
)

const (
	grpcPort = ":50051"
	httpPort = ":8080"
)

func main() {
	ctx := context.Background()
	repo := &service.Repository{}
	svc := service.NewService(repo)

	go func() {
		if err := grpc.RunGrpcServer(ctx, svc, grpcPort); err != nil {
			glog.Fatal(err)
		}
	}()

	if err := rest.RunRestServer(ctx, grpcPort, httpPort); err != nil {
		glog.Fatal(err)
	}

}
