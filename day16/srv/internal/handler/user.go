package handler

import (
	"context"
	__ "srv/proto"
)

type UserServer struct {
	__.UnimplementedUserServerServer
}

func (s *UserServer) Register(_ context.Context, in *__.RegisterRequest) (*__.RegisterReply, error) {

	return &__.RegisterReply{
		Code: 0,
		Msg:  "",
	}, nil
}
func (s *UserServer) Login(_ context.Context, in *__.LoginRequest) (*__.LoginReply, error) {

	return &__.LoginReply{
		Code: 0,
		Msg:  "",
		Id:   0,
	}, nil
}
