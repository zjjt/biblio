package handler

import (
	"context"

	userP "github.com/zjjt/biblio/user_service/proto/user"
	"github.com/zjjt/biblio/user_service/repository"
	"golang.org/x/crypto/bcrypt"
)

//Service is the struct that handles all incoming traffic and requests
type Service struct {
	repo repository.Repository
}

//InitService initialize the service handler
func InitService(repo repository.Repository) *Service {
	return &Service{repo}
}

//CreateUser does as its name says
func (s *Service) CreateUser(ctx context.Context, req *userP.User, res *userP.Response) error {
	//generate hash version of user password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Pwd), bcrypt.DefaultCost)
	if err != nil {
		res.Success = false
		res.Error = &userP.Error{
			Code:   400,
			Detail: err.Error(),
		}
		return nil
	}
	req.Pwd = string(hashPass)
	if err := s.repo.CreateUser(req); err != nil {
		res.Success = false
		res.Error = &userP.Error{
			Code:   400,
			Detail: err.Error(),
		}
		return nil
	}
	res.Success = true
	return nil
}

//GetUserByName does as its name says
func (s *Service) GetUserByName(ctx context.Context, req *userP.Request, res *userP.Response) error {
	user, err := s.repo.GetUserByName(req.UserName)
	if err != nil {
		res.Success = false
		res.Error = &userP.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return nil
	}
	// Compare the password with the hashed password stored in database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(req.UserPwd)); err != nil {
		res.Success = false
		res.Error = &userP.Error{
			Code:   500,
			Detail: err.Error(),
		}
		return nil
	}
	res.Success = true
	res.User = user
	return nil
}
