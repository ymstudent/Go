package main


//http2 server push 相关介绍：https://blog.golang.org/h2push
import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
	<title>http test</title>
	<script src="/assets/app.js"></script>
</head>
<body>
	<h1 style="color:red;">welcome, Ginner</h1>
</body>
</html>
`))

func main()  {
	r := gin.Default()
	r.Static("../assets", "../assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(context *gin.Context) {
		if pusher := context.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("../assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		context.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})

	r.RunTLS(":8081", "../testdata/server.pem", "../testdata/server.key")
}
