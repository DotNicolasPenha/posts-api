package usersHandler

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/responses"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/user"
	"github.com/gin-gonic/gin"
)

func (h *handler) loginHandler(ctx *gin.Context) {
	var LoginUserDTO user.LoginUserDTO
	ctx.ShouldBindBodyWithJSON(&LoginUserDTO)
	token, err := h.service.LoginUser(LoginUserDTO)
	if err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	responses.OK_DATA(ctx, gin.H{
		"token": token,
	})
}
