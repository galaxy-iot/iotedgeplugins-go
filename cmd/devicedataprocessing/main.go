package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	grpcplugin "github.com/galaxy-iot/iotedgeplugins-go/pkg/plugins/grpc"
	"github.com/wh8199/log"
	"google.golang.org/grpc"
)

type Processor struct {
	grpcplugin.UnimplementedProcessorServer
}

func (p *Processor) ProcessDataOnce(ctx context.Context, ds *grpcplugin.DataSet) (*grpcplugin.DataSet, error) {
	log.Info(ds)
	return ds, nil
}

// processor functions
func (p *Processor) ProcessDataStream(s grpcplugin.Processor_ProcessDataStreamServer) error {
	for {
		ds, err := s.Recv()
		if err != nil {
			log.Error(err)
			return err
		}

		log.Info(ds)

		if err := s.Send(ds); err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}

var (
	port = 9100
)

func init() {
	flag.IntVar(&port, "port", 9100, "grpc server port")
}

func main() {
	s := grpc.NewServer()
	grpcplugin.RegisterProcessorServer(s, &Processor{})

	// listen tcp connection
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// start grpc server
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
