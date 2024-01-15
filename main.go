package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/asaba/controller"
	"github.com/asaba/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/asaba_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	print(&db)

	productController := controller.NewProductController(db)
	productLogController := controller.NewProductLogController(db)

	e := echo.New()
	e.Use(middleware.CORS())

	api := e.Group("/api")
	api.POST("/products", func(c echo.Context) error {
		var createProducts models.CreateProducts
		err := c.Bind(&createProducts)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Message: err.Error(),
			})
		}
		err = productController.CreateProducts(context.Background(), createProducts)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Message: fmt.Sprintf("error while creating data : %s", err.Error()),
			})
		}
		return c.JSON(http.StatusOK, models.Response{
			Message: "success",
		})
	})

	api.GET("/products", func(c echo.Context) error {
		var products models.Products
		products, err = productController.GetProducts(context.Background())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Message: fmt.Sprintf("error while getting data : %s", err.Error()),
			})
		}
		return c.JSON(http.StatusOK, models.Response{
			Message: "success",
			Data:    products,
		})
	})

	api.PUT("/products", func(c echo.Context) error {
		var updateProducts models.UpdateProducts
		err := c.Bind(&updateProducts)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Message: err.Error(),
			})
		}

		err = productController.UpdateProducts(context.Background(), updateProducts)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Message: fmt.Sprintf("error while Updating data : %s", err.Error()),
			})
		}
		return c.JSON(http.StatusOK, models.Response{
			Message: "success",
		})
	})

	api.POST("/product-logs", func(c echo.Context) error {
		var productLog models.ProductLog
		err := c.Bind(&productLog)

		if err != nil {
			return c.JSON(http.StatusBadRequest, models.Response{
				Message: err.Error(),
			})
		}

		err = productLogController.CreateProductLog(context.Background(), productLog)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.Response{
				Message: fmt.Sprintf("error while creating Log : %s", err.Error()),
			})
		}

		return c.JSON(http.StatusOK, models.Response{
			Message: "success",
		})
	})

	e.Start(":8000")
}
