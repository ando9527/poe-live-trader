package server

import (
	"context"
	"log"
	"net"

	"github.com/ando9527/poe-live-trader/pkg/grpc/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Database interface{
	GetFilterList()*pb.FilterList
	RemoveFilter(filter string)error
	AddFilter(filter string, comment string) error
}

type Server struct{
	address string
	database Database
}

func NewServer(address string, database Database) *Server {
	return &Server{address: address, database: database}
}



func (s *Server)Run()  {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		logrus.Fatalf("failed to listen port：%v", err)
	}

	g := grpc.NewServer()
	pb.RegisterYourServiceServer(g, s)

	reflection.Register(g)
	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve：%v", err)
	}
}

func (s *Server) GetFilterList(ctx context.Context, in *pb.Request) (*pb.FilterList, error) {
	out:=&pb.FilterList{
		Filters:              []*pb.Filter{&pb.Filter{
			Value:                "123",
			Comment:              "yolo",
		}},
	}
	return out, nil
}