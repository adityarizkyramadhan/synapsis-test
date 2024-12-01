package main

import (
	"log"
	"net"
	"os"

	"github.com/adityarizkyramadhan/synapsis-test/book-service/config/database"
	pbBook "github.com/adityarizkyramadhan/synapsis-test/book-service/internal/handler/grpc"
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
	if err := database.AutoMigrate(&model.Book{}, &model.CategoryBook{}, &model.Borrowing{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repoBook := repository.Newbook(db)
	serviceBook := service.NewBook(repoBook)
	grpcHandler := grpcImplementation.NewBook(serviceBook)
	grpcServer := grpc.NewServer()
	pbBook.RegisterBookHandlerServer(grpcServer, grpcHandler)

	repoCategoryBook := repository.NewCategoryBook(db)
	serviceCategoryBook := service.NewCategoryBook(repoCategoryBook)
	grpcHandlerCategoryBook := grpcImplementation.NewCategoryBook(serviceCategoryBook)
	pbBook.RegisterCategoryBookHandlerServer(grpcServer, grpcHandlerCategoryBook)

	repoBorrowing := repository.NewBorrowing(db)
	serviceBorrowing := service.NewBorrowing(repoBorrowing)
	grpcHandlerBorrowing := grpcImplementation.NewBorrowing(serviceBorrowing)
	pbBook.RegisterBorrowingHandlerServer(grpcServer, grpcHandlerBorrowing)

	listener, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("gRPC server is running on port " + os.Getenv("GRPC_PORT"))
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
