package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

//ShouldBind方法自动进行参数绑定
type Student struct {
	Name 	string `form:"name" binding:"required"`
	Classes string `form:"classes"`
}

type Register struct {
	UserName string `form:"name"`
	Phone 	 string `form:"phone"`
	Password string `form:"pwd"`
}

func main()  {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var student Student
		if err := context.ShouldBind(&student); err != nil {
			log.Println(err.Error())
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)
		context.Writer.Write([]byte("hello, " + student.Name))
	})

	r.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var register Register
		if err := context.ShouldBind(&register); err != nil {
			log.Println(err.Error())
		}

		fmt.Println(register.UserName)
		fmt.Println(register.Phone)
		fmt.Println(register.Password)
		context.Writer.Write([]byte(register.UserName + " register success"))
	})

	r.Run(":8081")
}