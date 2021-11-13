package user_service

import (
	tool "music_api_server/Tool"
	userModel "music_api_server/repositories/user"
	req "music_api_server/request"
)

//UserService Implement the user_service Protocol
type UserService struct {
}

func (user *UserService)Register(req *req.RegisterRequest) error {
	bcrypt := tool.Bcrypt{
		Cost: 10,
	}

	hashPassword , err := bcrypt.MakePassword([]byte(req.Password))
	if err != nil{
		return err
	}

	req.Password = string(hashPassword)
	return userModel.CreateUser(req) // the model will return the error
}

func (user *UserService)Login(req *req.LoginRequest) error {
	return nil
}