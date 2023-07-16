package server

import (
	"github.com/gin-gonic/gin"
	"github.com/magmel48/social-network/internal/api"
	"github.com/magmel48/social-network/internal/db"
	"github.com/magmel48/social-network/internal/repositories/users"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	repository *users.Repository
	logger     *zap.Logger
}

func New(repository *users.Repository, logger *zap.Logger) *Server {
	return &Server{repository: repository, logger: logger}
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
	var body api.PostLoginJSONRequestBody
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
	var body api.PostUserRegisterJSONRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if body.Password == nil ||
		body.FirstName == nil ||
		body.SecondName == nil ||
		body.City == nil ||
		body.Biography == nil ||
		(body.Birthdate == nil && body.Age == nil) {
		c.Status(http.StatusBadRequest)
		return
	}

	var birthday time.Time
	if body.Birthdate == nil {
		// rough estimate
		birthday = time.Now().AddDate(-*body.Age, 0, 0)
	}

	user, err := s.repository.Register(c, db.User{
		FirstName: *body.FirstName,
		LastName:  *body.SecondName,
		Password:  *body.Password,
		Birthday:  birthday,
	}, body.City, body.Biography)
	if err != nil {
		s.logger.Error("register error", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id": strconv.Itoa(int(user.ID)),
	})
}

func (s *Server) GetUserSearch(c *gin.Context, params api.GetUserSearchParams) {
	//TODO implement me
	panic("implement me")
}
