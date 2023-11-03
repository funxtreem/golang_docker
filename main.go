package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(mysql.Open("root:cakratendados@tcp(172.17.0.2:3306)/db_coba"))
	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Go Docker Running")
		fmt.Println("Connection Established")
		fmt.Println("This Service Running on Port: 9000")
		HandleRequest()
	}

}

func SelectTable1(c *gin.Context) {
	type TableStruct struct {
		Id     uint64 `json:"id"`
		Number int    `json:"number"`
		Text   string `json:"string"`
	}
	var DataOutput []TableStruct

	result := db.Table("table1").Select("*").Scan(&DataOutput)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, DataOutput)
	} else {
		c.JSON(http.StatusOK, "Error")
	}

}

func SelectTable2(c *gin.Context) {
	type TableStruct struct {
		Code string `json:"code"`
		Name string `json:"name"`
	}
	var DataOutput []TableStruct

	result := db.Table("table2").Select("*").Scan(&DataOutput)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, DataOutput)
	} else {
		c.JSON(http.StatusOK, "Error")
	}

}

func HandleRequest() {
	myRouter := gin.Default()
	myRouter.GET("/SelectTable1", SelectTable1)
	myRouter.GET("/SelectTable2", SelectTable2)

	err := myRouter.Run(":9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
