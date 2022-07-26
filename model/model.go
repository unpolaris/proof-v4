package model

// Response 返回数据的结构体
type Response struct {
	Code int         `json:"code" example:"200" extensions:"x-order=000"` //状态码
	Msg  string      `json:"msg" example:"OK" extensions:"x-order=001"`   //消息
	Data interface{} `json:"data" extensions:"x-order=002"`               //数据
}

type ResAutoProof struct {
	Suc  []*AutoProofIdHash `json:"suc"`
	Fail []*AutoProofIdErr  `json:"fail"`
}

type AutoProofIdHash struct {
	Id     string `json:"id"`
	TxHash string `json:"tx_hash"` //自动存证hash
}

type AutoProofIdErr struct {
	Id  string `json:"id"`
	Err string `json:"err"` //错误原因
}

