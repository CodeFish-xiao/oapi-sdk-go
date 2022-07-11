// Package baike code generated by oapi sdk gen
package larkbaike

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *BaikeService {
	b := &BaikeService{config: config}
	b.Classification = &classification{service: b}
	b.Draft = &draft{service: b}
	b.Entity = &entity{service: b}
	return b
}

// 业务域服务定义
type BaikeService struct {
	config         *larkcore.Config
	Classification *classification
	Draft          *draft
	Entity         *entity
}

// 资源服务定义
type classification struct {
	service *BaikeService
}
type draft struct {
	service *BaikeService
}
type entity struct {
	service *BaikeService
}

// 资源服务方法定义
func (c *classification) List(ctx context.Context, req *ListClassificationReq, options ...larkcore.RequestOptionFunc) (*ListClassificationResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, c.service.config, http.MethodGet,
		"/open-apis/baike/v1/classifications", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListClassificationResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *classification) ListByIterator(ctx context.Context, req *ListClassificationReq, options ...larkcore.RequestOptionFunc) (*ListClassificationIterator, error) {
	return &ListClassificationIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (d *draft) Create(ctx context.Context, req *CreateDraftReq, options ...larkcore.RequestOptionFunc) (*CreateDraftResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPost,
		"/open-apis/baike/v1/drafts", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateDraftResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (d *draft) Update(ctx context.Context, req *UpdateDraftReq, options ...larkcore.RequestOptionFunc) (*UpdateDraftResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, d.service.config, http.MethodPut,
		"/open-apis/baike/v1/drafts/:draft_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateDraftResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *entity) Create(ctx context.Context, req *CreateEntityReq, options ...larkcore.RequestOptionFunc) (*CreateEntityResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodPost,
		"/open-apis/baike/v1/entities", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateEntityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *entity) Get(ctx context.Context, req *GetEntityReq, options ...larkcore.RequestOptionFunc) (*GetEntityResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodGet,
		"/open-apis/baike/v1/entities/:entity_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetEntityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *entity) Highlight(ctx context.Context, req *HighlightEntityReq, options ...larkcore.RequestOptionFunc) (*HighlightEntityResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodPost,
		"/open-apis/baike/v1/entities/highlight", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &HighlightEntityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *entity) List(ctx context.Context, req *ListEntityReq, options ...larkcore.RequestOptionFunc) (*ListEntityResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodGet,
		"/open-apis/baike/v1/entities", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListEntityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *entity) ListByIterator(ctx context.Context, req *ListEntityReq, options ...larkcore.RequestOptionFunc) (*ListEntityIterator, error) {
	return &ListEntityIterator{
		ctx:      ctx,
		req:      req,
		listFunc: e.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (e *entity) Match(ctx context.Context, req *MatchEntityReq, options ...larkcore.RequestOptionFunc) (*MatchEntityResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodPost,
		"/open-apis/baike/v1/entities/match", []larkcore.AccessTokenType{larkcore.AccessTokenTypeUser, larkcore.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &MatchEntityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *entity) Search(ctx context.Context, req *SearchEntityReq, options ...larkcore.RequestOptionFunc) (*SearchEntityResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodPost,
		"/open-apis/baike/v1/entities/search", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchEntityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *entity) SearchByIterator(ctx context.Context, req *SearchEntityReq, options ...larkcore.RequestOptionFunc) (*SearchEntityIterator, error) {
	return &SearchEntityIterator{
		ctx:      ctx,
		req:      req,
		listFunc: e.Search,
		options:  options,
		limit:    req.Limit}, nil
}
func (e *entity) Update(ctx context.Context, req *UpdateEntityReq, options ...larkcore.RequestOptionFunc) (*UpdateEntityResp, error) {
	// 发起请求
	rawResp, err := larkcore.SendRequest(ctx, e.service.config, http.MethodPut,
		"/open-apis/baike/v1/entities/:entity_id", []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateEntityResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
