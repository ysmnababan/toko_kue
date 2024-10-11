package handler

import (
	"net/http"
	"toko_kue/repository"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CR repository.CategoryRepoI
}

func (h *CategoryHandler) GetAllCategory(e echo.Context) error {
	
	return e.JSON(http.StatusOK, "")
}
func (h *CategoryHandler) GetById(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}
func (h *CategoryHandler) AddCategory(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}
func (h *CategoryHandler) UpdateCategory(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}
func (h *CategoryHandler) DeleteCategory(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}
