package controllers

import (
	"zad4/config"
	"zad4/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return err
	}
	config.DB.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func GetProducts(c echo.Context) error {
	var products []models.Product
	config.DB.Preload("Category").Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	result := config.DB.Preload("Category").First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found."})
	}
	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found."})
	}
	if err := c.Bind(&product); err != nil {
		return err
	}
	config.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	config.DB.Delete(&models.Product{}, id)
	return c.NoContent(http.StatusNoContent)
}
