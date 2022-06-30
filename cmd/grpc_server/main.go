package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	//_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	//_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/nikitads9/note-service-api/internal/app/api/note_v1"
	"github.com/nikitads9/note-service-api/internal/app/repository"
	"github.com/nikitads9/note-service-api/internal/app/service/note"
	pb "github.com/nikitads9/note-service-api/pkg/note_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	//_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

const (
	grpcAdress  = ":50051"
	httpAddress = ":8000"
)

const (
	host     = "localhost"
	dbPort   = "5444"
	user     = "postgres"
	password = "notes_pass"
	dbName   = "notes_db"
	ssl      = "disable"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(startGRPC())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(startHTTP())
	}()

	wg.Wait()
}

func startGRPC() error {
	//nolint
	list, err := net.Listen("tcp", grpcAdress)
	if err != nil {
		return fmt.Errorf("failed to create listener %v", err.Error())
	}
	defer list.Close()

	DbDsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, dbPort, user, password, dbName, ssl)

	db, err := sqlx.Open("pgx", DbDsn)
	if err != nil {
		return fmt.Errorf("failed to establish connection with database")
	}
	defer db.Close()

	noteRepository := repository.NewNoteRepository(db)
	noteService := note.NewNoteService(noteRepository)

	s := grpc.NewServer()
	pb.RegisterNoteV1Server(s, note_v1.NewNoteV1(noteService))
	if err = s.Serve(list); err != nil {
		return fmt.Errorf("failed to process gRPC server %v", err.Error())
	}

	return nil
}

func startHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterNoteV1HandlerFromEndpoint(ctx, mux, grpcAdress, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpAddress, mux)
}
