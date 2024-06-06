package forestroute

import (
	"context"
	"fmt"
	"forest_natura/forestroute"
	"forest_natura/service/iservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GrpcServer struct {
	ForestService iservice.ForestService
	s             *grpc.Server
}

func (g *GrpcServer) Run(ctx context.Context, addr string) {
	var listenCfg net.ListenConfig

	lis, err := listenCfg.Listen(ctx, "tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("error while creating listen for grpc server ::: %d", err))
	}

	g.s = grpc.NewServer()
	reflection.Register(g.s)
	forestroute.RegisterForestRouteServer(g.s, &ForestGrpc{
		service: g.ForestService,
	})
	err = g.s.Serve(lis)
	if err != nil {
		panic(fmt.Sprintf("Error while serving server ::: $d", err))
	}

}
