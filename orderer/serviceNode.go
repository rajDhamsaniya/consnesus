package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "study/GitHub/consensus/protoc/orderer"

	"./kafka"

	"github.com/gogo/protobuf/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type server struct{}

const (
	ordererPort string = ":50054"
)

//var producer sarama.AsyncProducer

func (s *server) SubmitTx(ctx context.Context, in *pb.EndorsedTx) (*empty.Empty, error) {

	msg, err := proto.Marshal(in)
	if err != nil {
		fmt.Println("can not convert")
	}
	//fmt.Println("Aye")
	//var producer sarama.AsyncProducer
	producer := kafka.GetProducer()
	kafka.SendTx(producer, msg)
	// fmt.Println(producer)
	return &empty.Empty{}, nil
}
func (s *server) GetBlocks(in *pb.BlockId, stream pb.Orderer_GetBlocksServer) error {

	return nil
}
func (s *server) GetSpecificBlock(ctx context.Context, in *pb.BlockId) (*pb.Block, error) {

	return &pb.Block{}, nil
}

func main() {

	producer := kafka.StartProducer()
	fmt.Println("Initialization")
	go kafka.StartConsumer(producer)
	go kafka.ConnectGossipService()
	lis, err := net.Listen("tcp", ordererPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrdererServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
