package grpc

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"sso/internal/api/grpc/controller"
	"sso/internal/errs"
	"sso/pkg/contracts/ssocontract"
)

type Config interface {
	Port() int
}

type Server struct {
	config     Config
	controller *controller.Controller
	server     *grpc.Server
	log        *slog.Logger
}

func New(config Config, controller *controller.Controller, log *slog.Logger) *Server {
	return &Server{
		config:     config,
		controller: controller,
		log:        log,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port()))
	if err != nil {
		return errs.Wrap("failed to listen", err)
	}

	s.server = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
		),
	)

	ssocontract.RegisterSSOServer(s.server, s.controller)

	go func() {
		err = s.server.Serve(listener)
		if err != nil {
			panic(errs.Wrap("failed to serve", err))
		}
	}()

	return nil
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
