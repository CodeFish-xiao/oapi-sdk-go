// Package optical_char_recognition code generated by oapi sdk gen
package larkoptical_char_recognition

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *OpticalCharRecognitionService {
	o := &OpticalCharRecognitionService{config: config}
	o.Image = &image{service: o}
	return o
}

// 业务域服务定义
type OpticalCharRecognitionService struct {
	config *larkcore.Config
	Image  *image
}

// 资源服务定义
type image struct {
	service *OpticalCharRecognitionService
}

// 资源服务方法定义
func (i *image) BasicRecognize(ctx context.Context, req *BasicRecognizeImageReq, options ...larkcore.RequestOptionFunc) (*BasicRecognizeImageResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, i.service.config, http.MethodPost,
		"/open-apis/optical_char_recognition/v1/image/basic_recognize", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BasicRecognizeImageResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
