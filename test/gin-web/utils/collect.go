package utils

import (
	"gin-web/types"
	"fmt"
)
var storeData []types.UserInfo

func CollectArray(userArr []types.UserInfo){
	storeData = append(storeData,userArr...)
	fmt.Println("storedata=============11:", storeData)
}

func Collect(user types.UserInfo){
	storeData = append(storeData,user)
	fmt.Println("storedata==============22:", storeData)
}