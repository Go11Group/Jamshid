package strorage

import pb "user_service/proto"

type StorageI struct {
	UserService UserServiceI
}
type UserServiceI interface {
	CreateUser(basket *pb.UserCreateRequest) (*pb.Void, error)
	UpdatedUser(basket *pb.UserUpdatedRequest) (*pb.Void, error)
	DeleteUser(basket *pb.GetByIdRequest) (*pb.Void, error)
	GetUser(basket *pb.UserFilterRequest) ([]*pb.UserResponse, error)
	GetUserDepositCard(basket *pb.GetByIdRequest) ([]*pb.UserDeposit, error)
	GetUserCreditCard(basket *pb.GetByIdRequest) ([]*pb.UserCredit, error)
	GetUserCardAmount(basket *pb.GetByIdRequest) (*pb.UserAmount, error)
	GetUserCard(basket *pb.GetByIdRequest) (*pb.UserCard, error)
}
