package main

import (
	"context"
	"github.com/Ankr-network/go-ethereum-jsonrpc/jsonrpc"
	"log"
	"net/http"
)

type TestServer struct {
}

type Request struct {
	Param1 int    `json:"param_1"`
	Param2 string `json:"param_2"`
}

type Response struct {
	Param1     int `json:"param_1"`
	Param2Size int `json:"param_2_size"`
}

func (s *TestServer) TestRequest(ctx context.Context, request Request) (*Response, error) {

	return &Response{
		Param1:     request.Param1,
		Param2Size: len(request.Param2),
	}, nil

}

func main()  {
	jsonRpc := jsonrpc.NewServer()
	if err := jsonRpc.RegisterName("namespace", &TestServer{}); err != nil{
		log.Fatalf("failed to register namespace")
	}
	log.Printf("listening localhost:8801")
	if err := http.ListenAndServe("localhost:8801", jsonRpc); err != nil {
		log.Fatalf("failed to register namespace")
	}
}