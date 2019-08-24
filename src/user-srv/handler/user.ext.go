package handler

import (
	"context"
	"dianying/src/share/errors"
	"dianying/src/share/pb"
	"dianying/src/share/utils/log"
	"dianying/src/user-srv/db"
	"go.uber.org/zap"
)

type UserServiceExtHandler struct {
	logger *zap.Logger
}



func NewUserServiceExtHandler() *UserServiceExtHandler {
	return &UserServiceExtHandler{
		logger:log.Instance(),
	}
}

func (u *UserServiceExtHandler)RegistAccount(ctx context.Context, req *pb.RegistAccountReq,res *pb.RegistAccountRes) error{
	userName := req.GetUserName()
	password := req.GetPassword()
	email := req.GetEmail()
	user,err := db.FindUserByEmail(email)
	if err != nil{
		u.logger.Error("error",zap.Error(err))
		return errors.ErrorUserFailed
	}
	if user != nil{
		return errors.ErrorUserAlready
	}
	err = db.InsertUser(userName,password,email)
	if err != nil{
		u.logger.Error("error",zap.Error(err))
		return errors.ErrorUserFailed
	}
	return nil
}

func (u *UserServiceExtHandler) LoginAccount(ctx context.Context,req *pb.LoginAccountReq,res *pb.LoginAccountRes) error {
	email := req.GetEmail()
	password := req.GetPassword()
	user,err := db.FindUserByEmailPassword(email,password)
	if err != nil{
		u.logger.Error("error",zap.Error(err))
		return errors.ErrorUserFailed
	}
	if user == nil{
		return errors.ErrorUserLoginFailed
	}

	res.Email = user.Email
	res.Phone = user.Phone
	res.UserId = user.UserId
	res.UserName = user.UserName
	return nil
}

func (u *UserServiceExtHandler) WantScore(ctx context.Context,req *pb.WantScoreReq,res *pb.WantScoreRes) error {
	panic("implement me")
}

func (u *UserServiceExtHandler) UpdateUserProfile(ctx context.Context,req *pb.UpdateUserProfileReq,res *pb.UpdateUserProfileRes) error {
	userEmail := req.GetUserEmail()
	userName := req.GetUserName()
	userPhone := req.GetUserPhone()
	userId := req.GetUserId()
	if userEmail != ""{
		err := db.UpdateUserEmailProfile(userId,userEmail)
		if err != nil{
			u.logger.Error("error",zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	if userName != ""{
		err := db.UpdateUserNameProfile(userId,userName)
		if err != nil {
			u.logger.Error("error",zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	if userPhone != ""{
		err := db.UpdateUserPhoneProfile(userId,userPhone)
		if err != nil {
			u.logger.Error("error",zap.Error(err))
			return errors.ErrorUserFailed
		}
	}
	return nil
}
