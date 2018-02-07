package router

import (
	"net/http"

	"../handlers"
	"github.com/gin-gonic/gin"
)

// Init blablaba
func Init() {
	// Creates a default gin router
	router := gin.Default() // Grouping routes
	// group： v1
	v1 := router.Group("/v1")
	{
		v1.GET("/hello", handlers.Hello)

		v1.GET("/get", handlers.Get)

		//URL中的name參數
		v1.GET("/hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})

		//常規參數
		v1.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			lastname := c.Query("lastname")
			c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
		})
	}

	//根據website的路由規則
	router.LoadHTMLGlob("templates/*")
	v2 := router.Group("/v2")
	{
		v2.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{"title": "hello Gin."})
		})

	}
	// 404 NotFound
	router.NoRoute(func(c *gin.Context) {
		//https://golang.org/src/net/http/status.go
		c.HTML(200, "404.html", gin.H{})
	})

	router.Run(":8000") // listen and serve on 0.0.0.0:8000
}
