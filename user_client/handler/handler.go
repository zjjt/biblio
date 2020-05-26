package handler

import (
	"context"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/gorilla/schema"
	"github.com/micro/go-micro/v2/client"
	userP "github.com/zjjt/biblio/user_service/proto/user"
)

var (
	serviceClient userP.UserService
	decoder       = schema.NewDecoder()
)

//API is a struct used in the rest api
type API struct{}

//Init initializes the user service client
func Init() {
	serviceClient = userP.NewUserService("biblio.service.user", client.DefaultClient)
}

//Signup creates a new user in the system
func (s *API) Signup(req *restful.Request, res *restful.Response) {
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	//make use of anonymous struct
	args := new(struct {
		username string
		password string
	})
	err = decoder.Decode(args, req.Request.PostForm)
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	//we do not care about the response of the GRPC service
	//We only check for an error and if no errors occured we format a json response and send it
	//to the client
	_, err = serviceClient.CreateUser(context.TODO(), &userP.User{Name: args.username, Pwd: args.password})
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	response := struct {
		success bool
		detail  string
	}{true, "the user has been created"}
	res.WriteEntity(response)
}

//Login logs the user in
func (s *API) Login(req *restful.Request, res *restful.Response) {

	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	//make use of anonymous struct
	args := new(struct {
		username string
		password string
	})
	err = decoder.Decode(args, req.Request.PostForm)
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	//we do not care about the response of the GRPC service
	//We only check for an error and if no errors occured we format a json response and send it
	//to the client
	_, err = serviceClient.GetUserByName(context.TODO(), &userP.Request{UserName: args.username, UserPwd: args.password})
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	response := struct {
		success bool
		token   string
		detail  string
	}{true, "", "this user is now connected"}
	res.WriteEntity(response)
}
