package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

var codeStr = "nL_e6rYu-Ixpo7NDva0K429hCQHOSVUJcmjwE1T8t3ybgZGBz5AiPkslFqRfWMdX"

// EncodeStrUin 编码字符串格式的uin
func EncodeStrUin(uin string) string {
	return base64.NewEncoding(codeStr).WithPadding('*').EncodeToString([]byte(uin))
	// return base64.StdEncoding.EncodeToString([]byte(uin))
}

func EncodeStrUinComm(uin string) string {
	// return base64.NewEncoding(codeStr).WithPadding('*').EncodeToString([]byte(uin))
	return base64.StdEncoding.EncodeToString([]byte(uin))
}

// EncodeUin 编码uin
func EncodeUin(uin int64) string {
	return base64.NewEncoding(codeStr).WithPadding('*').EncodeToString([]byte(strconv.FormatInt(uin, 10)))
}

// DecodeUin 解码Uin
func DecodeUin(enUin string) int64 {
	strUin, err := base64.NewEncoding(codeStr).WithPadding('*').DecodeString(enUin)
	if err != nil {
		fmt.Println("decode error:", err)
		return 10
	}
	uin, _ := strconv.ParseInt(string(strUin), 10, 64)
	return uin
}

// DecodeUin 解码Uin 返回str
func DecodeUinStr(enUin string) string {
	strUin, err := base64.NewEncoding(codeStr).WithPadding('*').DecodeString(enUin)
	if err != nil {
		fmt.Println("decode error:", err)
		return ""
	}
	return string(strUin)
}

func main() {
	fmt.Println("uin=====", 12345789)
	fmt.Println("encode c429505189=====", EncodeStrUinComm("4295051879"))
	fmt.Println("encode 4295051879=====", EncodeStrUin("4295051879"))
	fmt.Println("encode int=====", EncodeUin(12345789))
	fmt.Println("encode c=====", EncodeStrUinComm("45644646467474564464646747"))
	fmt.Println("encode=====", EncodeStrUin("45644646467474564464646747"))
	fmt.Println("decode int=====", DecodeUin("7e-q7KnkoKcINv**"))
	fmt.Println("decode=====", DecodeUinStr("7e-q7KnkoKcINv**"))
	fmt.Println("decode=====", DecodeUinStr("7e4s7evs7eCP7wSP7ivk7wvP7wvs7eCl7eS*"))
}
