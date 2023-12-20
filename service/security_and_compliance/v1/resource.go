// Code generated by Lark OpenAPI.

package larksecurity_and_compliance

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type V1 struct {
	OpenapiLog *openapiLog // openapi_log
}

func New(config *larkcore.Config) *V1 {
	return &V1{
		OpenapiLog: &openapiLog{config: config},
	}
}

type openapiLog struct {
	config *larkcore.Config
}

// ListData
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=list_data&project=security_and_compliance&resource=openapi_log&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/security_and_compliancev1/listData_openapiLog.go
func (o *openapiLog) ListData(ctx context.Context, req *ListDataOpenapiLogReq, options ...larkcore.RequestOptionFunc) (*ListDataOpenapiLogResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/security_and_compliance/v1/openapi_logs/list_data"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, o.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListDataOpenapiLogResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, o.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
