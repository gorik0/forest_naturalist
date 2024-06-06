package generator

import (
	"forest_natura/domain"
	"math/rand"
	"strings"
	"time"
)

var AvailableRoute = strings.Split("forest river town field swamp", " ")

func GenerateRandomPOintsForRoute() (p []*domain.Point) {

	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := 0; i < r.Intn(10); i++ {
		p = append(p, &domain.Point{i})
	}
	return p
}
