package service

import (
	"context"
	pb "user_service/proto"
	"user_service/strorage"
)

type UserService struct {
	St *strorage.StorageI
	pb.UnimplementedUserServiceServer
}

func NewUserService(st *strorage.StorageI) *UserService {
	return &UserService{
		St: st,
	}
}

func (ser *UserService) CreateUser(request *pb.UserCreateRequest) (*pb.Void, error) {
	response, err := ser.St.UserService.CreateUser(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (ser *UserService) UpdateUser(request *pb.UserUpdatedRequest) (*pb.Void, error) {
	response, err := ser.St.UserService.UpdatedUser(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (ser *UserService) DeleteUser(request *pb.GetByIdRequest) (*pb.Void, error) {
	response, err := ser.St.UserService.DeleteUser(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (ser *UserService) GetUser(ctx context.Context, request *pb.UserFilterRequest) ([]*pb.UserResponse, error) {
	response, err := ser.St.UserService.GetUser(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (ser *UserService) GetUserCreditCard(request *pb.GetByIdRequest) ([]*pb.UserCredit, error) {
	response, err := ser.St.UserService.GetUserCreditCard(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (ser *UserService) GetUserDepositCard(request *pb.GetByIdRequest) ([]*pb.UserDeposit, error) {
	response, err := ser.St.UserService.GetUserDepositCard(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (ser *UserService) GetUserCard(request *pb.GetByIdRequest) (*pb.UserCard, error) {
	response, err := ser.St.UserService.GetUserCard(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (ser *UserService) GetUserById(request *pb.GetByIdRequest) (*pb.UserAmount, error) {
	response, err := ser.St.UserService.GetUserCardAmount(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
