package auth

import (
	"context"
	"log"

	userpb "github.com/acme-corp/user-service/proto/gen/go/user/v1"
	"google.golang.org/grpc"
)

// RegisterUser creates a new user account with the provided details.
func RegisterUser(ctx context.Context, conn *grpc.ClientConn, name, email, phone string) (*userpb.CreateUserResponse, error) {
	client := userpb.NewUserServiceClient(conn)

	req := &userpb.CreateUserRequest{
		Name:        name,
		Email:       email,
		PhoneNumber: phone,
	}

	log.Printf("registering user: name=%s email=%s phone=%s", name, email, phone)

	resp, err := client.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	log.Printf("user registered: id=%s", resp.GetUserId())
	return resp, nil
}
