package user

import (
	"github.com/gin-gonic/gin"
)

func NewUserHandler(r *gin.RouterGroup, repo *UserRepository) {

	r.POST("/register", Register)
	r.POST("/login", Login)
}

/* func (b *bookHandler) create(c *gin.Context) {
	bookBody := &api.Book{}
	if err := c.Bind(&bookBody); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.CannotBindGivenData))
		return
	}

	if err := bookBody.Validate(strfmt.NewFormats()); err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	book, err := b.repo.create(responseToBook(bookBody))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, BookToResponse(book))
}

func (b *bookHandler) getAll(c *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(c)

	books, count := b.repo.getAll(pageIndex, pageSize)

	paginatedResult := pagination.NewFromGinRequest(c, count)
	paginatedResult.Items = booksToResponse(books)

	c.JSON(http.StatusOK, paginatedResult)
} */
