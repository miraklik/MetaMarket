package suite

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"internal/internal/config"
	"net"
	"os"
	"strconv"
	"testing"

	ssov1 "github.com/GolangLessons/protos/gen/go/sso"
	"google.golang.org/grpc"
)

const (
	grpcHost = "127.0.0.1"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadPath(configPath())

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	// Запустить при deploy
	/*creds, err := credentials.NewClientTLSFromFile("D:/Program Files/MetaMarket/backend/server-cert.pem", "")
	if err != nil {
		t.Fatalf("failed to load TLS credentials: %v", err)
	}*/

	conn, err := grpc.Dial(
		grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: ssov1.NewAuthClient(conn),
	}
}

func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}

func configPath() string {
	key := "./config/local.yaml"

	if v := os.Getenv(key); v != "" {
		return v
	}

	return "../config/local_tests.yaml"
}
