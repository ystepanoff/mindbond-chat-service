package main

import (
	"flotta-home/mindbond/chat-service/pkg/client"
	"fmt"
	"log"
	"net"

	"flotta-home/mindbond/chat-service/pkg/config"
	"flotta-home/mindbond/chat-service/pkg/db"
	pb "flotta-home/mindbond/chat-service/pkg/pb"
	services "flotta-home/mindbond/chat-service/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)
	authClient := client.InitAuthServiceClient(c.AuthServiceUrl)
	translatorClient := client.InitTranslatorServiceClient(c.TranslatorServiceUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Chat service on", c.Port)

	s := services.Server{
		H:                h,
		AuthClient:       authClient,
		TranslatorClient: translatorClient,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
