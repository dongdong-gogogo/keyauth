package api

import (
	"github.com/dongdong-gogogo/keyauth/utils"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"

	"github.com/dongdong-gogogo/keyauth/apps/token"
)

func (h *handler) IssueToken(r *restful.Request, w *restful.Response) {

	req := token.NewIssueTokenRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

func (h *handler) ValidateToken(r *restful.Request, w *restful.Response) {
	accessToken := utils.GetToken(r.Request)

	req := token.NewValidateTokenRequest(accessToken)
	ins, err := h.service.ValidateToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, ins)
}

func (h *handler) RevolkToken(r *restful.Request, w *restful.Response) {
	req := token.NewRevolkTokenRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	set, err := h.service.RevolkToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}
