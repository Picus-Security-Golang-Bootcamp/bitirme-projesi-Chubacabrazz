package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	repo *CategoryRepository
}

func CategoryHandler(r *gin.RouterGroup, repo *CategoryRepository) {
	h := &categoryHandler{repo: repo}

	r.GET("/", h.getAll)
	/* r.POST("/create", h.create)
	r.GET("/:id", h.getByID)
	r.PUT("/:id", h.update)
	r.DELETE("/:id", h.delete) */
}

func (b *categoryHandler) getAll(c *gin.Context) {
	categories, err := b.repo.GetAll()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, categories)
}
