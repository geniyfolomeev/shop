package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop"
	"strconv"
)

func (h *Handler) createProduct(c *gin.Context) {
	var input shop.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Product.CreateProduct(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) getAllProducts(c *gin.Context) {
	products, err := h.services.GetAllProducts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) getProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	product, err := h.services.GetProductById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *Handler) updateProduct(c *gin.Context) {

}

func (h *Handler) deleteProduct(c *gin.Context) {

}
