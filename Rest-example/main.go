package main

import (
	"cc/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()

	r.GET("/",homeFunc)
	
	r.GET("/ping",pingFunc)
	r.POST("/ping",postPingFunc)

	r.GET("/detail/:id",detailFunc)

	api := r.Group("/ping") 
	{
		v1 := api.Group("/v1") 
		{
			v1.Use(middlewares.CustomMiddleware)
			v1.GET("/a",homeFunc)
			v1.GET("/b",pingFunc)
		}
		v2 := api.Group("/v2") 
		{
			v2.GET("/c",homeFunc)
			v2.GET("/d",pingFunc)
		}
	}
	r.Run(":8080")
}

func homeFunc(context *gin.Context) {
	var data  = gin.H{
		"message" : "Hello from Hieu Hoa Hong",
	}
	context.JSON(http.StatusOK, data)
}

func pingFunc(context *gin.Context) {
	context.String(http.StatusOK,"Ping")
}

func postPingFunc(context *gin.Context) {
	context.JSON(http.StatusOK,gin.H{
		"message" : "Post ping OK",
	})
}
func detailFunc(context *gin.Context) {
	name := context.DefaultQuery("name","Hieu") // get query params
	id := context.Param("id") // get param
	address := context.DefaultPostForm("address", "aaa") // get data from Post form res.body
	context.JSON(http.StatusOK, gin.H{
		"id" : id,
		"message" : "Hello " + name + " from " + address,
	})
	log.Print(("I'm in a local middleware"))
}

// func CustomMiddleware(context *gin.Context) {
// 	log.Print(("I'm in a global middleware"))
// 	if(true) { 
// 		context.Next()
// 	}
// }