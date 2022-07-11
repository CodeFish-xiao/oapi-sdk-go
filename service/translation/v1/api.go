// Package translation code generated by oapi sdk gen
package larktranslation

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *TranslationService {
	t := &TranslationService{config: config}
	t.Text = &text{service: t}
	return t
}

// 业务域服务定义
type TranslationService struct {
	config *larkcore.Config
	Text   *text
}

// 资源服务定义
type text struct {
	service *TranslationService
}

// 资源服务方法定义
func (t *text) Detect(ctx context.Context, req *DetectTextReq, options ...larkcore.RequestOptionFunc) (*DetectTextResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/translation/v1/text/detect", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DetectTextResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *text) Translate(ctx context.Context, req *TranslateTextReq, options ...larkcore.RequestOptionFunc) (*TranslateTextResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/translation/v1/text/translate", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &TranslateTextResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
