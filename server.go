package main

import (
	controller "api/Controller"
	database "api/Database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	db, err := database.New()

	if err != nil {
		fmt.Print("database connection fail.")
	} else {
		fmt.Print("database connection successful.")
	}

	router.GET("/products", func(ctx *gin.Context) {
		products := controller.GetProducts(db, ctx)
		ctx.JSON(http.StatusOK, products)
	})

	router.POST("/order", func(ctx *gin.Context) {
		order, err := controller.CreateOrder(db, ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			ctx.JSON(http.StatusOK, order)
		}
	})

	router.GET("/orders", func(ctx *gin.Context) {
		orders := controller.GetOrders(db, ctx)
		ctx.JSON(http.StatusOK, orders)
	})

	router.Run(":8080")
}
