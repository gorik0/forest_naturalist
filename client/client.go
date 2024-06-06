package main

import (
	"context"
	"fmt"
	"forest_natura/domain"
	"forest_natura/forestroute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"time"
)

var routePoints []*domain.Point

func main() {
	ctx := context.Background()

	client := makeGrpcClient()

	var answer string
	for {
		makeUSerInput(answer)

		//::::SWITCH ON ANSWER

		switch answer {
		case "1":
			{
				//:::MAKE Register Animal
				registerAnimal(ctx, client)
				continue
			}
		case "2":
			{
				//:::MAKE Make Route
				makeRoouteSummary(ctx, client)
				continue
			}

		case "3":
			//:::MAKE Make Route
			makeRoute(ctx, client)
			continue
		default:
			println("Exit")
			break
		}

	}

}

func makeUSerInput(answer string) {
	println("Enter your action ::::\n\t 1.) Register Animal \n\t2.)MakeRoute\n\t3.)MakeRouteSummary\n\t\t :::::")
	_, err := fmt.Scan(&answer)
	if err != nil {
		panic("BAD INPUT!!!")
	}
}

func makeGrpcClient() forestroute.ForestRouteClient {
	cc, err := grpc.Dial(":9000", grpc.WithTimeout(time.Second*10), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Sprintf("Error in client creating ::: %d", err))
	}
	client := forestroute.NewForestRouteClient(cc)
	return client
}

func makeRoute(ctx context.Context, client forestroute.ForestRouteClient) {
	var routeName string
	println("Enter your route to geting points for ... ")
	_, err := fmt.Scan(&routeName)
	if err != nil {
		panic("bad input")
	}
	route, err := client.MakeRoute(ctx, &forestroute.Route{
		Name: routeName,
	})
	if err != nil {
		panic(fmt.Errorf("Eror while making route :::%w", err))

	}
	for {
		recv, err := route.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(fmt.Errorf("Eror while getting points for route :::%w", err))
		}
		routePoints = append(routePoints, &domain.Point{Position: int(recv.Position)})

	}

	println(fmt.Sprintf("::::: POINTS FOR ROUTE(count) -->>> %s ::::::", len(routePoints)))
	println("DONE!!!")

}

func makeRoouteSummary(ctx context.Context, client forestroute.ForestRouteClient) {
	route, err := client.MakeRouteSummary(ctx)
	if err != nil {
		panic(fmt.Errorf("error while making route:::$w", err))
	}
	for _, p := range routePoints {
		time.Sleep(time.Millisecond * 300)
		err := route.Send(&forestroute.Point{
			Position: int32(p.Position),
		})
		if err != nil {
			panic(fmt.Errorf("Error while sending points ::: %s", err))
		}
	}
	recv, err := route.CloseAndRecv()
	if err != nil {
		panic(fmt.Errorf("Error while getting route summary", err))
	}
	println(fmt.Sprintf("::::: ROUTE SUMMARY -->>> %s ::::::", recv.StepsCount))
	println("DONE!!!")

}

func registerAnimal(ctx context.Context, client forestroute.ForestRouteClient) {
	var animal string
	println("Enter your animal to register for ... ")
	_, err := fmt.Scan(&animal)
	if err != nil {
		panic("bad input")
	}
	unknown, err := client.RegisterAnimal(ctx, &forestroute.Animal{
		Name: animal,
	})
	if err != nil {
		panic(fmt.Errorf("Error in getting registering animal::: %s", err))
	}
	println("Animal:: ")
	println(fmt.Sprintf("::::: KNOWN ANIMAL -->>> %s ::::::", unknown.Known))
	println("DONE!!!")
}
