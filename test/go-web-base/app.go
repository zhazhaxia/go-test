package main

import (
    "fmt"
    "net/http"
)

func defaultRoute(response http.ResponseWriter, request *http.Request) {
    fmt.Println(request.URL.Path)        //通过 request，执行http请求的相关操作
    response.Write([]byte("这里是默认路由")) //通过 response，执行http的响应相关操作
}

func secontRoute(response http.ResponseWriter, request *http.Request) {
    fmt.Println("二级路由~~~")
    response.Write([]byte("我是二级路由"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", defaultRoute) // 默认路由

    mux.HandleFunc("/second", secontRoute) //二级路由

    fmt.Println("go server is running on port:9000")
    http.ListenAndServe(":9000", mux) //设置监听的端口
}