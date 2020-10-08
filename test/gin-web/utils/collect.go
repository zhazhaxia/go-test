package utils

import (
	"gin-web/types"
	"time"
	"fmt"
)
const maxBufferSize = 6
const timeout = 5 * time.Second

var timer *(time.Timer)
var storeData []types.UserInfo


func CollectArray(userArr []types.UserInfo){
	storeData = append(storeData,userArr...)
	// fmt.Println("storedata=============11:", storeData)
	ActToSend(storeData)
}

func Collect(user types.UserInfo){
	storeData = append(storeData,user)
	// fmt.Println("storedata==============22:", storeData)
	ActToSend(storeData)
}

func ActToSend(userArr []types.UserInfo){
	length := len(userArr)
	if(length >= maxBufferSize){
		fmt.Println("max!!!!!!!!=========before",len(storeData))
		storeData = storeData[0:0]
		fmt.Println("max!!!!!!!!=========",len(storeData))
		return
	}
	fmt.Println("timer data!!!=========timer:",timer)
	if timer != nil {
		return
	}
	timer = time.AfterFunc(timeout, func() {
		timer = nil
		fmt.Println("time!!!!!!!!=========before",len(storeData))
		storeData = storeData[0:0]
		fmt.Println("time!!!!!!!!=========",len(storeData))
	})
}