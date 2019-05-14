package main

import (
	"context"
	pb "github.com/andrianay/userservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

type repository interface {
	Create(user *pb.User) (*pb.User, error)
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	users []*pb.User
}

// Create a new consignment
func (repo *Repository) Create(user *pb.User) (*pb.User, error) {
	updated := append(repo.users, user)
	repo.users = updated
	return user, nil
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

func (s *service) Create(ctx context.Context, user *pb.User) (*pb.Response, error) {
	user, err := s.repo.Create(user)

	if err != nil {
		return nil, err
	}
	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &pb.Response{User: user}, nil
}

func (s *service) Get(context.Context, *pb.User) (*pb.Response, error) {
	panic("implement me")
}

func (s *service) GetAll(context.Context, *pb.Request) (*pb.Response, error) {
	panic("implement me")
}

func (s *service) Auth(context.Context, *pb.User) (*pb.Token, error) {
	panic("implement me")
}

func (s *service) ValidateToken(context.Context, *pb.Token) (*pb.Token, error) {
	panic("implement me")
}

func main() {
	repo := &Repository{}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterUserServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
