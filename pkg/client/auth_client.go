package client

import (
	"context"
	"flotta-home/mindbond/chat-service/pkg/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func InitAuthServiceClient(url string) AuthServiceClient {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return AuthServiceClient{
		Client: pb.NewAuthServiceClient(cc),
	}
}

func (a *AuthServiceClient) Validate(token string) (*pb.ValidateResponse, error) {
	request := &pb.ValidateRequest{
		Token: token,
	}
	return a.Client.Validate(context.Background(), request)
}

func (a *AuthServiceClient) LookupById(userId int64) (*pb.LookupResponse, error) {
	request := &pb.LookupRequest{
		UserId: userId,
	}
	return a.Client.Lookup(context.Background(), request)
}

func (a *AuthServiceClient) LookupByEmail(email string) (*pb.LookupResponse, error) {
	request := &pb.LookupRequest{
		Email: email,
	}
	return a.Client.Lookup(context.Background(), request)
}

func (a *AuthServiceClient) LookupByHandle(handle string) (*pb.LookupResponse, error) {
	request := &pb.LookupRequest{
		Handle: handle,
	}
	return a.Client.Lookup(context.Background(), request)
}
