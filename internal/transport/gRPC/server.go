package gRPC

import (
	"github.com/sirupsen/logrus"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/service"
	grpcHandler "gitlab.digital-spirit.ru/study/artem_crud/internal/transport/gRPC/handler"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/transport/gRPC/proto"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Server *grpc.Server
}

func (s *Server) Run(network string, port string, services *service.Service) error {
	lis, err := net.Listen(network, ":"+port)
	if err != nil {
		return err
	}

	s.Server = grpc.NewServer()
	proto.RegisterRecordsServer(s.Server, grpcHandler.NewRecordHandler(services))

	logrus.Printf("server listening at %v", lis.Addr())
	return s.Server.Serve(lis)
}
