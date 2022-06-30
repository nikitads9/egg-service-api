package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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

	client := pb.NewNoteV1Client(con)

	var res *pb.AddNoteResponse
	res, err = client.AddNote(ctx, &pb.AddNoteRequest{
		Title:   "title1",
		Content: "fhdshjdsgd",
	})
	if err != nil {
		log.Printf("failed to add note: %v\n", err.Error())
	}

	fmt.Println("note with id", res.GetResult().GetId(), "added")

	var addedID *pb.MultiAddResponse

	addedID, err = client.MultiAdd(ctx, &pb.MultiAddRequest{
		Notes: []*pb.MultiAddRequest_Notes{
			{
				Title:   "title11",
				Content: "ffdsjfdjf",
			},
			{
				Title:   "title2",
				Content: "sometext",
			},
			{
				Title:   "title3",
				Content: "more text",
			},
		},
	})
	if err != nil {
		log.Printf("failed to add notes: %v\n", err.Error())
	}
	fmt.Printf("added %v IDs", addedID.GetResult().Count)

	_, err = client.RemoveNote(ctx, &pb.RemoveNoteRequest{Id: int64(2)})
	if err != nil {
		log.Printf("failed to remove note: %v\n", err.Error())
	}

	notes, err := client.GetList(ctx, &emptypb.Empty{})
	if err != nil {
		log.Printf("failed to get all notes: %v\n", err.Error())
	}
	fmt.Printf("%v\n", notes.GetResults())

	_, err = client.UpdateNote(ctx, &pb.UpdateNoteRequest{
		Id:      3,
		Title:   "newtitle",
		Content: "newcontent",
	})
	if err != nil {
		log.Printf("failed to update a note: %v\n", err.Error())
	}

	note, err := client.GetNote(ctx, &pb.GetNoteRequest{
		Id: 3,
	})
	if err != nil {
		log.Printf("failed to get note: %v\n", err.Error())
	}
	fmt.Printf("%v\n", note)
}
