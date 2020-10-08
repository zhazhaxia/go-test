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
		var userinfo types.UserInfo
		if err := c.ShouldBindJSON(&userinfoJsonArr); err == nil&&len(userinfoJsonArr.ListInfo)>0 {
			utils.CollectArray(userinfoJsonArr.ListInfo)
		} else if err := c.ShouldBindJSON(&userinfo); err == nil{
			utils.Collect(userinfo)
		}else{
			c.JSON(http.StatusBadRequest, gin.H{"error": "Login information is not complete"})
			return
		}
		fmt.Println("1111111111===:", userinfoJsonArr.ListInfo)
		fmt.Println("2222222222===:", userinfo)

		fmt.Println("len===:", len(userinfoJsonArr.ListInfo))
		// 将map格式化
		jresult, err := json.MarshalIndent(userinfoJsonArr, "", "   ")
    if err != nil {
        fmt.Println("%v\n", err)
    }
		fmt.Println("receive value===:", string(jresult))
		// if userinfoJsonArr.UserName != "milu" || userinfoJsonArr.Age != 18 {
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

