package main

import (
	"log"
	"net"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/config/database"
	pb "github.com/adityarizkyramadhan/synapsis-test/book-service/internal/handler/grpc"
	grpcImplementation "github.com/adityarizkyramadhan/synapsis-test/book-service/internal/handler/grpc/implementation"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/repository"
	"github.com/adityarizkyramadhan/synapsis-test/book-service/internal/service"
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
	if err := database.AutoMigrate(&model.Book{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repoBook := repository.Newbook(db)
	serviceBook := service.NewBook(repoBook)
	grpcHandler := grpcImplementation.NewBook(serviceBook)
	grpcServer := grpc.NewServer()
	pb.RegisterBookHandlerServer(grpcServer, grpcHandler)

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
