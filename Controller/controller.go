package controller

import (
	entities "api/Entities"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB, ctx *gin.Context) []entities.Product {
	var products []entities.Product

	db.Model(entities.Product{}).Joins("Catagory").Find(&products)

	// page index
	product := entities.Product{}
	catagory := entities.Catagory{}
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	// items per request/page
	item, err := strconv.Atoi(ctx.Query("item"))
	if err != nil || item < 0 {
		item = len(products)
	}

	// get filter key
	product.Gender = ctx.Query("gender")
	product.Size = ctx.Query("size")
	catagory.Color = ctx.Query("color")
	catagory.Pattern = ctx.Query("pattern")
	catagory.Figure = ctx.Query("figure")

	products = productsFilter(products, product.Gender, product.Size, catagory)

	// slice products list from item and page
	if item*(page-1) >= len(products) {
		return []entities.Product{}
	}

	if item*page >= len(products) {
		products = products[item*(page-1):]
	} else {
		products = products[item*(page-1) : item*page]
	}

	return products
}

func CreateOrder(db *gorm.DB, ctx *gin.Context) (entities.Order, error) {
	var order entities.Order
	var product entities.Product

	err := ctx.BindJSON(&order)
	if err != nil {
		return entities.Order{}, err
	}

	// check existing product
	err = db.Model(&entities.Product{}).Where("ID = ?", order.Product.ID).First(&product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.Order{}, errors.New("product not found")
	}

	// check product Quantity
	if product.Quantity < order.Quantity {
		return entities.Order{}, errors.New("not enough product")
	}

	// check existing user (if not have user or ID : create it)
	err = db.Model(&entities.User{}).Where("ID = ?", order.User.ID).First(&order.User).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		db.Save(&order.User)
	}

	order.Product = product
	db.Save(&order)
	return order, nil
}

func GetOrders(db *gorm.DB, ctx *gin.Context) []entities.Order {
	var orders []entities.Order
	layout := "2006-01-02"
	db.Model(&orders).Joins("Product").Joins("User").Find(&orders)

	// get filter order
	startDate, err := time.Parse(layout, ctx.Query("start"))
	endDate, err := time.Parse(layout, ctx.Query("end"))

	if err == nil {
		orders = paidOrdersFilter(orders, startDate, endDate)
	}

	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}

	item, err := strconv.Atoi(ctx.Query("items"))
	if err != nil || item < 0 {
		item = len(orders)
	}

	// slice orders list from item and page
	if item*(page-1) >= len(orders) {
		return []entities.Order{}
	}

	if item*page >= len(orders) {
		orders = orders[item*(page-1):]
	} else {
		orders = orders[item*(page-1) : item*page]
	}

	orders = orders[item*(page-1) : item*page]
	return orders
}

func productsFilter(products []entities.Product, gender string, size string, catagory entities.Catagory) []entities.Product {
	var ProductsFiltered []entities.Product

	for _, prod := range products {
		if gender != "" && prod.Gender != gender {
			continue
		}

		if size != "" && prod.Size != size {
			continue
		}

		if catagory.Color != "" && prod.Catagory.Color != catagory.Color {
			continue
		}

		if catagory.Pattern != "" && prod.Catagory.Pattern != catagory.Pattern {
			continue
		}

		if catagory.Figure != "" && prod.Catagory.Figure != catagory.Figure {
			continue
		}
		ProductsFiltered = append(ProductsFiltered, prod)
	}
	return ProductsFiltered
}

func paidOrdersFilter(orders []entities.Order, startDate time.Time, endDate time.Time) []entities.Order {
	var ordersFiltered []entities.Order
	for _, order := range orders {
		if order.Status == "paid" && order.CreatedAt.After(startDate) && order.CreatedAt.Before(endDate) {
			ordersFiltered = append(ordersFiltered, order)
		}
	}
	return ordersFiltered
}
