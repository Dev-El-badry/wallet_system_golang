package gapi

import (
	"context"

	db "github.com/Dev-El-badry/wallet-system/db/sqlc"
	"github.com/Dev-El-badry/wallet-system/pb"
	"github.com/Dev-El-badry/wallet-system/util"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to has password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		Email:          req.GetEmail(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			switch pgErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exists: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}
	return resp, nil
}
