package http

import (
	"net/http"

	"github.com/adityarizkyramadhan/synapsis-test/user-service/internal/service"
	"github.com/adityarizkyramadhan/synapsis-test/user-service/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	serviceUser *service.User
}

func NewUser(serviceUser *service.User) *User {
	return &User{serviceUser: serviceUser}
}

func (u *User) GetByID(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)

	user, err := u.serviceUser.GetByID(ctx, id)
	if err != nil {
		ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, user)
}
