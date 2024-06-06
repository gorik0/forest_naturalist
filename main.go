package main

import (
	"context"
	forestroute "forest_natura/grpc"
	service2 "forest_natura/service"
)

func main() {
	println("Helo gorik!")
	service := service2.NewForestService()
	s := forestroute.GrpcServer{
		ForestService: service,
	}
	ctx := context.Background()
	s.Run(ctx, ":9000")

}
