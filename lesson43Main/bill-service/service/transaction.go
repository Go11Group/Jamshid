package service

import (
	pb "bill_service/genproto"
	"context"
)

func (service *CardService) CreateTerminal(ctx context.Context, in *pb.CreateTerminalRequest)(*pb.Void,error) {
	return service.St.Terminal.CreateTerminal(in)
}
func (service *CardService) UpdateTerminal(ctx context.Context, in *pb.UpdateTerminalRequest)(*pb.Void,error) {
	return service.St.Terminal.UpdateTerminal(in);
}
func (service *CardService) DeleteTerminal(ctx context.Context, in *pb.ByIdRequest)(*pb.Void,error) {
	return service.St.Terminal.DeleteTerminal(in);
}
func (service *CardService) GetTerminal(ctx context.Context, in *pb.TerminalFilterRequest)(*pb.TerminalsResponse,error) {
	return service.St.Terminal.GetTerminal(in);
}
func (service *CardService) GetTerminalById(ctx context.Context, in *pb.ByIdRequest)(*pb.TerminalResponse,error) {
	return service.St.Terminal.GetTerminalById(in);
}


