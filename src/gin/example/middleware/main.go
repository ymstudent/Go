package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method

		fmt.Println("请求path: ", path)
		fmt.Println("请求method: ", method)

		context.Next() //此处跳转执行业务逻辑
		fmt.Println(context.Writer.Status())
	}
}

func main()  {
	r := gin.Default()
	r.Use(RequestInfos())

	r.GET("/query", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{} {
			"code": 1,
			"msg": context.FullPath(),
		})
	})

	r.Run(":8081")
}
