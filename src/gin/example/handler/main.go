package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.Handle("GET", "/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		name := context.DefaultQuery("name", "")
		fmt.Println(name)

		context.Writer.Write([]byte("hello," + name))
	})

	r.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		user := context.DefaultPostForm("user", "")
		fmt.Println(user)
		password := context.PostForm("password")
		fmt.Println(password)

		context.Writer.Write([]byte("user login"))
	})

	r.DELETE("/user/:id", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		id := context.Param("id")
		fmt.Println(id)

		context.Writer.Write([]byte("delete user id: " + id))
	})

	r.Run(":8081")
}
