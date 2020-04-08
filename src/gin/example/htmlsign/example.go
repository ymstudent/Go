package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func formAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

func main()  {
	r := gin.Default()
	r.Delims("{[{", "}]}")
	r.SetFuncMap(template.FuncMap{
		"formAsDate": formAsDate,
	})
	r.LoadHTMLGlob("../templates/*")

	r.GET("/raw", func(c *gin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", gin.H{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})
	r.Run(":8081")
}
