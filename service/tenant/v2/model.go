// Package tenant code generated by oapi sdk gen
package larktenant

import (
	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 生成枚举值

// 生成数据类型

type Avatar struct {
	AvatarOrigin *string `json:"avatar_origin,omitempty"`
	Avatar72     *string `json:"avatar_72,omitempty"`
	Avatar240    *string `json:"avatar_240,omitempty"`
	Avatar640    *string `json:"avatar_640,omitempty"`
}

// builder开始
type AvatarBuilder struct {
	avatarOrigin     string
	avatarOriginFlag bool
	avatar72         string
	avatar72Flag     bool
	avatar240        string
	avatar240Flag    bool
	avatar640        string
	avatar640Flag    bool
}

func NewAvatarBuilder() *AvatarBuilder {
	builder := &AvatarBuilder{}
	return builder
}

func (builder *AvatarBuilder) AvatarOrigin(avatarOrigin string) *AvatarBuilder {
	builder.avatarOrigin = avatarOrigin
	builder.avatarOriginFlag = true
	return builder
}
func (builder *AvatarBuilder) Avatar72(avatar72 string) *AvatarBuilder {
	builder.avatar72 = avatar72
	builder.avatar72Flag = true
	return builder
}
func (builder *AvatarBuilder) Avatar240(avatar240 string) *AvatarBuilder {
	builder.avatar240 = avatar240
	builder.avatar240Flag = true
	return builder
}
func (builder *AvatarBuilder) Avatar640(avatar640 string) *AvatarBuilder {
	builder.avatar640 = avatar640
	builder.avatar640Flag = true
	return builder
}

func (builder *AvatarBuilder) Build() *Avatar {
	req := &Avatar{}
	if builder.avatarOriginFlag {
		req.AvatarOrigin = &builder.avatarOrigin

	}
	if builder.avatar72Flag {
		req.Avatar72 = &builder.avatar72

	}
	if builder.avatar240Flag {
		req.Avatar240 = &builder.avatar240

	}
	if builder.avatar640Flag {
		req.Avatar640 = &builder.avatar640

	}
	return req
}

// builder结束

type Tenant struct {
	Name      *string `json:"name,omitempty"`
	DisplayId *string `json:"display_id,omitempty"`
	TenantTag *int    `json:"tenant_tag,omitempty"`
	TenantKey *string `json:"tenant_key,omitempty"`
	Avatar    *Avatar `json:"avatar,omitempty"`
}

// builder开始
type TenantBuilder struct {
	name          string
	nameFlag      bool
	displayId     string
	displayIdFlag bool
	tenantTag     int
	tenantTagFlag bool
	tenantKey     string
	tenantKeyFlag bool
	avatar        *Avatar
	avatarFlag    bool
}

func NewTenantBuilder() *TenantBuilder {
	builder := &TenantBuilder{}
	return builder
}

func (builder *TenantBuilder) Name(name string) *TenantBuilder {
	builder.name = name
	builder.nameFlag = true
	return builder
}
func (builder *TenantBuilder) DisplayId(displayId string) *TenantBuilder {
	builder.displayId = displayId
	builder.displayIdFlag = true
	return builder
}
func (builder *TenantBuilder) TenantTag(tenantTag int) *TenantBuilder {
	builder.tenantTag = tenantTag
	builder.tenantTagFlag = true
	return builder
}
func (builder *TenantBuilder) TenantKey(tenantKey string) *TenantBuilder {
	builder.tenantKey = tenantKey
	builder.tenantKeyFlag = true
	return builder
}
func (builder *TenantBuilder) Avatar(avatar *Avatar) *TenantBuilder {
	builder.avatar = avatar
	builder.avatarFlag = true
	return builder
}

func (builder *TenantBuilder) Build() *Tenant {
	req := &Tenant{}
	if builder.nameFlag {
		req.Name = &builder.name

	}
	if builder.displayIdFlag {
		req.DisplayId = &builder.displayId

	}
	if builder.tenantTagFlag {
		req.TenantTag = &builder.tenantTag

	}
	if builder.tenantKeyFlag {
		req.TenantKey = &builder.tenantKey

	}
	if builder.avatarFlag {
		req.Avatar = builder.avatar
	}
	return req
}

// builder结束

// 生成请求和响应结果类型，以及请求对象的Builder构造器

type QueryTenantRespData struct {
	Tenant *Tenant `json:"tenant,omitempty"`
}

type QueryTenantResp struct {
	*larkcore.RawResponse `json:"-"`
	larkcore.CodeError
	Data *QueryTenantRespData `json:"data"`
}

func (resp *QueryTenantResp) Success() bool {
	return resp.Code == 0
}

// 生成消息事件结构体

// 生成请求的builder构造器
// 1.1 生成body的builder结构体
