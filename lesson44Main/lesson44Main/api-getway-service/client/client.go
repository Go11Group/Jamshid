package client

import (
	pb "api_get_way/proto"
	"context"
)

type UserClient struct {
	Client pb.UserServiceClient
}

func NewUserService(conn pb.UserServiceClient) *UserClient {
	return &UserClient{
		Client: conn,
	}
}

func (cl *UserClient) CreateUser(request *pb.UserCreateRequest) (*pb.Void, error) {
	response, err := cl.Client.CreateUser(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (cl *UserClient) UpdateUser(request *pb.UserUpdatedRequest) (*pb.Void, error) {
	response, err := cl.Client.UpdateUser(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (cl *UserClient) DeleteUser(request *pb.GetByIdRequest) (*pb.Void, error) {
	response, err := cl.Client.DeleteUser(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (cl *UserClient) GetUser(request *pb.UserFilterRequest) ([]*pb.UserResponse, error) {
	response, err := cl.Client.GetUser(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (cl *UserClient) GetUserCreditCard(request *pb.GetByIdRequest) (*pb.UserCredit, error) {
	response, err := cl.Client.GetUserCreditCard(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (cl *UserClient) GetUserDepositCard(request *pb.GetByIdRequest) (*pb.UserDeposit, error) {
	response, err := cl.Client.GetUserDepositCard(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (cl *UserClient) GetUserCard(request *pb.GetByIdRequest) (*pb.UserCard, error) {
	response, err := cl.Client.GetUserCard(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

//func (cl *UserClient) GetUserById(request *pb.GetByIdRequest) (*pb.UserAmount, error) {
//	response, err := cl.Client.
//	if err != nil {
//		return nil, err
//	}
//	return response, nil
//}
