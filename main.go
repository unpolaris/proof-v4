package main

import (
	"github.com/gin-gonic/gin"
)

// @Title 存证溯源接口
// @Version 0.1
// @Description 尽量通用的存证溯源平台，目前是简单用户体系。
// @SecurityDefinitions.ApiKey ApiKeyAuth
// @In header
// @Name Authorization
// @BasePath /
func main() {
	router := gin.Default()

	router.POST("/api/auto_proofs", AutoProofs)
	router.POST("/v1/proof/GetProofs", GetProofs)
	router.POST("/v1/proof/List", ListProof)
	router.Run(":5000")
}

// AutoProofs 添加存证信息
// @Tags 存证 proof
// @Summary 添加批量自动存证
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param FZM-Ca-App-Id header string false "应用id"
// @Param FZM-Ca-App-Key header string false "应用公钥"
// @Param FZM-Ca-Sign header string false "签名"
// @Param FZM-Ca-Nonce header string false "随机字符串"
// @Param FZM-Ca-Timestamp header string false "时间戳"
// @Param FZM-Ca-Sign-Type header string false "签名方式，HMAC-SHA256（默认）或MD5"
// @Param data body []model.ReqAutoProof true "添加批量自动存证"
// @Success 200 {object} model.Response{data=model.ResAutoProof} "{"code":200,"msg":"OK","data":{}}"
// @Router /api/auto_proofs [post]
func AutoProofs(c *gin.Context) {}

// GetProofs 获取多个指定hash的存证信息
// @Summary 获取多个指定hash的存证信息
// @Description get proof by hashes
// @Tags Proof
// @Accept application/json
// @Produce application/json
// @Param input body swagger.ClientRequest{params=[]rpcutils.Hashes} true "INPUT"
// @Success 200 {object} swagger.ListProofResult
// @Failure 400 {object} swagger.ServerResponse{error=string}
// @Router /v1/proof/GetProofs [post]
func GetProofs(c *gin.Context) {}

// ListProof 获取存证列表
// @Summary 获取存证列表
// @Description list proof of organization/sender
// @Tags Proof
// @Accept application/json
// @Produce application/json
// @Param input body swagger.ClientRequest{params=[]swagger.Query} true "INPUT"
// @Success 200 {object} swagger.ListProofResult
// @Failure 400 {object} swagger.ServerResponse{error=string}
// @Router /v1/proof/List [post]
func ListProof(c *gin.Context) {}
