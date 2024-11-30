package main

import (
	"log"
	"net"
	"os"

	"github.com/adityarizkyramadhan/synapsis-test/category-service/config/database"
	pb "github.com/adityarizkyramadhan/synapsis-test/category-service/internal/handler/grpc"
	grpcImplementation "github.com/adityarizkyramadhan/synapsis-test/category-service/internal/handler/grpc/implementation"
	"github.com/adityarizkyramadhan/synapsis-test/category-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/category-service/internal/repository"
	"github.com/adityarizkyramadhan/synapsis-test/category-service/internal/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := database.AutoMigrate(&model.Category{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repoCategory := repository.NewCategory(db)
	serviceCategory := service.NewCategory(repoCategory)
	grpcHandler := grpcImplementation.NewCategory(serviceCategory)
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryHandlerServer(grpcServer, grpcHandler)

	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port " + os.Getenv("GRPC_PORT"))
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
