package server

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/magmel48/social-network/internal/api"
	"github.com/magmel48/social-network/internal/db"
	"github.com/magmel48/social-network/internal/repositories/cities"
	"github.com/magmel48/social-network/internal/repositories/user_cities"
	"github.com/magmel48/social-network/internal/repositories/users"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type Handlers struct {
	users      *users.Repository
	cities     *cities.Repository
	userCities *user_cities.Repository
	logger     *zap.Logger
}

func New(
	users *users.Repository,
	cities *cities.Repository,
	userCities *user_cities.Repository,
	logger *zap.Logger) *Handlers {
	return &Handlers{users: users, cities: cities, userCities: userCities, logger: logger}
}

func (h *Handlers) GetDialogUserIdList(c *gin.Context, userId api.UserId) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) PostDialogUserIdSend(c *gin.Context, userId api.UserId) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) PutFriendDeleteUserId(c *gin.Context, userId api.UserId) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) PutFriendSetUserId(c *gin.Context, userId api.UserId) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) PostLogin(c *gin.Context) {
	// FIXME resolve as options pattern?
	m, _ := h.jwt()
	m.LoginHandler(c)
}

func (h *Handlers) PostPostCreate(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) PutPostDeleteId(c *gin.Context, id api.PostId) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) GetPostFeed(c *gin.Context, params api.GetPostFeedParams) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) GetPostGetId(c *gin.Context, id api.PostId) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) PutPostUpdate(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) GetUserGetId(c *gin.Context, id api.UserId) {
	dbID, err := strconv.Atoi(id)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	user, err := h.users.FindByID(c, int32(dbID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	userID := strconv.Itoa(int(user.ID))
	age := int(time.Now().Sub(user.Birthday).Seconds() / 31207680)

	response := api.User{
		Id:         &userID,
		FirstName:  &user.FirstName,
		SecondName: &user.LastName,
		Birthdate:  &api.BirthDate{Time: user.Birthday},
		Age:        &age,
	}

	if user.Biography.Valid {
		response.Biography = &user.Biography.String
	}

	userCity, err := h.userCities.FindByUserID(c, user.ID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			c.Status(http.StatusInternalServerError)
			return
		}
	} else {
		city, err := h.cities.FindByID(c, userCity.CityID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		response.City = &city.Name
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handlers) PostUserRegister(c *gin.Context) {
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

	user, err := h.users.Register(c, db.User{
		FirstName: *body.FirstName,
		LastName:  *body.SecondName,
		Password:  *body.Password,
		Birthday:  birthday,
	}, body.City, body.Biography)
	if err != nil {
		h.logger.Error("register user error", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}

	if body.City != nil {
		city, err := h.cities.UpsertCity(c, *body.City)
		if err != nil {
			h.logger.Error("upsert city error", zap.Error(err))
			c.Status(http.StatusInternalServerError)
			return
		}

		err = h.userCities.Create(c, user.ID, city.ID)
		if err != nil {
			h.logger.Error("create user <-> city connection error", zap.Error(err))
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id": strconv.Itoa(int(user.ID)),
	})
}

func (h *Handlers) GetUserSearch(c *gin.Context, params api.GetUserSearchParams) {
	//TODO implement me
	panic("implement me")
}
