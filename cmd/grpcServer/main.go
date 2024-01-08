package main

import (
	"database/sql"
	"net"

	"github.com/RomeroGabriel/go-gRPC/internal/db"
	"github.com/RomeroGabriel/go-gRPC/internal/pb"
	"github.com/RomeroGabriel/go-gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	dbc, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer dbc.Close()

	categoryDb := db.NewCategoryDb(dbc)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServicer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServicer, categoryService)
	reflection.Register(grpcServicer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServicer.Serve(lis); err != nil {
		panic(err)
	}
}
