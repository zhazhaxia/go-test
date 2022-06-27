package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	m := map[string]interface{}{"app_id": "you_api", "app_sign": "md5_base_16", "timestamp": "1473655478000"}
	json_str, _ := json.Marshal(m)
	fmt.Println(string(json_str[:]))

	var values url.Values

	values.Set("test", fmt.Sprint("777"))
	fmt.Println("lb===", values.Get("test"))

	fmt.Println("aaa===", values.Encode())

}
