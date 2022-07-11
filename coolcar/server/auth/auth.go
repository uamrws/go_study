package auth

import (
	"context"
	"go.uber.org/zap"
	authpb "go_study/coolcar/server/auth/api/gen"
)

type Service struct {
	authpb.UnimplementedAuthServiceServer
	Logger *zap.Logger
}

func (s *Service) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("receive request code", zap.String("code", request.Code))
	return &authpb.LoginResponse{AccessToke: "asd", ExpiresIn: 123877823}, nil
}
