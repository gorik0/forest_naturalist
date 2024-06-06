package forestroute

import (
	"context"
	"forest_natura/forestroute"
	"forest_natura/service/iservice"
)

type ForestGrpc struct {
	forestroute.UnimplementedForestRouteServer
	service iservice.ForestService
}

func (f *ForestGrpc) MakeRouteSummary(forestroute.ForestRoute_MakeRouteSummaryServer) error {

	return nil
}
func (f *ForestGrpc) MakeRoute(*forestroute.Route, forestroute.ForestRoute_MakeRouteServer) error {
	return nil
}
func (f *ForestGrpc) RegisterAnimal(c context.Context, animal *forestroute.Animal) (*forestroute.IsAnimalUnknown, error) {
	if unknown, err := f.service.IsUnknonwAnimal(animal.Name); err != nil {

		return &forestroute.IsAnimalUnknown{Known: false}, err
	} else {
		return &forestroute.IsAnimalUnknown{Known: unknown}, nil
	}

}
