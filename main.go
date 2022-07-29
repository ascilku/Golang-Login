package main

import (
	"fmt"
	"go-4/handler"
	"go-4/member"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		newRepository := member.NewRepository(db)
		newSevices := member.NewSevices(newRepository)
		newHandler := handler.NewHandler(newSevices)

		routes := gin.Default()
		api := routes.Group("api/v1")
		api.POST("member", newHandler.SaveHandler)
		api.POST("login", newHandler.SaveLogin)
		routes.Run(":6000")

	}
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// routes := gin.Default()
	// api := routes.Group("api/v1")
	// api.GET("/member", handler)
	// routes.Run(":6000")

}

// func handler(h *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/tes_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		var keyMain []member.Member
// 		db.Find(&keyMain)
// 		h.JSON(http.StatusOK, keyMain)
// 	}
// }
