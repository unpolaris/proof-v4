package swagger

type ListProofResult struct {
	ServerResponse
	Result []Proof `json:"result"` // 返回结果
}

type Proof struct {
	Basehash           string      `json:"basehash"`                       // 增量存证依赖的主hash
	EvidenceName       string      `json:"evidenceName"`                   // 存证名称
	Prehash            string      `json:"prehash"`                        // 增量存证前一个hash
	ProofBlockHash     string      `json:"proof_block_hash"`               // 区块hash
	ProofBlockTime     int         `json:"proof_block_time"`               // 上链时间
	ProofData          string      `json:"proof_data"`                     // 存证数据
	ProofDeleted       string      `json:"proof_deleted"`                  // 删除存证交易hash
	ProofDeletedFlag   bool        `json:"proof_deleted_flag"`             // 删除标志
	ProofDeletedNote   string      `json:"proof_deleted_note"`             // 删除备注
	ProofHeight        int         `json:"proof_height"`                   // 存证高度
	ProofHeightIndex   int         `json:"proof_height_index"`             //  存证高度索引
	ProofId            string      `json:"proof_id"`                       // 存证id
	ProofNote          string      `json:"proof_note"`                     // 存证备注
	ProofOrganization  string      `json:"proof_organization"`             // 组织
	ProofOriginal      string      `json:"proof_original"`                 // 来源
	ProofSender        string      `json:"proof_sender"`                   // 存证发起者
	ProofTxHash        string      `json:"proof_tx_hash"`                  // 交易哈希
	SourceHash         interface{} `json:"source_hash"`                    // 依赖交易哈希
	UpdateHash         string      `json:"update_hash"`                    // 更新依赖主哈希
	UpdateVersion      int         `json:"update_version"`                 // 更新版本
	Version            int         `json:"version"`                        // 存证版本
	UserAuthType       int         `json:"user_auth_type,omitempty"`       // 用户认证类型
	UserEmail          string      `json:"user_email,omitempty"`           // 用户邮箱
	UserIcon           string      `json:"user_icon,omitempty"`            // 用户头像链接地址
	UserName           string      `json:"user_name,omitempty"`            // 用户名
	UserPhone          string      `json:"user_phone,omitempty"`           // 用户手机号
	UserRealName       string      `json:"user_real_name,omitempty"`       // 用户真是名称
	UserEnterpriseName string      `json:"user_enterprise_name,omitempty"` // 用户企业名称
}

