// デバッグ用コード
package main

import (
	"log"
	"src/model"
	"src/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db := model.Connect()
	db = db.Debug()
	
	router := gin.Default()

    router.POST("/login", controller.login)
    router.GET("/logout", controller.logout)
    router.GET("/sign-up", controller.signUp)
    router.GET("/ranking", controller.showRanking)
    router.GET("/score", controller.getScore)
    router.GET("/result", controller.getResult)
	router.GET("/test", controller.test)

    if err := router.Run(); err != nil {
        log.Fatal("Server Run Failed.: ", err)
    }

}
