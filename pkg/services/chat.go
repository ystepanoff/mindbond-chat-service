package services

import (
	"context"
	"flotta-home/mindbond/chat-service/pkg/client"
	"flotta-home/mindbond/chat-service/pkg/db"
	"flotta-home/mindbond/chat-service/pkg/models"
	pb "flotta-home/mindbond/chat-service/pkg/pb"
	"fmt"
	"net/http"
)

type Server struct {
	H                db.Handler
	AuthClient       client.AuthServiceClient
	TranslatorClient client.TranslatorServiceClient
}

func (s *Server) AddContact(ctx context.Context, req *pb.AddContactRequest) (*pb.AddContactResponse, error) {
	var contact models.Contact
	if result, err := s.AuthClient.Validate(req.Token); err != nil || result.Status != http.StatusOK {
		return &pb.AddContactResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("Validation error: %s", result.Status),
		}, nil
	}
	contactUser, err := s.AuthClient.LookupByHandle(req.Handle)
	if err != nil {
		return &pb.AddContactResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("Lookup error: %s", err),
		}, nil
	}
	if contactUser.Status != http.StatusOK {
		return &pb.AddContactResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("Lookup error: %s", contactUser.Error),
		}, nil
	}
	contact.UserId = req.UserId
	contact.ContactId = contactUser.UserId
	if result := s.H.DB.Create(&contact); result.Error != nil {
		return &pb.AddContactResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.AddContactResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) RemoveContact(ctx context.Context, req *pb.RemoveContactRequest) (*pb.RemoveContactResponse, error) {
	var contact models.Contact
	if result, err := s.AuthClient.Validate(req.Token); err != nil || result.Status != http.StatusOK {
		return &pb.RemoveContactResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("Validation error: %s", result.Status),
		}, nil
	}
	contactUser, err := s.AuthClient.LookupByHandle(req.Handle)
	if err != nil || contactUser.Status != http.StatusOK {
		return &pb.RemoveContactResponse{
			Status: http.StatusInternalServerError,
			Error:  fmt.Sprintf("Lookup error: %s", contactUser.Status),
		}, nil
	}
	if result := s.H.DB.First(&contact, req.UserId, contactUser.UserId); result.Error != nil {
		return &pb.RemoveContactResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}
	if result := s.H.DB.Delete(&contact); result.Error != nil {
		return &pb.RemoveContactResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.RemoveContactResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) FetchContacts(req *pb.FetchContactsRequest, stream pb.ChatService_FetchContactsServer) error {
	var contacts []models.Contact
	if result, err := s.AuthClient.Validate(req.Token); err != nil || result.Status != http.StatusOK {
		return fmt.Errorf("Validation error: %s", result.Status)
	}
	if result := s.H.DB.Find(&contacts, req.UserId); result.Error != nil {
		return fmt.Errorf("DB error: %s", result.Error.Error())
	}
	for _, contact := range contacts {
		contactUser, err := s.AuthClient.LookupById(contact.ContactId)
		if err != nil || contactUser.Status != http.StatusOK {
			return fmt.Errorf("Lookup error: %s", contactUser.Status)
		}
		if err := stream.Send(&pb.UserContact{
			UserId: contactUser.UserId,
			Handle: contactUser.Handle,
		}); err != nil {
			return fmt.Errorf("Stream error: %s", err)
		}
	}
	return nil
}

func (s *Server) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	var chat models.Chat
	chat.User1ID = req.User1Id
	chat.User2ID = req.User2Id
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

func (s *Server) AddMessage(ctx context.Context, req *pb.AddMessageRequest) (*pb.AddMessageResponse, error) {
	var chat models.Chat
	var message models.Message
	if result := s.H.DB.First(&chat, req.UserFromId, req.UserToId); result.Error == nil {
		message.ChatId = chat.Id
	} else if result := s.H.DB.First(&chat, req.UserToId, req.UserFromId); result.Error == nil {
		message.ChatId = chat.Id
	} else {
		return &pb.AddMessageResponse{
			Status: http.StatusNotFound,
			Error:  "Chat entity does not exist",
		}, nil
	}
	message.Original = req.Message
	result, err := s.TranslatorClient.Translate(req.Message, req.UserFromLanguage, req.UserToLanguage)
	if err != nil {
		return &pb.AddMessageResponse{
			Status: http.StatusBadRequest,
			Error:  fmt.Sprintf("Translation error: %s", err),
		}, nil
	}
	message.Translation = result.Translation
	if result := s.H.DB.Create(&message); result.Error != nil {
		return &pb.AddMessageResponse{
			Status: http.StatusBadRequest,
			Error:  result.Error.Error(),
		}, nil
	}
	return &pb.AddMessageResponse{
		Status: http.StatusOK,
	}, nil
}
