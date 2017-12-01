package server

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/gosample/workshop/proto/list"
	"github.com/gosample/workshop/proto/notifier"
	"github.com/gosample/workshop/proto/users"
	"github.com/gosample/workshop/shared"
)

type Server struct{}

func (s *Server) CreateUser(ctx context.Context, in *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	userID, err := shared.CreateUser(in.Email)

	response := new(users.CreateUserResponse)
	if err == nil {
		log.Printf("[user.Create] New user ID: %s", userID)

		createInitialItem(userID)
		go notify(in.Email)

		response.Message = "User created successfully"
		response.Id = userID
		response.Code = http.StatusCreated
	} else {
		response.Message = err.Error()
		response.Code = http.StatusInternalServerError
	}

	return response, err
}

// Call notifier service
func notify(email string) {
	_, err := shared.NotifierClient.Email(context.Background(), &notifier.EmailRequest{
		Email: email,
	})
	log.Printf("[user.Create] Notifying user: %v", err)
}

// Create initial item in todo list
func createInitialItem(userID string) {
	_, err := shared.ListClient.CreateItem(context.Background(), &list.CreateItemRequest{
		Message: "Welcome to Workshop!",
		UserId:  userID,
	})
	if err != nil {
		log.Printf("[user.Create] Cannot create item: %v", err)
	}
}
