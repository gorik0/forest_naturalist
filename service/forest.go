package service

import (
	"fmt"
	"forest_natura/domain"
	generator2 "forest_natura/generator"
	"forest_natura/service/iservice"
	"slices"
	"strings"
)

type ForestUsecase struct {
	Animals []string
}

func (f *ForestUsecase) MakeRoute(point []*domain.Point) domain.RouteSummary {
	return domain.RouteSummary{

		StepsCount: len(point),
	}

}

func (f *ForestUsecase) MakeRouteSummary(routeName string) (error, []*domain.Point) {
	var points []*domain.Point
	if contain := slices.Contains(generator2.AvailableRoute, routeName); !contain {
		return fmt.Errorf("route is unknown!!!"), points
	}
	return nil, generator2.GenerateRandomPOintsForRoute()
}

func MakeAnimapls() []string {
	return strings.Split("elephant pantera tiger hyppo beaver", " ")
}
func (f *ForestUsecase) IsUnknonwAnimal(animal string) (bool, error) {

	return slices.Contains(f.Animals, animal), nil
}

var _ iservice.ForestService = &ForestUsecase{}
