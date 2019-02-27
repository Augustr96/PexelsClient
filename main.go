package main

import (
	"context"
	"fmt"
	"github.com/Augustr96/GoPexels"
	"github.com/Augustr96/unifiedproto/goout/pexels"
	"github.com/Augustr96/unifiedproto/goout/utils"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

const (
	port = ":1081"
	restPort = ":50053"
)

type Server struct {
	Gpc *GoPexels.Client
}

func (s *Server) GetPhoto(ctx context.Context, in *pexels.GetPhotoRequest) ( response *pexels.GetPhotoResponse, err error){
	response = &pexels.GetPhotoResponse{
		Status: &utils.Status{
			Code: utils.StatusCode_STATUS_CODE_OK,
			Message: "ok",
		},
	}
	res, err := s.Gpc.GetPhoto(in.GetId())
	if err != nil {
		response.Status.Message = err.Error()
		return response, nil
	}

	response.Status.Message = res.Url
	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	go runRest()

	server := new(Server)
	server.Gpc = GoPexels.NewClient(os.Getenv("PexelsToken"))

	pexels.RegisterPexelsServer(s, server)

	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
	}

}

func runRest() error {
	ctx := context.Background()

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pexels.RegisterPexelsHandlerFromEndpoint(ctx, gwmux, port, opts)
	if err != nil {
		goto Err
	}
	mux.Handle("/", gwmux)

	err = http.ListenAndServe(restPort, mux)
	if err != nil {
		goto Err
	}

Err:
	if err != nil {
		panic(err)
	}
	return err
}
