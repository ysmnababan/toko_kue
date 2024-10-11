package handler

import (
	"net/http"
	"strconv"
	"toko_kue/helper"
	"toko_kue/models"
	"toko_kue/repository"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CR repository.CategoryRepoI
}

func (h *CategoryHandler) GetAllCategory(e echo.Context) error {
	data, err := h.CR.GetAllCategory()
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Categories fetched successfully",
		"data":    data,
	})
}
func (h *CategoryHandler) GetById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	data, err := h.CR.GetById(id)
	if err != nil {
		return helper.ParseError(err, e)
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Categories fetched successfully",
		"data":    data,
	})
}

func (h *CategoryHandler) AddCategory(e echo.Context) error {
	cat := models.Category{}
	err := e.Bind(&cat)
	if err != nil {
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	if cat.Name == "" || cat.Code == "" {
		return helper.ParseError(helper.ErrParam, e)
	}
	data, err := h.CR.AddCategory(&cat)
	if err != nil {
		return helper.ParseError(err, e)
	}
	return e.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Category created successfully",
		"data":    data,
	})
}

func (h *CategoryHandler) UpdateCategory(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	cat := models.Category{}
	err = e.Bind(&cat)
	if err != nil {
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	if cat.Name == "" || cat.Code == "" {
		return helper.ParseError(helper.ErrParam, e)
	}
	cat.CategoryID = uint(id)

	data, err := h.CR.UpdateCategory(&cat)

	if err != nil {
		return helper.ParseError(err, e)
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Categories updated successfully",
		"data":    data,
	})
}
func (h *CategoryHandler) DeleteCategory(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ParseError(helper.ErrInvalidId, e)
	}

	data, err := h.CR.DeleteCategory(id)
	if err != nil {
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Categories deleted successfully",
		"data":    data,
	})
}
