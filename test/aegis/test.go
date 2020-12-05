package main

import (
	"fmt"

	"git.code.oa.com/aegis/aegis-go-sdk/aegis"
	"git.code.oa.com/aegis/aegis-go-sdk/aegis/models"
	// "git.code.oa.com/aegis/aegis-go-sdk/aegis/services"
)

func main() {
	var a models.Aegis
	fmt.Println(a)
	fmt.Println(models.Aegis)
	fmt.Println(aegis.Info)
	// services.DoRequest("t")

}
