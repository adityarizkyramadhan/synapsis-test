package http

import "github.com/adityarizkyramadhan/synapsis-test/user-service/internal/service"

type User struct {
	serviceUser *service.User
}

func NewUser(serviceUser *service.User) *User {
	return &User{serviceUser: serviceUser}
}
