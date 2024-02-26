package main

import (
	"database/sql"
	"net"

	"github.com/gstanleysilva/go-grpc/internal/database"
	"github.com/gstanleysilva/go-grpc/internal/pb"
	"github.com/gstanleysilva/go-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	//Create a new GRPC Server
	grpcServer := grpc.NewServer()

	//Link GRPC Server with Category Service
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	//Register reflection service on GRPC Server - EVANS
	reflection.Register(grpcServer)

	//Open Communication TCP Port
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}

	//Start GRPC Server
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
