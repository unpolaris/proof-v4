package swagger

type ClientRequest struct {
	Method string      `json:"method,omitempty"` // 方法
	Params interface{} `json:"params"`           // 参数
	ID     uint64      `json:"id"`               // 请求标识
}

type ServerResponse struct {
	ID     uint64      `json:"id"`     // 请求标识
	Result interface{} `json:"result"` // 返回结果
	Error  interface{} `json:"error"`  // 错误描述
}
