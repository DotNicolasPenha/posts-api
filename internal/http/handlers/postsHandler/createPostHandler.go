package postsHandler

import (
	"github.com/DotNicolasPenha/Posts-CRUD/internal/http/responses"
	"github.com/DotNicolasPenha/Posts-CRUD/internal/modules/post"
	"github.com/gin-gonic/gin"
	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

func (h *handler) createPostHandler(ctx *gin.Context) {
	var postToCreate post.CreatePostDTO
	if err := ctx.ShouldBindJSON(&postToCreate); err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	userID := ctx.MustGet("userID").(uuid.UUID)
	postToCreate.AuthorID = userID
	if err := h.service.AddPost(postToCreate); err != nil {
		responses.BadRequest(ctx, err)
		return
	}
	responses.OK_CREATED(ctx, "post created")
}
