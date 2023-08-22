package services

import (
	"context"
	"flotta-home/mindbond/chat-service/pkg/db"
	"flotta-home/mindbond/chat-service/pkg/models"
	pb "flotta-home/mindbond/chat-service/pkg/pb"
	"net/http"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	var chat models.Chat
	chat.User1ID = req.User1ID
	chat.User2ID = req.User2ID
	if result := s.H.DB.Create(&chat); result.Error != nil {
		return &pb.CreateChatResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.CreateChatResponse{
		Status: http.StatusCreated,
		ChatId: chat.Id,
	}, nil
}

func (s *Server) FindChat(ctx context.Context, req *pb.FindChatRequest) (*pb.FindChatResponse, error) {
	var chat models.Chat
	if result := s.H.DB.First(&chat, req.User1Id, req.User2Id); result.Error != nil {
		return &pb.FindChatResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	if chat.DeletedAt.Valid == true {
		return &pb.FindChatResponse{
			Status: http.StatusNotFound,
			Error:  "This chat was previously deleted",
		}, nil
	}
	data := &pb.FindChatData{
		ChatId: chat.Id,
	}
	return &pb.FindChatResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) DeleteChat(ctx context.Context, req *pb.DeleteChatRequest) (*pb.DeleteChatResponse, error) {
	var chat models.Chat
	if result := s.H.DB.First(&chat, req.ChatId); result.Error != nil {
		return &pb.DeleteChatResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	if chat.DeletedAt.Valid == true {
		return &pb.DeleteChatResponse{
			Status: http.StatusNotFound,
			Error:  "This chat was previously deleted",
		}, nil
	}
	if result := s.H.DB.Delete(&chat); result.Error != nil {
		return &pb.DeleteChatResponse{
			Status: http.StatusBadRequest,
			Error:  "Unknown error",
		}, nil
	}
	return &pb.DeleteChatResponse{
		Status: http.StatusOK,
	}, nil
}
