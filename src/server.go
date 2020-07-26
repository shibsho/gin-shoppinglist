package main

import (
	"log"
	"net/http"

	controller "projects/practice/gin-shoppinglist/src/controllers/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// サーバーを起動する
	serve()
}

func serve() {
	// デフォルトのミドルウェアでginのルーターを作成
	// LoggerとアプリケーションクラッシュをキャッチするRecoveryミドルウェアを保有している
	router := gin.Default()

	router.StaticFS("/web", http.Dir("sta"))

	// 静的ファイルのパスを設定
	router.Static("/views", "./views")

	// ルーターの設定
	// URLへのアクセスに対して静的ページを返す(localhost/shoppingappにアクセスするとstatic/inex.htmlが返る)
	router.StaticFS("/shoppingapp", http.Dir("views/static"))

	// 全商品情報を返す
	router.GET("/fetchAllProducts", controller.FetchAllProducts)
	// 1つの商品情報の状態を返す
	router.GET("/fetchProduct", controller.FindProduct)
	// 商品登録
	router.POST("/addProduct", controller.AddProduct)
	// 商品更新
	router.POST("/changeStateProduct", controller.ChangeStateProduct)
	// 商品削除
	router.POST("/deleteProduct", controller.DeleteProduct)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Faild: ", err)
	}
}
