package client

import (
	"context"
	"flotta-home/mindbond/chat-service/pkg/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TranslatorServiceClient struct {
	Client pb.TranslatorServiceClient
}

func InitTranslatorServiceClient(url string) TranslatorServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return TranslatorServiceClient{
		Client: pb.NewTranslatorServiceClient(cc),
	}
}

func (c *TranslatorServiceClient) Translate(message string, fromLang string, toLang string) (*pb.TranslateResponse, error) {
	request := &pb.TranslateRequest{
		Message:  message,
		FromLang: fromLang,
		ToLang:   toLang,
	}
	return c.Client.Translate(context.Background(), request)
}
