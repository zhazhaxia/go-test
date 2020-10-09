package main

import (
	"gin-web/types"
	"gin-web/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	router := gin.New()

	router.POST("/loginTest", func(c *gin.Context) {
		var userinfoJsonArr types.UserList
		//绑定两种类型，但是性能会降低
		if err := c.ShouldBindJSON(&userinfoJsonArr); err == nil&&len(userinfoJsonArr.ListInfo)>0 {
			utils.CollectArray(userinfoJsonArr.ListInfo)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Login information is not complete"})
			return
		}

		// 将map格式化
		jresult, err := json.MarshalIndent(userinfoJsonArr, "", "   ")
    if err != nil {
        fmt.Println("%v\n", err)
    }
		fmt.Println("receive value===:", string(jresult))
		c.JSON(http.StatusOK, gin.H{"status": "It's OK"})
	})

	router.GET("/loginTest", func(c *gin.Context) {
			firstname := c.DefaultQuery("username", "DefaultName")
			lastname := c.Query("age") // 是 c.Request.URL.Query().Get("lastname") 的简写

			c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.Run(":9000")
}

