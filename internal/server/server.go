package server

import (
	"github.com/gin-gonic/gin"
	"github.com/magmel48/social-network/internal/api"
	"github.com/magmel48/social-network/internal/db"
	"github.com/magmel48/social-network/internal/repositories/users"
	"net/http"
	"strconv"
)

type Server struct {
	repository *users.Repository
}

func New(repository *users.Repository) *Server {
	return &Server{repository: repository}
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
	var body api.PostLoginJSONBody
	err := c.BindJSON(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if body.Id == nil || body.Password == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(*body.Id)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := s.repository.Login(c, db.User{ID: int32(id), Password: *body.Password})
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": strconv.Itoa(int(user.ID)),
	})
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
