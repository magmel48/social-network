package server

import (
	"github.com/gin-gonic/gin"
	"github.com/magmel48/social-network/internal/api"
)

type Server struct{}

func New() *Server {
	return &Server{}
}

func (s *Server) GetDialogUserIdList(c *gin.Context, userId api.UserId) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostDialogUserIdSend(c *gin.Context, userId api.UserId) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PutFriendDeleteUserId(c *gin.Context, userId api.UserId) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PutFriendSetUserId(c *gin.Context, userId api.UserId) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostLogin(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostPostCreate(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PutPostDeleteId(c *gin.Context, id api.PostId) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetPostFeed(c *gin.Context, params api.GetPostFeedParams) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetPostGetId(c *gin.Context, id api.PostId) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PutPostUpdate(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetUserGetId(c *gin.Context, id api.UserId) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PostUserRegister(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetUserSearch(c *gin.Context, params api.GetUserSearchParams) {
	//TODO implement me
	panic("implement me")
}
