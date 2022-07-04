package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nikitads9/egg-service-api/pkg/egg_api"
	"google.golang.org/grpc"
)

const grpcAdress = "localhost:50051"

func main() {
	ctx := context.Background()
	//nolint
	con, err := grpc.Dial(grpcAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err.Error())
	}
	defer con.Close()

	client := pb.NewEggNutritionClient(con)

	var res *pb.AddMealResponse
	res, err = client.AddMeal(ctx, &pb.AddMealRequest{
		UserId: int64(4221),
		Weight: float32(390),
	})
	if err != nil {
		log.Printf("failed to add meal: %v\n", err.Error())
	}

	fmt.Println("meal with id", res.GetResult().GetId(), "added")

	res, err = client.AddMeal(ctx, &pb.AddMealRequest{
		UserId: int64(4221),
		Weight: float32(325),
	})
	if err != nil {
		log.Printf("failed to add meal: %v\n", err.Error())
	}

	fmt.Println("meal with id", res.GetResult().GetId(), "added")

	meals, err := client.GetList(ctx, &pb.GetListRequest{UserId: int64(4221)})
	if err != nil {
		log.Printf("failed to get all meals: %v\n", err.Error())
	}
	fmt.Printf("%v\n", meals.GetResults())

	/*_, err = client.RemoveMeal(ctx, &pb.RemoveMealRequest{
		Id:     2,
		UserId: int64(4221),
	})
	if err != nil {
		log.Printf("failed to remove meal: %v\n", err.Error())
	}

	_, err = client.UpdateMeal(ctx, &pb.UpdateMealRequest{
		Id:       1,
		UserId:   4221,
		MealDate: "Mon Jan _2 15:04:05 2006",
		Weight:   float32(330),
	})
	if err != nil {
		log.Printf("failed to update a meal: %v\n", err.Error())
	}
	*/
	meal, err := client.GetMeal(ctx, &pb.GetMealRequest{
		Id:     1,
		UserId: 4221,
	})
	if err != nil {
		log.Printf("failed to get meal: %v\n", err.Error())
	}
	fmt.Printf("%v\n", meal)
}
