// Code generated by Lark OpenAPI.

package larkauthen

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type V1 struct {
	AccessToken            *accessToken            // access_token
	OidcAccessToken        *oidcAccessToken        // oidc.access_token
	OidcRefreshAccessToken *oidcRefreshAccessToken // oidc.refresh_access_token
	RefreshAccessToken     *refreshAccessToken     // refresh_access_token
	UserInfo               *userInfo               // user_info
}

func New(config *larkcore.Config) *V1 {
	return &V1{
		AccessToken:            &accessToken{config: config},
		OidcAccessToken:        &oidcAccessToken{config: config},
		OidcRefreshAccessToken: &oidcRefreshAccessToken{config: config},
		RefreshAccessToken:     &refreshAccessToken{config: config},
		UserInfo:               &userInfo{config: config},
	}
}

type accessToken struct {
	config *larkcore.Config
}
type oidcAccessToken struct {
	config *larkcore.Config
}
type oidcRefreshAccessToken struct {
	config *larkcore.Config
}
type refreshAccessToken struct {
	config *larkcore.Config
}
type userInfo struct {
	config *larkcore.Config
}

// Create
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=create&project=authen&resource=access_token&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/authenv1/create_accessToken.go
func (a *accessToken) Create(ctx context.Context, req *CreateAccessTokenReq, options ...larkcore.RequestOptionFunc) (*CreateAccessTokenResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/authen/v1/access_token"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeApp}
	apiResp, err := larkcore.Request(ctx, apiReq, a.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateAccessTokenResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, a.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Create
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=create&project=authen&resource=oidc.access_token&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/authenv1/create_oidcAccessToken.go
func (o *oidcAccessToken) Create(ctx context.Context, req *CreateOidcAccessTokenReq, options ...larkcore.RequestOptionFunc) (*CreateOidcAccessTokenResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/authen/v1/oidc/access_token"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeApp}
	apiResp, err := larkcore.Request(ctx, apiReq, o.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateOidcAccessTokenResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, o.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Create
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=create&project=authen&resource=oidc.refresh_access_token&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/authenv1/create_oidcRefreshAccessToken.go
func (o *oidcRefreshAccessToken) Create(ctx context.Context, req *CreateOidcRefreshAccessTokenReq, options ...larkcore.RequestOptionFunc) (*CreateOidcRefreshAccessTokenResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/authen/v1/oidc/refresh_access_token"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeApp}
	apiResp, err := larkcore.Request(ctx, apiReq, o.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateOidcRefreshAccessTokenResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, o.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Create
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=create&project=authen&resource=refresh_access_token&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/authenv1/create_refreshAccessToken.go
func (r *refreshAccessToken) Create(ctx context.Context, req *CreateRefreshAccessTokenReq, options ...larkcore.RequestOptionFunc) (*CreateRefreshAccessTokenResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/authen/v1/refresh_access_token"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeApp}
	apiResp, err := larkcore.Request(ctx, apiReq, r.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateRefreshAccessTokenResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, r.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// Get
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=get&project=authen&resource=user_info&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/authenv1/get_userInfo.go
func (u *userInfo) Get(ctx context.Context, options ...larkcore.RequestOptionFunc) (*GetUserInfoResp, error) {
	// 发起请求
	apiReq := &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	apiReq.ApiPath = "/open-apis/authen/v1/user_info"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, u.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetUserInfoResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, u.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
