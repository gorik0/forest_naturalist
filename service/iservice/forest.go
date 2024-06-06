package iservice

import (
	"forest_natura/domain"
)

type ForestService interface {
	IsUnknonwAnimal(string) (bool, error)
	MakeRoute(point []*domain.Point) domain.RouteSummary
	MakeRouteSummary(routeName string) (error, []*domain.Point)
}
