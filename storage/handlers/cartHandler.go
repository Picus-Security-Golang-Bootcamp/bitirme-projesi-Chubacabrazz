package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Chubacabrazz/picus-storeApp/storage/services"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type cartHandler struct {
	service services.Service
}

func NewBasketHandler(r *gin.RouterGroup, service services.Service) {
	h := &cartHandler{service: service}

	r.GET("/:ID", h.getBasket)
	r.POST("/", h.createBasket)
	r.DELETE("/:ID", h.deleteBasket)

	r.POST("/item", h.addItem)
	r.DELETE("/:ID/item/:ProductID", h.deleteItem)
	r.PUT(":ID/item/:item/quantity/:quantity", h.updateItem)
}

func (r *cartHandler) getBasket(g *gin.Context) {
	ID := g.Param("ID")
	result, err := r.service.Get(g.Request.Context(), ID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
	}
	if result == nil {
		g.JSON(http.StatusNotFound, "")
	}
	g.JSON(http.StatusOK, result)
}

func (r *cartHandler) createBasket(g *gin.Context) {
	entity := new(CreateBasketRequest)

	if err := g.Bind(entity); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}

	if b, err := r.service.Create(g.Request.Context(), entity.Buyer); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	} else {

		g.JSON(http.StatusCreated, map[string]string{"ID": b.ID})
	}
}

func (r *cartHandler) deleteBasket(g *gin.Context) {
	ID := g.Param("ID")
	_, err := r.service.Delete(g.Request.Context(), ID)

	if errors.Cause(err) == sql.ErrNoRows {
		g.JSON(http.StatusNotFound, err.Error())
	}

	if err != nil {
		g.JSON(http.StatusInternalServerError, err.Error())
	}

	g.JSON(http.StatusAccepted, "")

}
func (r *cartHandler) addItem(g *gin.Context) {
	req := new(AddItemRequest)

	if err := g.Bind(req); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}

	if ProductID, err := r.service.AddItem(g.Request.Context(), req.BasketId, req.Sku, req.Quantity, req.Price); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	} else {
		g.JSON(http.StatusCreated, map[string]string{"ID": ProductID})
	}
}
func (r *cartHandler) updateItem(g *gin.Context) {

	ID := g.Param("ID")
	ProductID := g.Param("ProductID")
	quantity, err := strconv.Atoi(g.Param("quantity"))

	if len(ID) == 0 || len(ProductID) == 0 || err != nil || quantity <= 0 {
		g.JSON(http.StatusBadRequest, "Failed to update item. BasketId or BasketItem Id is null or empty.")
	}
	if err := r.service.UpdateItem(g.Request.Context(), ID, ProductID, quantity); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}
	g.JSON(http.StatusAccepted, "")

}

func (r *cartHandler) deleteItem(g *gin.Context) {

	ID := g.Param("ID")
	ProductID := g.Param("ProductID")

	if len(ID) == 0 || len(ProductID) == 0 {
		g.JSON(http.StatusBadRequest, "Failed to delete item. BasketId or BasketItem Id is null or empty.")
	}
	if err := r.service.DeleteItem(g.Request.Context(), ID, ProductID); err != nil {
		g.JSON(http.StatusBadRequest, err.Error())
	}
	g.JSON(http.StatusAccepted, "")
}

type (
	CreateBasketRequest struct {
		Buyer string `json:"buyer" validate:"required"`
	}

	AddItemRequest struct {
		BasketId string  `json:"basketId"  validate:"required"`
		Sku      string  `json:"sku"  validate:"required"`
		Quantity int     `json:"quantity" validate:"required,gte=0,lte=20"`
		Price    float64 `json:"price" validate:"required,gte=0"`
	}
)
