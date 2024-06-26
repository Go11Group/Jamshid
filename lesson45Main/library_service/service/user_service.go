package service

import (
	"context"
	pb "library_service/genproto"
	strorage "library_service/storage"
)

type UserService struct {
	Serv strorage.StorageI
	pb.UnimplementedUserServiceServer
}

func NewUserService(st strorage.StorageI) *UserService {
	return &UserService{Serv: st}
}
func (serv *UserService) CreateUser(ctc context.Context, in *pb.CreateUserRequest) (*pb.Void, error) {
	return serv.Serv.UserI.CreateUser(in)
}
func (serv *UserService) UpdateUser(ctc context.Context, in *pb.UpdatedUserRequest) (*pb.Void, error) {
	return serv.Serv.UserI.UpdateUser(in)
}
func (serv *UserService) DeleteUser(ctc context.Context, in *pb.ByIdRequest) (*pb.Void, error) {
	return serv.Serv.UserI.DeleteUser(in)
}
func (serv *UserService) GetUser(ctc context.Context, in *pb.UserFilterRequest) (*pb.UsersResponse, error) {
	return serv.Serv.UserI.GetUser(in)
}
