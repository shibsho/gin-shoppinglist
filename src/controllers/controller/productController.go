package controller

import (
	"strconv"

	entity "projects/practice/gin-shoppinglist/src/models/entity"

	db "projects/practice/gin-shoppinglist/src/models/db"

	"github.com/gin-gonic/gin"
)

// 商品の購入状態
const (
	notPurchased = 0
	Purchased    = 1
)

func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()
	c.JSON(200, resultProducts)
}

func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")
	productID, _ := strconv.Atoi(productIDStr)
	product := db.FindProduct(productID)
	c.JSON(200, product)
}

func AddProduct(c *gin.Context) {
	productName := c.PostForm("productName")
	productMemo := c.PostForm("productMemo")

	product := entity.Product{
		Name:  productName,
		Memo:  productMemo,
		State: notPurchased,
	}

	db.InsertProduct(&product)
}

func ChangeStateProduct(c *gin.Context) {
	reqProductID := c.PostForm("productID")
	reqProductState := c.PostForm("productState")

	productID, _ := strconv.Atoi(reqProductID)
	productState, _ := strconv.Atoi(reqProductState)

	changeState := notPurchased
	// 商品が未購入状態の場合、購入状態に変換する
	if productState == notPurchased {
		changeState = Purchased
	}

	db.UpdateStateProduct(productID, changeState)
}

func DeleteProduct(c *gin.Context) {
	productIDStr := c.PostForm("productID")

	productID, _ := strconv.Atoi(productIDStr)
	db.DeleteProduct(productID)
}
