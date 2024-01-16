package gapi

import (
	"fmt"
	db "github.com/haodam/Bank-Go/db/sqlc"
	"github.com/haodam/Bank-Go/pb"
	"github.com/haodam/Bank-Go/token"
	"github.com/haodam/Bank-Go/util"
	"github.com/haodam/Bank-Go/worker"
)

// Server servers gRPC request for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil

}
