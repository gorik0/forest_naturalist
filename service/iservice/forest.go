package iservice

import (
	"forest_natura/domain"
)

type ForestService interface {
	IsUnknonwAnimal(string) (bool, error)
	MakeRouteSum(point []*domain.Point) domain.RouteSummary
	MakeRoute(routeName string) (error, []*domain.Point)
}
