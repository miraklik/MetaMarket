package authgrpc

import (
	"context"
	"errors"
	"internal/internal/lib/logger/sl"
	"internal/internal/services/auth"
	"log/slog"

	stroge "internal/internal/storage"

	ssov1 "github.com/GolangLessons/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
	log  *slog.Logger
}

func Register(gRPCServer *grpc.Server, auth Auth, logger *slog.Logger) {
	ssov1.RegisterAuthServer(gRPCServer, &serverAPI{
		auth: auth,
		log:  logger,
	})
}

func (s *serverAPI) Login(
	ctx context.Context,
	in *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	log := s.log.With(slog.String("op", "serverAPI.Login"))

	if in.Email == "" {
		log.Warn("email is missing in request")
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		log.Warn("password is missing in request")
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if in.GetAppId() == 0 {
		log.Warn("app_id is missing in request")
		return nil, status.Error(codes.InvalidArgument, "app_id is required")
	}

	// Вызов функции аутентификации
	log.Info("calling auth.Login", slog.String("email", in.Email), slog.Int("appID", int(in.GetAppId())))
	token, err := s.auth.Login(ctx, in.GetEmail(), in.GetPassword(), int(in.GetAppId()))
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCredentials) {
			log.Warn("login failed due to invalid credentials", sl.Err(err))
			return nil, status.Error(codes.InvalidArgument, "invalid email or password")
		}

		log.Error("internal error during login", sl.Err(err))
		return nil, status.Error(codes.Internal, "failed to login")
	}

	log.Info("user logged in successfully", slog.String("token", token))
	return &ssov1.LoginResponse{Token: token}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	in *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	if in.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if in.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	uid, err := s.auth.RegisterNewUser(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		if errors.Is(err, stroge.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &ssov1.RegisterResponse{UserId: uid}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	in *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	if in.UserId == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	isAdmin, err := s.auth.IsAdmin(ctx, in.GetUserId())
	if err != nil {
		if errors.Is(err, stroge.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		return nil, status.Error(codes.Internal, "failed to check admin status")
	}

	return &ssov1.IsAdminResponse{IsAdmin: isAdmin}, nil
}
