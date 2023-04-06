package main

import (
	"database/sql"

	"github.com/eduardo-villasboas/gRPC-hands-on/internal/database"

	"github.com/eduardo-villasboas/gRPC-hands-on/internal/service"

	"google.golang.org/grpc"

	"net"

	"google.golang.org/grpc/reflection"

	"github.com/eduardo-villasboas/gRPC-hands-on/internal/pb"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)

	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}

}	