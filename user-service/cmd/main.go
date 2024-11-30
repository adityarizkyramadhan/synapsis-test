package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/adityarizkyramadhan/synapsis-test/user-service/config/cache"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/config/database"
	pb "github.com/adityarizkyramadhan/synapsis-test/user-service/internal/handler/grpc"
	grpcImplementation "github.com/adityarizkyramadhan/synapsis-test/user-service/internal/handler/grpc/implementation"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/model"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/repository"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/service"
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
	if err := database.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	redis := cache.NewRedis()
	_, err = redis.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	repoUser := repository.NewUser(db, redis)
	serviceUser := service.NewUser(repoUser)
	grpcHandler := grpcImplementation.NewUser(serviceUser)
	grpcServer := grpc.NewServer()
	pb.RegisterUserHandlerServer(grpcServer, grpcHandler)
	listener, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server is running on port " + os.Getenv("GRPC_PORT"))
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
