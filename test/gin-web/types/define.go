package types

type UserInfo struct {
	UserName     string `json:"username"`
	Age uint `json:"age""`
}

type UserList struct {
	ListInfo     []UserInfo `json:"listinfo"`
}