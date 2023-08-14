package gapi

import (
	"fmt"

	db "github.com/Dev-El-badry/wallet-system/db/sqlc"
	"github.com/Dev-El-badry/wallet-system/pb"
	"github.com/Dev-El-badry/wallet-system/token"
	"github.com/Dev-El-badry/wallet-system/util"
)

type Server struct {
	pb.UnimplementedSimpleWalletServer
	store      db.Store
	tokenMaker token.Maker
	config     util.Config
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{store: store, config: config, tokenMaker: tokenMaker}

	return server, nil
}
