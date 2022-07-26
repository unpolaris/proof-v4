package model

type ReqAutoProof struct {
	Id     string        `json:"id"`
	Detail string        `json:"detail" validate:"required" label:"detail 存证内容"`
	Ext    *ProofExtInfo `json:"ext"`
}

type ProofExtInfo struct {
	Note        string           `json:"note"`
	Ext         string           `json:"ext"`
	IsIncrement bool             `json:"is_increment"`
	TemplateId  int64            `json:"template_id"`
	ExtDetail   *ProofExtDetail  `json:"ext_detail"`
	NoteDetail  *ProofNoteDetail `json:"note_detail"`
	Version     string           `json:"version"`
	Prikey      string           `json:"prikey"` //自定义私钥
}

type ProofExtDetail struct {
	BaseHash  string `json:"basehash"`
	PreHash   string `json:"prehash"`
	ProofName string `json:"存证名称"`
	ProofType string `json:"存证类型"`
}

type ProofNoteDetail struct {
	Version      int            `json:"version"`
	UserName     string         `json:"userName"`
	UserIcon     string         `json:"userIcon"`
	EvidenceName string         `json:"evidenceName"` //基础存证的名字
	StepName     string         `json:"stepName"`     //增量存证的名字
	Banners      []*ProofBanner `json:"banners"`      //增量不用，基础可能要
}


type ProofBanner struct {
	Hash   string `json:"hash"`
	Target string `json:"target"`
}
