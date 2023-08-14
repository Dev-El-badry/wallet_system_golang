package gapi

import (
	db "github.com/Dev-El-badry/wallet-system/db/sqlc"
	"github.com/Dev-El-badry/wallet-system/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.Users) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Email:             user.Email,
		FullName:          user.FullName,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
