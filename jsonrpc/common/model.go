package common

type Args struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Reply struct {
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}