package http

import (
	pbAuthor "github.com/adityarizkyramadhan/synapsis-test/api-gateway/internal/client/author/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthorRoutes struct {
	authorClient pbAuthor.AuthorHandlerClient
}

func NewAuthorRoutes(authorClient pbAuthor.AuthorHandlerClient) *AuthorRoutes {
	return &AuthorRoutes{authorClient: authorClient}
}

func (a *AuthorRoutes) Init() error {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	a.authorClient = pbAuthor.NewAuthorHandlerClient(conn)

	return nil
}
