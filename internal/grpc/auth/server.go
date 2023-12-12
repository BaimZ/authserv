package auth

import (
	"context"

	ssov1 "github.com/baimz/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	context.Context,
	*ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	panic("implement me")
}
func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
func (s *serverAPI) Login(
	context.Context,
	**ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	panic("implement me")
}
func (s *serverAPI) IsAdmin(
	context.Context,
	*ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
