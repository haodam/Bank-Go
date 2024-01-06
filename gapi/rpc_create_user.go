package gapi

import (
	"context"
	db "github.com/haodam/Bank-Go/db/sqlc"
	"github.com/haodam/Bank-Go/pb"
	"github.com/haodam/Bank-Go/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}
	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exitss: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to hash user: %s", err)
	}

	// converter data db.User on pb.User
	rsp := &pb.CreateUserResponse{
		User: converter(user),
	}

	return rsp, nil
}
