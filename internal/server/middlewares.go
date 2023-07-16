package server

import (
	jwtv2 "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/magmel48/social-network/internal/api"
	"github.com/magmel48/social-network/internal/db"
	"strconv"
	"time"
)

var identityKey = "id"

func (h *Handlers) jwt() (*jwtv2.GinJWTMiddleware, error) {
	return jwtv2.New(&jwtv2.GinJWTMiddleware{
		Realm:       "social-network",
		Key:         []byte("secret key"), // FIXME I don't care for that for the first tasks
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var body api.PostLoginJSONBody
			if err := c.ShouldBind(&body); err != nil {
				return nil, jwtv2.ErrMissingLoginValues
			}

			if body.Id == nil || body.Password == nil {
				return nil, jwtv2.ErrMissingLoginValues
			}

			id, err := strconv.Atoi(*body.Id)
			if err != nil {
				return nil, jwtv2.ErrInvalidAuthHeader
			}

			user, err := h.repository.Login(c, db.User{ID: int32(id), Password: *body.Password})
			if err != nil {
				return nil, jwtv2.ErrFailedAuthentication
			}

			return user, nil
		},
	})
}
