package auth

import (
	"gitee.com/dongdong-0421/keyauth/apps/token"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"

	"gitee.com/dongdong-0421/keyauth/utils"
)

// 通过 r.Filter()  加载中间件
func (a *KeyauthAuther) RestfulAuthHandlerFunc(
	req *restful.Request,
	resp *restful.Response,
	chain *restful.FilterChain,
) {

	// 1、认证中间件，获取用户的Token
	tkStr := utils.GetToken(req.Request)

	// 2、到用户中心验证Token合法性
	validateReq := token.NewValidateTokenRequest(tkStr)
	tk, err := a.auth.ValidateToken(req.Request.Context(), validateReq)
	if err != nil {
		response.Failed(resp.ResponseWriter, err)
		return
	}
	req.SetAttribute("token", tk)
	// chain 用于将请求flow下去
	chain.ProcessFilter(req, resp)
}
