package handler

import (
	"net/http"
	"strconv"
	"toko_kue/helper"
	"toko_kue/models"
	"toko_kue/repository"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	PR repository.ProductRepoI
}

func (h *ProductHandler) GetAllProduct(e echo.Context) error {
	data, err := h.PR.GetAllProduct()
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product fetched successfully",
		"data":    data,
	})
}
func (h *ProductHandler) GetProductById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	data, err := h.PR.GetProductById(id)
	if err != nil {
		return helper.ParseError(err, e)
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product fetched successfully",
		"data":    data,
	})
}

func (h *ProductHandler) AddProduct(e echo.Context) error {
	cat := models.Product{}
	err := e.Bind(&cat)
	if err != nil {
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	if cat.Name == "" || cat.Code == "" || cat.Price < 0 || cat.Stock < 0 {
		return helper.ParseError(helper.ErrParam, e)
	}

	data, err := h.PR.AddProduct(&cat)
	if err != nil {
		return helper.ParseError(err, e)
	}
	return e.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Product created successfully",
		"data":    data,
	})
}

func (h *ProductHandler) UpdateProduct(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	cat := models.Product{}
	err = e.Bind(&cat)
	if err != nil {
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	if cat.Name == "" || cat.Code == "" {
		return helper.ParseError(helper.ErrParam, e)
	}
	cat.ID = uint(id)
	data, err := h.PR.UpdateProduct(&cat)

	if err != nil {
		return helper.ParseError(err, e)
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product updated successfully",
		"data":    data,
	})
}
func (h *ProductHandler) DeleteProduct(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	data, err := h.PR.DeleteProduct(id)
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Product deleted successfully",
		"data":    data,
	})
}
