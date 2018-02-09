package handlersVue

import (
	"net/http"

	"../../../mssql"
	"github.com/gin-gonic/gin"
)

//GetData handler
func GetData(c *gin.Context) {
	mssql.Init()
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{
		"student": "Student",
	})
}
