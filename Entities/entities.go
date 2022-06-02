package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
}

type Product struct {
	gorm.Model
	Name       string   `josn:"name" gorm:"column:name"`
	Price      int      `json:"price" gorm:"column:price"`
	Quantity   int      `json:"quantity" gorm:"column:quantity"`
	Size       string   `json:"size" gorm:"column:size"`
	Gender     string   `json:"gender" gorm:"column:gender"`
	Catagory   Catagory `json:"catagory" gorm:"foreignKey:CatagoryID;references:ID"`
	CatagoryID uint64
}

type Catagory struct {
	gorm.Model
	Color   string `json:"color" gorm:"column:color"`
	Pattern string `json:"pattern" gorm:"column:pattern"`
	Figure  string `json:"figure" gorm:"column:figure"`
}

type Cart struct {
	gorm.Model
	User      User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	UserID    uint64
	ProductID uint64
}

type Order struct {
	gorm.Model
	Product   Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	Quantity  int     `json:"quantity" gorm:"column:quantity"`
	Status    string  `json:"status" gorm:"column:status"`
	User      User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
	ProductID uint64
	UserID    uint64
}
