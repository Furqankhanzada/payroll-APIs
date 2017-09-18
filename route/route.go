package route

import (
	"log"
	"github.com/gin-gonic/gin"
	"payroll-APIs/server"

)

type Users struct {
	Name string
	Phone string
}

var users  []Users

func Route(){
	router := gin.Default()
	router.GET("/api/get/user", func(c *gin.Context){
		UserCollection := server.GetCollection("test")
		err := UserCollection.Find(nil).All(&users)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{"data" : users})
	})
	router.Run(":4000")
}
