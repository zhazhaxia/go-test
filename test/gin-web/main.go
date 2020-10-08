package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName     string `json:"username"`
	Age uint `json:"age""`
}

func main() {
	router := gin.New()

	router.POST("/loginTest", func(c *gin.Context) {
		var userinfoJson UserInfo
		if err := c.ShouldBindJSON(&userinfoJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Login information is not complete"})
			return
		}
		// 将map格式化
		jresult, err := json.MarshalIndent(userinfoJson, "", "   ")
    if err != nil {
        fmt.Println("%v\n", err)
		}
		
		fmt.Println("receive value===:", string(jresult))
		// if userinfoJson.UserName != "milu" || userinfoJson.Age != 18 {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		// 	return
		// }
		c.JSON(http.StatusOK, gin.H{"status": "It's OK"})
	})

	router.GET("/loginTest", func(c *gin.Context) {
			firstname := c.DefaultQuery("username", "DefaultName")
			lastname := c.Query("age") // 是 c.Request.URL.Query().Get("lastname") 的简写

			c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":9000")
}

