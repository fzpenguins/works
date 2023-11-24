package serialize

type Response struct {
	Status uint        `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type TokenData struct {
	Data interface{} `json:"data"`
	Id   uint        `json:"id"`
}
