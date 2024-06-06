package forestroute

import (
	"context"
	"fmt"
	"forest_natura/domain"
	"forest_natura/forestroute"
	"forest_natura/service/iservice"
	"io"
)

type ForestGrpc struct {
	forestroute.UnimplementedForestRouteServer
	service iservice.ForestService
}

func (f *ForestGrpc) MakeRouteSummary(stream forestroute.ForestRoute_MakeRouteSummaryServer) error {

	var points []*domain.Point

	for {
		point, err := stream.Recv()
		if err != nil {
			if err == io.EOF {

				summaryFromService := f.service.MakeRouteSum(points)
				err := stream.SendAndClose(&forestroute.RouteSummary{
					StepsCount: int32(summaryFromService.StepsCount),
					Duration:   nil,
				})
				if err != nil {
					return fmt.Errorf("Error while service request:::%w", err)
				}
				return nil

			}
			return fmt.Errorf("Error while recienig point from client :::: %w", err)
		}
		points = append(points, &domain.Point{Position: int(point.Position)})

	}
}
func (f *ForestGrpc) MakeRoute(routeName *forestroute.Route, stream forestroute.ForestRoute_MakeRouteServer) error {

	err, points := f.service.MakeRoute(routeName.Name)
	fmt.Println("points", points)
	if err != nil {
		return fmt.Errorf("Error while  getting poiunts from service for routename %s :::: %w", routeName, err)
	}
	for _, point := range points {
		err := stream.Send(&forestroute.Point{Position: int32(point.Position)})
		if err != nil {
			return fmt.Errorf("error while sending point to client ::: %w", err)
		}

	}
	return nil
}
func (f *ForestGrpc) RegisterAnimal(c context.Context, animal *forestroute.Animal) (*forestroute.IsAnimalUnknown, error) {
	if unknown, err := f.service.IsUnknonwAnimal(animal.Name); err != nil {

		return &forestroute.IsAnimalUnknown{Known: false}, err
	} else {
		return &forestroute.IsAnimalUnknown{Known: unknown}, nil
	}

}
